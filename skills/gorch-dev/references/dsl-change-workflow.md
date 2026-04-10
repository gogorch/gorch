# DSL 变更流程

新增/修改/删除 DSL 指令或语法规则时，按以下清单执行。

## 1. 修改语法文件

1. 修改 `iantlr/gorchlang.g4`。
2. 规则结构要与现有 AST 监听模式保持一致（`EnterXxx` / `ExitXxx`）。

## 2. 重新生成解析器

优先使用已安装的 `antlr4` 命令；否则使用本地 jar 通过 java 执行。

示例：

```bash
rm -f internal/lang/iantlr/*.go internal/lang/iantlr/*.tokens internal/lang/iantlr/*.interp
java -jar /Users/ws/Project/antlr-4.13.2-complete.jar \
  -Dlanguage=Go -no-visitor -package iantlr \
  iantlr/gorchlang.g4 -o internal/lang/iantlr
mv internal/lang/iantlr/iantlr/* internal/lang/iantlr/ && rmdir internal/lang/iantlr/iantlr
```

不要手改 `internal/lang/iantlr/` 下的生成文件。

## 3. 串接 AST

1. 如有需要，在 `internal/lang/ast/` 新增 AST 节点文件。
2. 实现 holder 接口与监听器 enter/exit 处理逻辑。
3. 在以下位置注册新的 accept 路径：
`exedesc_stmt.go`、`leaf_snippet.go`，以及必要的父节点文件。

## 4. 串接运行时

1. 在 `internal/engine/processor.go` 的 `createProcessor` 增加映射。
2. 实现 prepare/execute 行为及错误语义。
3. 若指令需要运行时状态，在 `internal/engine/context.go` 增加对应方法。

## 5. 新增/更新测试

1. 解析器测试：`internal/lang/iparser/*`
2. 运行时测试：`internal/engine/*`
3. 同时覆盖正向与反向用例。

## 6. 同步文档

1. 更新 `README.md` 的特性说明与示例。
2. 如果你的环境跟踪扩展仓，记得同步 `vscode-gorch-extension/` 的关键字/snippet。

## 7. 验证

```bash
go test ./internal/lang/iparser
go test ./internal/engine
go test ./...
```
