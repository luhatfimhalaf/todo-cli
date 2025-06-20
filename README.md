# Todo CLI

A command-line interface (CLI) application built in Go to manage to-do tasks. This project allows users to add, edit, complete, delete, and list tasks with priorities, stored persistently in a JSON file.

## Features
- Add tasks with descriptions and priorities (1=high, 2=medium, 3=low)
- Edit task descriptions
- Set task priorities
- Mark tasks as completed
- Delete tasks
- List tasks (filter by completed or pending status)
- Persistent storage in `todos.json`

## Installation
1. Ensure Go is installed (version 1.16 or later).
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/todo-cli.git
   cd todo-cli
