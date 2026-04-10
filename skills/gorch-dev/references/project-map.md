# 项目地图

## 顶层目录

- `gorch.go`：公共 API 入口（`New`、`ParseDSL`、`ParseDSLInDir`、`RegOp`）。
- `README.md`：语言特性、运行时概念与示例说明。
- `examples/`：可运行 DSL、算子样例与生成代码输出。

## DSL 与解析器

- `iantlr/gorchlang.g4`：语法源文件（单一事实来源）。
- `internal/lang/iantlr/`：生成的 lexer/parser/listener 文件。
- `internal/lang/ast/`：AST 节点、监听器绑定、语义校验。
- `internal/lang/iparser/`：解析入口与语法/自检测试。

## 运行时引擎

- `internal/engine/engine.go`：engine/executor 生命周期。
- `internal/engine/processor.go`：AST 到可执行处理器的映射。
- `internal/engine/context.go`：运行时上下文与流程控制方法。
- `internal/engine/container.go`：DI 存储与 typed 对象传递。
- `internal/engine/operator_fields.go`：反射注入/导出的兜底路径。
- `internal/engine/*_test.go`：执行语义测试集合。

## 代码生成（gorchc）

- `gorchc/main.go`：CLI 参数入口（`-g`、`-i`、`-d`、`-cmd`、`-p`）。
- `gorchc/generate.go`：生成流程编排。
- `gorchc/tpls/operator.tpl`：生成算子包装代码（非反射路径标记）。
- `gorchc/tpls/operator_import.tpl`：生成注册所需 import 代码。

## 相关模块（可选）

- `vscode-gorch-extension/`：DSL 高亮、snippet 与 IDE 辅助。仅在语法关键字变更时同步更新。
