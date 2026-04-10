---
name: gorch-dev
description: 高效维护和演进 Gorch 项目，包括 DSL 语法/解析、运行时引擎行为、算子 DI 与容器 fast-path、gorchc 代码生成及回归测试。编辑 iantlr/、internal/lang/、internal/engine/、gorchc/、examples/，或新增 DSL 指令、调整执行语义、优化生成算子性能时使用此 skill。
---

# Gorch 开发指南

## 概览

使用此 skill 以低回归方式开展 Gorch 开发。
先按下方“任务路由”确定主流程，再执行对应参考文档里的清单。

## 当前 DSL 静态检查基线

`internal/lang/ast` 当前默认启用以下自检（`Primary.SelfCheck`）：

- 未注册算子检查：递归覆盖 `START` / `ON_FINISH` / `UNFOLD` / `GO` / `SWITCH` / `LOOP` / `RETRY` 等路径。
- 必填参数检查：`LOOP` 要求 `MAX_ITER|maxIter > 0`，`RETRY` 要求 `MAX_TIMES|maxTimes > 0`。
- GO/WAIT 双向死链路检查：`WAIT` 必须匹配 `GO`，`GO` 也必须匹配 `WAIT`。
- 不可达分支检查（增强版）：串行链路里，`BREAK` 及其可静态推导的终止结构（如括号串行、`UNFOLD`、全分支终止的 `SWITCH`）后续节点判定为不可达。

## 任务路由

改代码前先选一个主流程：

1. 修改 DSL 语法，或新增/删除指令  
阅读并遵循 [references/dsl-change-workflow.md](references/dsl-change-workflow.md)。

2. 修改运行时执行语义（Context/Processor/Engine/Container）  
阅读 [references/runtime-and-di.md](references/runtime-and-di.md)，先跑定向测试，再跑 `go test ./...`。

3. 修改 gorchc 生成流程或算子包装行为  
阅读 [references/runtime-and-di.md](references/runtime-and-di.md)，并在 `examples/gen` 验证生成代码。

4. 需要快速熟悉陌生模块  
先读 [references/project-map.md](references/project-map.md)。

## 强制规则

- 不要手改 `internal/lang/iantlr/` 下的生成文件；必须从 `iantlr/gorchlang.g4` 重新生成。
- DSL 语义变化时，解析器/AST/运行时测试要同步更新。
- 除非用户明确同意破坏性变更，否则保持向后兼容。
- 优先做最小范围修改；先跑相关包测试，再跑全量回归。

## 标准验证

完成前执行：

```bash
go test ./internal/lang/iparser
go test ./internal/engine
go test ./...
```

如果改了 gorchc 模板或生成流程，额外在 examples 里走一次真实生成流程验证输出。

## 参考文档

- [references/project-map.md](references/project-map.md)：模块分布与入口文件。
- [references/dsl-change-workflow.md](references/dsl-change-workflow.md)：指令改动从语法到运行时的完整清单。
- [references/runtime-and-di.md](references/runtime-and-di.md)：engine/container/gorchc 的 DI 行为与 fast-path 说明。
