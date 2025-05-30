# 🧮 CLI Calculator (Go)

A simple command-line calculator written in Go. It supports basic arithmetic operations: **add**, **subtract**, **multiply**, and **divide**. Each operation is implemented modularly in its own file.

---

## 📁 Project Structure

go-cli-apps/
├── go.mod
├── main.go
└── calculator/
    ├── add.go
    ├── subtract.go
    ├── multiply.go
    └── divide.go

## ⚙️ Setup Instructions

1. Clone or download the project

```bash
git clone https://github.com/your-username/go-cli-apps.git
cd go-cli-apps
```

2. Initialize GO Modules

```bash
go mod tidy
```

## ▶️ Usage

```bash
go run main.go <operation> <nunber1> <nunber2>
```

# ✅ Supported operations

| Operation | Command Example          | Output       |
| --------- | ------------------------ | ------------ |
| `add`     | `go run main.go add 5 3` | `Result: 8`  |
| `sub`     | `go run main.go sub 9 4` | `Result: 5`  |
| `multi`     | `go run main.go mul 6 7` | `Result: 42` |
| `div`     | `go run main.go div 8 2` | `Result: 4`  |
| `mod`     |   `go run main.go mod 10 2` | `Result: 0` |

Division by zero is gracefully handled

```bash
go run main.go div 8 0
# Output: Error: division by zero
```

## 📦 Modular Design

Each operation is implemented in its own Go file inside the calculator/ package:

```go
// calculator/add.go
package calculator

func Add(a, b float64) float64 {
    return a + b
}
```

## 🚧 Error Handling

- Prints usage instructions if incorrect argunents are passed.
- Returns error messages for invalid numbers or operatuins.
- Handles division by zero wirh a user-friendly error.


## 🔮 Future Enhancements

- Use commandline-flags for more flexibler input.
- Add more operations(power, factorial, square root).
- Add unit tests.
- Add support for interactive mode or expression parsing.
