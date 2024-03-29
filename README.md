glox-interpreter
- 

This is my attempt to implement the interpreter from the book [Crafting Interpreters](https://craftinginterpreters.com/) in Go.

## Usage

```go run main.go examples/sample.lox```

## Progress
- [x] Part 1: Welcome
- [x] Chapter 1: Introduction
- [x] Chapter 2: A Map of the Territory
- [x] Chapter 3: The Lox Language

- [ ] Part 2: A Tree-Walk Interpreter
- [x] Chapter 4: Scanning. 
The scanner package reads source code and produces a stream of tokens.

- [x] Chapter 5: Representing Code.
The grammar defines a set of rules to convert a stream of tokens into a tree representation, since grammar is recursive.
We need to describe the nodes in the tree(primitives) in golang code.
These primitives are Expression, which can be Binary, Unary, Grouped or the actual Literal value.
This code representation should not have behaviour attached to it. Hence we use a Visitor pattern.


The ast-codegen package takes types that describe 
The ast package represents the syntax tree of the source code.

- [x] Chapter 6: Parsing Expression. 
The parser package parses the stream of tokens into an abstract syntax tree.
GRAMMAR.md describes the grammar of the Lox language.
Recursive descent parser (used by GCC, V8) is a top-down parser.
Each rule in the grammar is a function in the parser package.

- [x] Chapter 7: Evaluating Expressions.
Executing the syntax tree of simple expressions.
Using Java Object/Boxes types to hold Lox values.
Translate to Java operations.
The JVM takes care of memory management.

