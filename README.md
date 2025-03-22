# gorch

## 0. gorch 介绍

*gorch* `/ˈɡɔːrʃ/`， 戈奇，你也可以直接叫`枸杞`。它是一个基于图描述的执行框架。


它设计了一套能够描述执行流程的 DSL 语言， 并与之对应的执行引擎， 能够根据 DSL 描述的执行流程执行程序。

它不是分布式的，而是一个单机的、无状态的执行框架。

## 1. 开发背景

21年的时候，我在百度搜索部门内开始着手一个PHP架构模块迁移Go的方案调研。彼时，内部不少C++模块都已经实现模块图化（基于DAG图实现的执行框架），也初识图化模块的好处:

* 清晰的依赖管理
* 可视化的并行、串行的执行
* 良好的容错性和可恢复性
* 策略与策略之间的解耦

所以我们在考虑方案时也要求将模块图化执行。

所以我开发了exgraph图执行引擎，并应用到搜索展现架构PHP迁移Go的项目中，并取得了不错的效果。感兴趣的同学可以看下介绍：[百度搜索exgraph图执行引擎设计与实践](https://mp.weixin.qq.com/s/Ovgzq3zoXteMMcAPrpn_DA)。

从百度离开后，我开始着手开发**gorch**，想要实现一个应用面更广的图执行引擎。


## 2. 基本概念

### 2.1 执行块

**gorch**中，内置三个执行块，分别是：

* `START` 执行块，表示程序的入口。
* `FRAGMENT` 执行块，表示一个执行片段。它可以被`START`块、或者另一个`FRAGMENT`块包含。语法解释器内部有做执行环检查。
* `REGISTER`执行块，用来注册用户的算子。

#### 2.1.1 START执行块

`START` 执行块基本语法如下：

```
START("my_start"){
	aoperator
}
```

上面表示在`my_start`的执行块中，只执行了`aoperator`算子。

> `START`执行块可以有多个，但其名称需要全局唯一

#### 2.1.2 FRAGMENT执行块

`FRAGMENT` 执行块基本语法如下：
```
FRAGMENT("my_fragment"){
	aoperator
}
```
上面表示在`my_fragment`的执行块中，只执行了`aoperator`算子。

> `FRAGMENT`执行块可以有多个，但其名称需要全局唯一

`FRAGMENT`块不可以单独执行，需要被`START`或者另一个`FRAGMENT`块包含。包含时的语法如下：

```
START("my_start"){
    UNFOLD("my_fragment")
}
```

上面表示在`my_start`执行块中，利用`UNFOLD`指令包含了`my_fragment`执行块。

#### 2.1.3 REGISTER执行块

`REGISTER` 执行块基本语法如下：

```
REGISTER("github.com/my_repo"){
    OPERATOR("finder2/my_operator", "MyOperator", 1)
    OPERATOR("finder1/my_operator", "MyOperator2", "myOperator3" 3)
}
```

上面表示在`github.com/my_repo`的执行块中，注册了两个算子。


`OPERATOR`指令语法如下：

```
// 3个参数：算子所在路径、算子struct名称、算子序号
// 没有填入算子名时，默认使用算子struct名称
OPERATOR("finder2/my_operator", "MyOperator", 1)

// 4个参数：算子所在路径、算子struct名称、算子的命名、算子序号
OPERATOR("finder2/my_operator", "MyOperator2", "myOperator3" 3)
```

> 算子名需要全局唯一
> 
> 算子序号需要也需要全局唯一

### 2.2 串行与并发

在`START`或者`FRAGMENT`执行块中，你可以自由选择算子的执行方式，是串行还是并发。

我们约定，用`->`表示串行执行，一般我们称为串行组，用`[]`表示并发执行，所以称为并发组。

例如：`aoperator -> boperator -> coperator`表示`aoperator`算子先执行，然后是`boperator`算子，最后是`coperator`算子。

再例如：`aoperator -> [boperator, coperator]`表示`aoperator`算子先执行，然后是`boperator`算子和`coperator`算子并发地执行。

### 2.2 可执行元素

下面，我们来介绍，在执行图中，可执行的元素。

### 2.2.1 算子

算子是图执行的最小单元。它是一个实现了`Processor`接口的struct。

```go
type MyOperator struct { }

func(o *MyOperator) Execute(ctx *gorch.Context) error {
    return nil
}
```

在`Execute`方法中，你可以实现自己的逻辑。

在图描述中，算子就是一个单词：

```
START("my_start"){
    aoperator
}
```
上面表示在`my_start`的执行块中，只执行了`aoperator`算子。

每个算子你都可以从执行图传递一定的参数到代码中。
例如：
```
START("my_start"){
    aoperator(val=1)
    -> boperator(val="2")
    -> coperator(val=true)
    -> doperator(val=10ms)
    -> eoperator(val=[1,2,3])
    -> foperator(val=["1", "2", "3"])
    -> goperator(val=[true, false])
    -> hoperator(val=[10s, 20ms])
}
```

算子的参数，支持4中类型：`int64`, `string`, `bool`, `time.Duration`。

用`[]`包裹起来后，表示是一个数组。

在算子的代码中，你可以用如下接口获取算子的参数：

```go
// ArgValue 参数值接口
type ArgValue interface {
	Int64(...int) int64
	Bool(...int) bool
	Str(...int) string
	Duration(...int) time.Duration

	Int64List() []int64
	BoolList() []bool
	StrList() []string
	DurationList() []time.Duration
}
```

```go
type aOp struct {}

func (a *aOp) Execute(ctx *Context) error {
    boolVal := ctx.Arg("val1").Bool()
    intVal := ctx.Arg("val2").Int64()
    strVal := ctx.Arg("val3").Str()
    durationVal := ctx.Arg("val4").Duration()

    boolList := ctx.Arg("val5").BoolList()
    intList := ctx.Arg("val6").Int64List()
    strList := ctx.Arg("val7").StrList()
    durationList := ctx.Arg("val8").DurationList()
    return nil
}
```

### 2.2.2 串行组

前面说到，在执行图上，用`->`表示串行执行，称为串行组。

实际上，一个串行组，也可以包含另一个串行组，就像这样：```aoperator -> (boperator -> coperator)```。

我们用`()`，在一个串行组内，又创建了一个子集的串行组。

为什么可以这样呢？

因为不论是算子、串行组、并发组，已经我们后面介绍的各种可执行元素，都在内部实现了`Processor`接口。

### 2.2.3 并发组

不做过多介绍。

### 2.2.4 SWITCH指令：选择执行分支

`SWITCH`指令也实现了`Processor`接口。

它的基本语法如下：

```
START("my_start"){
    SWITCH(SwitchOp) {
        CASE "changeValueTo1" => ChangeValueOP(val=1),
        CASE "changeValueTo2" => ChangeValueOP(val=2),
        CASE "changeValueTo3" => ChangeValueOP(val=3)
    }
}
```

`SWITCH`指令传入一个算子，这个算子在执行时，必须显示的调用`ctx.Switch()`方法，传入一个值，然后`SWITCH`指令会根据这个值，选择执行对应的分支。

### 2.2.5 包装算子

`WRAP`指令实现了`Processor`接口，同样可以嵌入在图上任意位置。

`WRAP`指令的基本语法如下：

```
START("my_start"){
    (wrapOp | aoperator)
        -> (wrapOp2 | wrapOp3 | (boperator -> coperator))
}
```

包装算子通过`|`符号，将算子包裹起来，借鉴了linux下管道的用法。

你可以包装一个算子，也可以包装实现了`Processor`接口的可执行元素。甚至你可以包装另一个包装算子。

需要注意的是，包装算子你应该用`()`包起来。

### 2.2.6 SKIP指令：跳过串行组

`SKIP`指令用于在串行组内，跳过剩余执行元素。

`SKIP`指令的基本语法如下：

```
START("my_start"){
    aoperator -> SKIP(boperator) -> coperator -> doperator
}
```

在上面的执行图中，进入`boperator`算子后，如果你调用`return ctx.SkipSerial()`，那么`coperator`算子将不会被执行，即`coperator`算子和`doperator`不会被执行。

如果你只想跳过`coperator`算子，你应该这么编排你的执行图：

```
START("my_start"){
    aoperator -> (SKIP(boperator) -> coperator) -> doperator
}
```

和你想的一样，用`()`把`SKIP(boperator) -> coperator`包裹起来就只跳过`coperator`算子。


### 2.2.7 GO指令：后台的去执行

很多时候，有些执行任务并不需要马上就获取结果，比如你要请求一个网络后端A，为了节省服务器执行时间，发送完请求就去执行别的任务，等到需要A后端的结果时，再去WAIT它。

`GO`指令就是为了满足这种需求的。

`GO`指令总是和`WAIT`指令一起使用。

它们的基本语法如下：

```
START("test"){

    GO(GoOperator(sleep=20ms, val=10), "goOperator")
    -> SleepOp(sleep=2ms) // 等待goroutine启动
    -> NothingOp(val="1", WAIT("goOperator", timeout=5ms))
}
```

调用`Go`指令，传入一个执行元素（即你可以传入算子，也可以是其他执行元素）以及名称，那么这个执行元素，就会在一个新的goroutine中执行。

`WAIT`指令，会等待`GO`指令启动的goroutine执行完毕。

`WAIT`指令不是单独存在的，它必须作为算子执行的最后参数出现，一个算子可以`WAIT`多个`GO`指令。

如果一个算子有自己的参数，那么`WAIT`指令必须放在其参数之后。

为适应不同的业务场景，`WAIT`指令有几种超时机制：

* `WAIT("goOperator", timeout=5ms)`：从当前时间开始计算，等待`goOperator`执行5ms，如果5ms内没有执行完毕，那么就返回超时错误。如果算子还没开始执行，就直接返回错误`routine not started`。
* `WAIT("goOperator", timeout=5ms, noCheckStart=true)`：从当前时间开始计算，等待`goOperator`执行5ms，如果5ms内没有执行完毕，那么就返回超时错误。不管算子有没有开始执行，反正就等待5ms。
* `WAIT("goOperator", totalTimeout=5ms)`：从算子开始执行时A点计算，一共等待`goOperator`执行5ms。如果算子还没开始执行，就返回错误`routine not started`。如果算子已经执行超时，就返回`routine execute timeout`。
