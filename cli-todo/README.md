![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

# 📝 Go CLI To-Do App

A simple command-line to-do list application written in Go. Tasks are stored in a JSON file for persistence.

---

## 📦 Features

- Add new tasks
- List existing tasks
- Mark tasks as done
- JSON-based file persistence
- Easy-to-extend structure

---

## 🛠️ Requirements

- [Go](https://golang.org/doc/install) 1.18 or later

---

## 🚀 Usage

```bash
go run main.go <command> [args]
```

## 📃 Commands

| Operation | Command Example                               | Output Example                       |
|-----------|------------------------------------------------|--------------------------------------|
| `add`     | `go run main.go add "Buy groceries"`           | `Task added.`                        |
| `list`    | `go run main.go list`                          | `1. [ ] Buy groceries`               |
| `done`    | `go run main.go done 1`                        | `Task marked as done.`              |
| `list`    | `go run main.go list`                          | `1. [x] Buy groceries`               |


## 📁 File Structure

``` bash
todo/
├── main.go         # CLI entry point
├── task/
│   └── task.go     # Task struct, logic for add/list/done/load/save
├── todo.json       # Auto-generated task file (ignored by Git)
```

## 🧠 Example

``` bash
go run main.go add "Finish Go project"
go run main.go add "Review pull request"
go run main.go list
go run main.go done 1
go run main.go list
```

## 💾 Data Format (todo.json)

```bash
[
  {
    "description": "Finish Go project",
    "done": true
  },
  {
    "description": "Review pull request",
    "done": false
  }
]
```

## 📒 Notes
- todo.json file is created int he smae diectory automatically
- Data is saved in between sessuins
- You can easily build a binary using :

```bash
go build -o todo
./todo add "Plan vacation"
```

## 🪪 License
This project is licensed under the [MIT License](LICENSE).
