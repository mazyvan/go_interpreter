# Persistio Scripting Language

**Persistio** is a simple, interpreted programming language implemented in Go. It is designed for learning, experimentation, and as a foundation for building your own language features. Persistio supports variables, functions (including closures), arrays, hashes (dictionaries), conditionals, and a small set of built-in functions.

---

## Features

- Lexical scoping
- Variadic functions
- Function literals
- Dynamic typing
- Built-in functions
- Functions as first-class citizens
- Anonymous functions
- Higher-order functions
- Closures
- Let bindings for variables
- First-class functions and closures
- Conditionals (`if`, `else`)
- Dot notation for accessing object properties
- Return statements
- Arrays and hashes
- String support and concatenation
- Built-in functions: `len`, `first`, `last`, `rest`, `push`, `puts`
- REPL for interactive exploration
- VS Code syntax highlighting extension

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer

### Build and Run

Clone the repository:

```sh
git clone https://github.com/mazyvan/persistio.git
cd persistio
```

Build and run the interpreter:

```sh
go run main.go
```

You will see a welcome message and a prompt:

```
Hello <yourname>! This is Ivan's programming language!
Feel free to type in commands
>>
```

Type Persistio code directly into the prompt!

---

## Language Syntax

### Variables

```persistio
let x = 5;
let y = 10;
let result = x + y;
```

### Functions

```persistio
let add = fn(a, b) {
  return a + b;
};
add(2, 3); // 5

// Closures
let newAdder = fn(x) {
  fn(y) { x + y };
};
let addTwo = newAdder(2);
addTwo(3); // 5
```

### Conditionals

```persistio
let x = 10;
if (x > 5) {
  return true;
} else {
  return false;
}
```

### Arrays

```persistio
let arr = [1, 2, 3];
arr[0]; // 1
arr[1 + 1]; // 3

let newArr = push(arr, 4); // [1, 2, 3, 4]
len(arr); // 3
first(arr); // 1
last(arr); // 3
rest(arr); // [2, 3]
```

### Hashes (Dictionaries)

```persistio
let myHash = {"name": "Ivan", "age": 30};
myHash["name"]; // "Ivan"
```

### Strings

```persistio
let greeting = "Hello, " + "world!";
len(greeting); // 13
```

---

## Built-in Functions

| Function | Description               | Example          | Result         |
| -------- | ------------------------- | ---------------- | -------------- |
| `len`    | Length of string or array | `len([1,2,3])`   | `3`            |
| `first`  | First element of array    | `first([1,2,3])` | `1`            |
| `last`   | Last element of array     | `last([1,2,3])`  | `3`            |
| `rest`   | All but first element     | `rest([1,2,3])`  | `[2,3]`        |
| `push`   | Add element to array      | `push([1,2], 3)` | `[1,2,3]`      |
| `puts`   | Print arguments           | `puts("Hello")`  | prints `Hello` |

---

## Examples

### Fibonacci

```persistio
let fib = fn(n) {
  if (n == 0) {
    return 0;
  } else {
    if (n == 1) {
      return 1;
    } else {
      return fib(n - 1) + fib(n - 2);
    }
  }
};
fib(10); // 55
```

### Factorial

```persistio
let fact = fn(n) {
  if (n == 0) {
    return 1;
  } else {
    return n * fact(n - 1);
  }
};
fact(5); // 120
```

---

## Development

### Running Tests

To run all tests:

```sh
go test ./...
```

### Project Structure

- `main.go`: Entry point, starts the REPL.
- `lexer`: Lexical analysis (tokenizer).
- `parser`: Parser for Persistio syntax.
- `ast`: Abstract Syntax Tree definitions.
- `object`: Runtime object system (integers, strings, arrays, etc).
- `evaluator`: Interpreter/evaluator.
- `repl`: Read-Eval-Print Loop.
- `examples`: Example Persistio scripts.
- `extensions/persistio`: VS Code extension for syntax highlighting.

---

## VS Code Extension

A VS Code extension for Persistio syntax highlighting is included in `extensions/persistio`.

To use:

1. Open the folder in VS Code.
2. Open the `extensions/persistio` folder.
3. Press `F5` to launch a new Extension Development Host window.
4. Open `.prs` files to see syntax highlighting.

---

## License

MIT License. See LICENSE for details.

---

Enjoy experimenting with **Persistio**! Contributions and suggestions are welcome.
