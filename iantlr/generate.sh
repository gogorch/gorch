#!/bin/sh

# 检查是否传入了参数
if [ -z "$1" ]; then
    echo "Usage: $0 [go|typescript]"
    exit 1
fi

# antlr4 -Dlanguage=Javascript -no-visitor -package alr gorchlang.g4 -o alrjs

# 定义 ANTLR 工具的别名
# 请确保你的环境中有 antlr-4.13.0-complete.jar 或者正确配置了 ANTLR4
# ANTLR4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
ANTLR4=antlr4

if [ "$1" = "go" ]; then
    echo "Generating Go parser in internal/lang/iantlr..."
    # 定义Go生成的输出目录和包名
    OUTPUT_DIR="../internal/lang/iantlr"
    PACKAGE="iantlr"
    
    # 清理旧的生成文件
    rm -rf ${OUTPUT_DIR}/*.go ${OUTPUT_DIR}/*.tokens ${OUTPUT_DIR}/*.interp
    
    # 运行 ANTLR4 生成 Go 代码
    ${ANTLR4} -Dlanguage=Go -no-visitor -package ${PACKAGE} gorchlang.g4 -o ${OUTPUT_DIR}

elif [ "$1" = "typescript" ]; then
    echo "Generating TypeScript parser in vscode-gorch-extension/src/lang..."
    # 定义TypeScript生成的输出目录
    OUTPUT_DIR="../vscode-gorch-extension/src/lang"
    
    # 清理旧的生成文件
    rm -rf ${OUTPUT_DIR}/*.ts ${OUTPUT_DIR}/*.tokens ${OUTPUT_DIR}/*.interp
    
    # 运行 ANTLR4 生成 TypeScript 代码
    ${ANTLR4} -Dlanguage=TypeScript -no-visitor gorchlang.g4 -o ${OUTPUT_DIR}

else
    echo "Invalid argument: $1"
    echo "Usage: $0 [go|typescript]"
    exit 1
fi

echo "Done."