This is an interpreter for a programming language called "Monkey", following
the book by Thorsten Ball: *Writing an Interpreter in Go*. 

Modifications to the original code from the book include (so far):

1. Support for unicode

Modification Ideas list:

1. Support for floating point numbers
2. Support for hex/octal notation of numbers

## Direnv

- install direnv
- make sure you add the line to the end of your .bashrc (hard tab required, that is NOT a space)
- add a .envrc to this directory with the line `layout go`
- authorize direnv with `direnv allow .`
- run tests such as `go test ./lexer` from the `src/monkey/` directory