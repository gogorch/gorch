# 运行时与 DI 说明

## 执行链路

1. DSL 解析为 AST（`iparser`）。
2. 基于 AST 构建处理器（`internal/engine/processor.go` 中的 `createProcessor`）。
3. 通过 `Executor.Execute` 运行；单次引擎执行共享同一个 container。

## DI 路径

当前有两条 DI 路径：

1. 反射元信息路径  
触发条件：算子未实现 `IsGenerateOperatorCode`。  
位置：`internal/engine/operator_fields.go` + `factory.go`。

2. 生成包装路径（推荐）  
触发条件：由 `gorchc/tpls/operator.tpl` 生成的包装算子。  
行为：包装层通过 `gorch.MutableTyped` / `gorch.RegisterTyped` 做 inject/extract，再调用真实算子。

## Fast-Path 方向

优化 DI 性能时遵循：

1. container 以内存里的 typed 指针缓存作为单一运行时存储。
2. `registerAny/mutableAny` 仅保留为动态入口（`Executor.Inject`），不再做旧路径镜像。
3. 生成代码持续走 typed API（`MutableTyped/RegisterTyped`）。
4. 先跑 `internal/engine`，再跑全量测试。

## 高风险文件

- `internal/engine/container.go`：并发与生命周期敏感。
- `internal/engine/context.go`：共享运行态和控制标记。
- `internal/engine/processor.go`：执行语义与顺序控制。

这些文件应尽量小步修改，并配合定向测试验证行为。

## 推荐回归命令

```bash
go test ./internal/engine -race
go test ./...
```
