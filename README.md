# Gorch 🚀

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)]()

*Gorch* `/ˈɡɔːrʃ/` （戈奇，也可以叫"枸杞"）是一个高性能的**基于图描述的执行框架**，专为复杂业务流程编排而设计。

## ✨ 核心特性

- 🎯 **DSL驱动**：直观的图描述语言，让复杂流程一目了然
- ⚡ **高性能**：基于Go协程的并发执行引擎
- 🔧 **代码生成**：自动生成算子注册和导入代码
- 🎨 **VS Code支持**：完整的IDE扩展，提供语法高亮、智能提示、错误检查
- 🔄 **灵活编排**：支持串行、并行、条件分支、异步执行等多种模式
- 🛡️ **容错机制**：内置错误处理、超时控制、状态恢复
- 📊 **可观测性**：完整的执行记录和性能监控

## 🏗️ 架构概览

Gorch采用分层架构设计，确保高性能和可扩展性：

```
┌─────────────────────────────────────────────────────────────┐
│                    Gorch DSL 语言层                          │
├─────────────────────────────────────────────────────────────┤
│  编译器 (gorchc)  │  VS Code 扩展  │  语法解析器 (ANTLR4)    │
├─────────────────────────────────────────────────────────────┤
│                    执行引擎 (Engine)                         │
├─────────────────────────────────────────────────────────────┤
│  协程池 (Pool)   │  日志系统 (MLog)  │  状态记录 (Recorder)   │
└─────────────────────────────────────────────────────────────┘
```

## 🚀 快速开始

### 安装

```bash
go get github.com/gogorch/gorch
```

### 基本使用

1. **定义算子**（业务逻辑单元）

```go
type MyOperator struct{}

func (o *MyOperator) Execute(ctx *gorch.Context) error {
    // 获取参数
    value := ctx.Arg("value").Int64()
    name := ctx.Arg("name").Str()

    // 执行业务逻辑
    fmt.Printf("Processing %s with value %d\n", name, value)
    return nil
}
```

2. **编写Gorch DSL**（`example.gorch`）

```gorch
// 注册算子
REGISTER("github.com/example/operators") {
    OPERATOR("myop", "MyOperator", "MyOp", 1)
}

// 定义执行流程
START("main") {
    MyOp(value=100, name="test")
    -> [ParallelOp1(), ParallelOp2()]  // 并行执行
    -> FinalOp()
}
```

3. **运行程序**

```go
func main() {
    // 解析DSL
    primary, err := gorch.ParseDSLInDir("./conf")
    if err != nil {
        log.Fatal(err)
    }

    // 创建执行引擎
    engine, err := gorch.New(primary)
    if err != nil {
        log.Fatal(err)
    }

    // 执行
    executor := engine.Start("main")
    result := executor.Execute()

    fmt.Printf("Execution result: %v\n", result.Error())
}
```

## 📖 开发背景

Gorch诞生于对高效业务流程编排的需求。在百度搜索部门的大规模系统迁移项目中，我们发现基于DAG图的执行框架具有显著优势：

- 🎯 **清晰的依赖管理** - 业务逻辑关系一目了然
- 🔄 **可视化执行流程** - 并行、串行执行路径清晰可见
- 🛡️ **优秀的容错性** - 支持错误恢复和状态管理
- 🔧 **模块解耦** - 策略与策略之间松耦合设计

基于这些实践经验，Gorch致力于打造一个更通用、更强大的图执行引擎。

