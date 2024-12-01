# CLI Todo Application

A simple command-line todo application built using GoLang. Below are the commands and their usage:

![CLI Todo App Screenshot](path/to/your/image.png)

## Commands

| Command                 | Description                                           | Example                        |
|-------------------------|-------------------------------------------------------|--------------------------------|
| `-help`                | Displays all the available commands.                  | `./todo -help`                |
| `-list`                | Prints all the todos in a table format.               | `./todo -list`                |
| `-add "todo_title"`    | Adds a new todo with the specified title.             | `./todo -add "Buy groceries"` |
| `-toggle index`        | Toggles the completion status of the todo at the index.| `./todo -toggle 2`            |
| `-delete index`        | Deletes the todo at the specified index.              | `./todo -delete 3`            |
| `-edit index "new_title"` | Edits the title of the todo at the specified index.  | `./todo -edit 1 "Go for a walk"` |
