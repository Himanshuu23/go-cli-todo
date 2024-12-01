package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Cmdflags struct {
    Add    string
    Del    int
    Edit   string
    Toggle int
    List   bool
}

func NewCmdFlags() *Cmdflags {
    cf := Cmdflags{}

    flag.StringVar(&cf.Add, "add", "", "Add a new todo by specifying title")
    flag.StringVar(&cf.Edit, "edit", "", "Edit a new todo by index & specifying a title")
    flag.IntVar(&cf.Del, "delete", -1, "Specify todo's index to delete")
    flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo's index to toggle its completion state")
    flag.BoolVar(&cf.List, "list", false, "To list all todos")

    flag.Parse()

    return &cf
}

func (cf *Cmdflags) Execute(todos *Todos, storage *Storage[Todos]) {
    switch {
    case cf.List:
        todos.print()
    case cf.Add != "":
        todos.add(cf.Add)
        storage.Save(*todos) // Save after adding
    case cf.Edit != "":
        parts := strings.SplitN(cf.Edit, ":", 2)
        if len(parts) != 2 {
            fmt.Println("Error: Invalid format for edit. Please use index:new_title format")
            os.Exit(1)
        }
        index, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Error: invalid index for edit")
            os.Exit(1)
        }
        todos.edit(index, parts[1])
        storage.Save(*todos) // Save after editing
    case cf.Toggle != -1:
        todos.toggle(cf.Toggle)
        storage.Save(*todos) // Save after toggling
    case cf.Del != -1:
        todos.delete(cf.Del)
        storage.Save(*todos) // Save after deleting
    default:
        fmt.Println("Invalid command")
    }
}