> 📚 相关阅读：[百度搜索exgraph图执行引擎设计与实践](https://mp.weixin.qq.com/s/Ovgzq3zoXteMMcAPrpn_DA)


## 🏗️ 核心概念

### 执行块 (Execution Blocks)

Gorch提供三种核心执行块，构建完整的业务流程：

| 执行块 | 作用 | 特点 |
|--------|------|------|
| `START` | 程序入口点 | 可定义多个，名称全局唯一 |
| `FRAGMENT` | 可复用的执行片段 | 支持嵌套调用，自动检测循环依赖 |
| `REGISTER` | 算子注册中心 | 管理所有业务算子的元数据 |

#### START 执行块

程序的入口点，定义主要的执行流程：

```gorch
START("user_service") {
    ValidateUser(userId=123)
    -> LoadUserProfile()
    -> [LoadPreferences(), LoadHistory()]  // 并行加载
    -> BuildResponse()
}
```

#### FRAGMENT 执行块

可复用的执行片段，支持模块化设计：

```gorch
FRAGMENT("user_validation") {
    CheckUserExists()
    -> ValidatePermissions()
    -> LogAccess()
}

START("main") {
    UNFOLD("user_validation")  // 引用fragment
    -> ProcessRequest()
}
```

#### REGISTER 执行块

算子注册中心，管理所有业务算子：

```gorch
REGISTER("github.com/example/operators") {
    OPERATOR("user", "UserValidator", "ValidateUser", 1)
    OPERATOR("profile", "ProfileLoader", "LoadProfile", 2)
    OPERATOR("response", "ResponseBuilder", "BuildResponse", 3)
}
```

### 执行模式 (Execution Patterns)

#### 串行执行 (`->`)
```gorch
START("sequential") {
    StepA() -> StepB() -> StepC()  // 顺序执行
}
```

#### 并行执行 (`[]`)
```gorch
START("parallel") {
    [TaskA(), TaskB(), TaskC()]  // 同时执行
}
```

#### 混合模式
```gorch
START("complex") {
    Init()
    -> [LoadData(), ValidateInput()]  // 并行
    -> ProcessData()                  // 串行
    -> [SaveResult(), SendNotification()]  // 并行
}
```

#### 条件分支 (`SWITCH`)
```gorch
START("conditional") {
    SWITCH(CheckCondition) {
        CASE "success" => ProcessSuccess(),
        CASE "error" => HandleError(),
        CASE "retry" => RetryOperation()
    }
}
```

### 高级特性

#### 异步执行 (`GO` & `WAIT`)

Gorch支持强大的异步执行模式，允许任务在后台运行，提高整体执行效率：

```gorch
START("async_demo") {
    // 启动后台任务
    GO(SlowNetworkCall(), "network_task")
    -> GO(DatabaseQuery(), "db_task")

    // 立即执行其他任务
    -> FastLocalOperation()

    // 等待后台任务完成
    -> ProcessResult(
        WAIT("network_task", timeout=30s),      // 等待网络任务，最多30秒
        WAIT("db_task", totalTimeout=45s)       // 从任务开始算，总共45秒
    )
}
```

**WAIT 指令参数详解：**

| 参数 | 说明 | 示例 |
|------|------|------|
| `timeout` | 等待超时时间（从开始等待时计算） | `timeout=30s` |
| `totalTimeout` | 总超时时间（从任务开始执行时计算） | `totalTimeout=60s` |
| `notCheckStart` | 允许等待未启动的任务 | `notCheckStart=true` |

**等待模式对比：**

```gorch
START("wait_modes") {
    GO(SlowTask(), "task1")
    -> GO(SlowTask(), "task2")
    -> GO(SlowTask(), "task3")

    // 立即执行其他操作，耗时10秒
    -> SomeOtherWork(duration=10s)

    // 不同的等待模式
    -> ProcessResults(
        // 模式1：timeout - 从现在开始等待20秒
        WAIT("task1", timeout=20s),

        // 模式2：totalTimeout - 从task2开始执行算起，总共30秒
        WAIT("task2", totalTimeout=30s),

        // 模式3：无限等待
        WAIT("task3")
    )
}
```

#### 包装算子 (`WRAP`)
```gorch
START("wrapped") {
    (LogWrapper | BusinessLogic())  // 为业务逻辑添加日志
    -> (RetryWrapper | NetworkCall())  // 为网络调用添加重试
}
```

#### 跳过控制 (`SKIP`)
```gorch
START("conditional_skip") {
    CheckCondition()
    -> SKIP(ProcessA())  // 可能被跳过
    -> ProcessB()        // 总是执行
}
```

## 🛠️ 工具链

### Gorch 编译器 (`gorchc`)

强大的代码生成工具，自动化繁琐的算子注册工作：

```bash
# 生成所有代码到指定目录
gorchc -g ./gen -d ./conf

# 仅生成导入代码
gorchc -i ./gen -d ./conf

# 打印解析结果（调试用）
gorchc -p -d ./conf
```

**自动生成内容：**
- 算子注册代码
- 导入声明
- 类型检查代码

### VS Code 扩展

完整的IDE支持，提升开发效率：

- 🎨 **语法高亮** - Gorch DSL语法着色
- 💡 **智能提示** - 算子名称、参数自动补全
- 🔍 **定义跳转** - 快速跳转到算子实现
- ❌ **错误检查** - 实时语法和语义检查
- 📊 **诊断信息** - 详细的错误和警告信息

安装方式：在VS Code扩展市场搜索 "Gorch Language Support"

## 📁 项目结构

```
gorch/
├── 📁 examples/           # 示例项目
│   ├── conf/             # Gorch DSL配置文件
│   ├── operators/        # 业务算子实现
│   └── gen/              # 自动生成的代码
├── 📁 gorchc/            # Gorch编译器
│   ├── main.go          # 编译器入口
│   ├── generate.go      # 代码生成逻辑
│   └── tpls/            # 代码模板
├── 📁 internal/          # 核心实现
│   ├── engine/          # 执行引擎
│   ├── lang/            # 语言解析器
│   └── ort/             # 运行时类型系统
├── 📁 pool/              # 协程和对象池
├── 📁 mlog/              # 日志系统
├── 📁 recorder/          # 执行记录器
└── gorch.go             # 公共API入口
```

## 🎯 算子开发

### 基本算子

实现 `Processor` 接口：

```go
type UserValidator struct{}

func (v *UserValidator) Execute(ctx *gorch.Context) error {
    // 获取参数
    userId := ctx.Arg("userId").Int64()
    strict := ctx.Arg("strict").Bool()

    // 执行业务逻辑
    if userId <= 0 {
        return fmt.Errorf("invalid user id: %d", userId)
    }

    return nil
}
```

### 参数类型支持

Gorch支持丰富的参数类型：

```gorch
START("param_demo") {
    MyOperator(
        intVal=42,                    // int64
        strVal="hello",               // string
        boolVal=true,                 // bool
        durationVal=30s,              // time.Duration
        intList=[1,2,3],              // []int64
        strList=["a","b","c"],        // []string
        boolList=[true,false],        // []bool
        durationList=[1s,2m,3h]       // []time.Duration
    )
}
```

## ⚡ 性能特性

### 高效的协程管理
- **协程池复用** - 避免频繁创建销毁协程的开销
- **智能调度** - 根据系统负载动态调整并发度
- **内存优化** - 对象池减少GC压力

### 零拷贝数据传递
- **上下文共享** - 算子间高效的数据传递
- **引用传递** - 避免不必要的数据复制
- **流式处理** - 支持大数据量的流式处理

### 可观测性
```go
// 启用执行记录
recorder := recorder.New()
executor := engine.Start("main").WithRecorder(recorder)

// 获取执行统计
stats := recorder.GetStats()
fmt.Printf("Total time: %v\n", stats.TotalTime)
fmt.Printf("Operator count: %d\n", stats.OperatorCount)
```

## 📚 最佳实践

### 1. 算子设计原则
- **单一职责** - 每个算子只做一件事
- **无状态** - 算子不应该保存状态信息
- **幂等性** - 重复执行应该产生相同结果
- **错误处理** - 合理的错误处理和恢复机制

### 2. 性能优化建议
```go
// ✅ 推荐：使用对象池
func (o *MyOperator) Execute(ctx *gorch.Context) error {
    buffer := pool.GetBuffer()
    defer pool.PutBuffer(buffer)
    // 使用buffer...
}

// ✅ 推荐：合理使用并行
START("optimized") {
    PrepareData()
    -> [ProcessA(), ProcessB(), ProcessC()]  // CPU密集型任务并行
    -> AggregateResults()
}
```

### 3. 错误处理模式

#### 基础错误处理
```go
func (o *NetworkOperator) Execute(ctx *gorch.Context) error {
    // 设置超时
    ctx.SetTimeout(30 * time.Second)

    // 重试机制
    for i := 0; i < 3; i++ {
        if err := o.doNetworkCall(); err == nil {
            return nil
        }
        time.Sleep(time.Duration(i+1) * time.Second)
    }

    return fmt.Errorf("network call failed after 3 retries")
}
```

#### 状态码模式 (`NewOperatorStates`)

Gorch支持更精细的状态码管理，用于区分不同的执行结果：

```go
package operators

import "github.com/gogorch/gorch"

// 定义算子状态集合
var (
    userStates = gorch.NewOperatorStates("UserOperator")

    // 定义不同级别的状态码
    UserNotFound    = userStates.Fatal(1001, "user not found")
    UserInactive    = userStates.Info(1002, "user is inactive")
    UserValidated   = userStates.Info(1003, "user validation success")
    UserCacheHit    = userStates.Info(1004, "user loaded from cache")
)

type UserOperator struct{}

func (o *UserOperator) Execute(ctx *gorch.Context) error {
    userId := ctx.Arg("userId").Int64()

    // 根据不同情况返回对应状态码
    if userId <= 0 {
        return UserNotFound  // Fatal级别，会中断执行
    }

    user, err := o.loadUser(userId)
    if err != nil {
        return UserNotFound
    }

    if !user.IsActive {
        return UserInactive  // Info级别，记录日志但继续执行
    }

    // 检查缓存命中
    if user.FromCache {
        return UserCacheHit
    }

    return UserValidated
}
```

**状态码级别说明：**
- `Fatal(code, msg)` - 致命错误，会中断整个执行流程
- `Info(code, msg)` - 信息状态，记录日志但不中断执行

**使用场景：**
- 🔍 **业务状态追踪** - 精确记录每个算子的执行状态
- 📊 **监控和统计** - 基于状态码进行业务指标统计
- 🐛 **问题诊断** - 快速定位问题发生的具体环节
- 📈 **性能分析** - 分析不同执行路径的性能表现

## 🤝 贡献指南

我们欢迎所有形式的贡献！

### 如何贡献
1. **Fork** 本仓库
2. **创建** 特性分支 (`git checkout -b feature/amazing-feature`)
3. **提交** 更改 (`git commit -m 'Add amazing feature'`)
4. **推送** 到分支 (`git push origin feature/amazing-feature`)
5. **创建** Pull Request

### 开发环境设置
```bash
# 克隆仓库
git clone https://github.com/gogorch/gorch.git
cd gorch

# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 构建编译器
go build -o gorchc ./gorchc
```

### 代码规范
- 遵循 Go 官方代码规范
- 添加必要的单元测试
- 更新相关文档
- 提交信息使用英文，格式清晰

## 📖 相关资源

### 文档
- [快速开始指南](examples/README.md)
- [API 参考文档](https://pkg.go.dev/github.com/gogorch/gorch)
- [VS Code 扩展使用指南](vscode-gorch-extension/USAGE_GUIDE.md)

### 示例项目
- [基础示例](examples/) - 展示核心功能的使用

### 社区
- [GitHub Issues](https://github.com/gogorch/gorch/issues) - 问题反馈和功能请求
- [GitHub Discussions](https://github.com/gogorch/gorch/discussions) - 社区讨论

## 🔄 版本历史

### v1.0.0 (当前版本)
- ✅ 核心执行引擎
- ✅ Gorch DSL 语言支持
- ✅ 代码生成工具 (gorchc)
- ✅ VS Code 扩展
- ✅ 完整的文档和示例

### 路线图
- 🔄 可视化流程编辑器
- 🔄 云原生部署支持

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE)。

## 🙏 致谢

感谢所有为 Gorch 项目做出贡献的开发者们！

特别感谢：
- [ANTLR4](https://github.com/antlr/antlr4) - 强大的语法解析器生成器
- [VS Code](https://code.visualstudio.com/) - 优秀的代码编辑器平台
- Go 社区 - 提供了优秀的生态系统

---

<div align="center">

**如果 Gorch 对你有帮助，请给我们一个 ⭐️**

[🏠 首页](https://github.com/gogorch/gorch) • [📖 文档](https://pkg.go.dev/github.com/gogorch/gorch) • [🐛 问题反馈](https://github.com/gogorch/gorch/issues) • [💬 讨论](https://github.com/gogorch/gorch/discussions)

</div>
