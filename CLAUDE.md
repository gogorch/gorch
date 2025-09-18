# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gorch is a high-performance graph-based execution framework for complex business process orchestration, written in Go. It provides a Domain-Specific Language (DSL) for defining execution workflows and includes a complete toolchain with compiler and VS Code extension.

## Common Development Commands

### Build and Test
```bash
# Build all packages
go build -v ./...

# Run all tests with coverage
go test -coverprofile=coverage.txt -covermode=atomic ./...

# Run tests for specific package
go test ./internal/engine/...

# Build the Gorch compiler
go build -o gorchc ./gorchc
```

### Code Generation
```bash
# Generate all code (registration + imports) from DSL files
gorchc -g ./gen -d ./conf

# Generate only import code
gorchc -i ./gen -d ./conf

# Debug: print parsed DSL structure
gorchc -p -d ./conf
```

### VS Code Extension Development
```bash
# Navigate to extension directory
cd vscode-gorch-extension

# Install dependencies
npm install

# Compile extension
npm run compile

# Package extension
npm run package
```

## Architecture Overview

### Core Components
- **gorch.go**: Main public API entry point
- **internal/engine/**: Execution engine with graph traversal, operator scheduling, and runtime management
- **internal/lang/**: DSL parser built with ANTLR4 for syntax analysis and AST generation
- **internal/ort/**: Runtime type system for parameter handling and type validation
- **gorchc/**: Code generation compiler that produces operator registration code
- **pool/**: Goroutine and object pool management for performance optimization
- **mlog/**: Structured logging system with operator-level tracking
- **recorder/**: Execution recording and performance monitoring

### DSL Language Structure
The Gorch DSL uses three main execution blocks:
- **START**: Entry points for execution flows with unique names
- **FRAGMENT**: Reusable execution fragments that can be unfolded
- **REGISTER**: Operator registration center managing business logic operators

Execution patterns include serial (`->`), parallel (`[]`), conditional (`SWITCH/CASE`), and async (`GO`/`WAIT`) operations.

### Code Generation Workflow
1. Write business operators implementing the `Processor` interface
2. Create `.gorch` DSL files defining workflows in `/conf` directory
3. Run `gorchc -g ./gen -d ./conf` to generate registration code
4. Import generated code with `_ "github.com/gogorch/gorch/examples/gen"`
5. Parse DSL and execute workflows using the engine API

### Key Design Patterns
- **Operator States**: Fatal states stop execution, Info states log but continue
- **Context Sharing**: Zero-copy data passing via shared execution context
- **Async Execution**: GO/WAIT pattern with timeout and totalTimeout controls
- **Operator Wrapping**: Cross-cutting concerns via WRAP syntax
- **Skip Control**: Conditional execution with SKIP directive

### Performance Considerations
- Goroutine pooling prevents frequent creation/destruction overhead
- Object pooling reduces GC pressure for high-frequency operations
- Smart scheduling adapts concurrency based on system load
- Reference-based data passing minimizes memory allocations