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


### 2.1 算子

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



