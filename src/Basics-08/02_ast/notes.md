# Abstract syntax tree

An AST (abstract syntax tree) is a tree-like data structure that completely describes 
the operation of a program written in any programming language.

Each node in this tree is part of the program. To see this try parsing  various Go programs
[here](http://goast.yuroyoro.net/).

---

## Lexer

```go
func sum(a, b int) int {
    return a + b
}
```

transformation into
```go
FUNC_DEF IDENT("sum") '(' IDENT("a") ',' IDENT('b') TYPE_INT ')' TYPE_INT '{'
RETURN IDENT('a') PLUS IDENT('b')
'}'
```

In the first stage, the compiler breaks the written program into small parts—lexemes—and 
assigns each lexeme its own pre-prepared identifier—a token. If a token isn't found for a 
lexeme, a compilation error will be returned. After all the tokens are generated, they are 
sent to the next stage of parsing.

---

## Parser

A parser is used in compilers to generate an AST from tokens and, possibly, perform some 
optimizations at this stage. But besides tokens, a parser lacks another important component 
for composing an AST: a grammar.

Using tokens and a language grammar, a compiler can generate an AST and also discard parts 
unnecessary for generating machine code. The point is that all these commas, parentheses, 
and arrows are only needed by us, programmers, to understand what we're writing.

For a compiler, this is all junk that would only complicate the generation of correct code. 
Therefore, a parser generates a tree in such a way as to eliminate all unnecessary parts of
a written program and retain only what is definitely useful. This is where the name AST comes
from.

## Useful links

* [Golang AST Visualizer](http://goast.yuroyoro.net/)
* [Understand go programm with go/parser](https://medium.com/justforfunc/understanding-go-programs-with-go-parser-c4e88a6edb87)
* [Basic AST Traversal in Go](https://zupzup.org/go-ast-traversal/)