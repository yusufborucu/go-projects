# gopassgen

`gopassgen` is a simple and customizable CLI tool built with Go and Cobra for generating secure random passwords.

---

## Features

- Generate strong passwords with a single command
- Customize password length
- Optionally include:
  - Uppercase letters
  - Numbers
  - Symbols
- Fast and lightweight (only uses Go standard library + Cobra)

---

## Installation

```bash
git clone https://github.com/yusufborucu/go-projects/gopassgen.git
cd gopassgen
go build -o gopassgen
```

Or run without building:

```bash
go run main.go generate --length 16 --uppercase --numbers --symbols
```

---

## Usage

```bash
./gopassgen generate [flags]
```

### Available Flags

| Flag                 | Description                      | Example                    |
|----------------------|----------------------------------|----------------------------|
| `--length`, `-l`     | Length of the password           | `--length 20`              |
| `--uppercase`, `-u`  | Include uppercase letters        | `--uppercase`              |
| `--numbers`, `-n`    | Include numbers                  | `--numbers`                |
| `--symbols`, `-s`    | Include special characters       | `--symbols`                |

### Example

```bash
./gopassgen generate --length 16 --uppercase --numbers --symbols
```

Output:

```
Generated password: W9@fXc7$zN1#kUp2
```

---

## Project Structure

```
gopassgen/
├── cmd/
│   ├── root.go
│   └── generate.go
└── main.go
```

---

## Built With

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)

---

## License

MIT