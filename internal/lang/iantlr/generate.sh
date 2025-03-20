#!/bin/sh

# alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
rm -rf alr
antlr4 -Dlanguage=Go -no-visitor -package alr gorchlang.g4 -o alr