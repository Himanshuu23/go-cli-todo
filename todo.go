package main

import (
    "fmt"
    "time"
    "errors"
    "strconv"
    "os"

    "github.com/aquasecurity/table"
)

type Todo struct {
    Title      string
    IsComplete bool
    CreatedAt  time.Time
    CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
    t := Todo{Title: title, IsComplete: false, CreatedAt: time.Now()}
    *todos = append(*todos, t)
    fmt.Println("New todo created successfully!")
}

func (todos *Todos) print() {
    table := table.New(os.Stdout)
    table.SetRowLines(false)
    table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
    
    for index, t := range *todos {
        completedAt := ""
        isComplete := "❌"
        
        if t.IsComplete {
            isComplete = "✅"
            if t.CompletedAt != nil {
                completedAt = t.CompletedAt.Format(time.RFC1123)
            }
        }
        
        table.AddRow(strconv.Itoa(index), t.Title, isComplete, t.CreatedAt.Format(time.RFC1123), completedAt)
    }

    table.Render()
}

func (todos *Todos) ValidateIndex(index int) error {
    if index < 0 || index >= len(*todos) {
        return errors.New("Invalid index")
    }
    
    return nil
}

func (todos *Todos) delete(index int) {
    err := todos.ValidateIndex(index)
    if err != nil {
        fmt.Println(err)
        return
    }

    *todos = append((*todos)[:index], (*todos)[index+1:]...)
    fmt.Println("Deleted successfully!")
}

func (todos *Todos) toggle(index int) {
    err := todos.ValidateIndex(index)
    if err != nil {
        fmt.Println(err)
        return
    }
   
    now := time.Now()
    (*todos)[index].IsComplete = !(*todos)[index].IsComplete
    (*todos)[index].CompletedAt = &now
    fmt.Println("Toggled successfully!")
}

func (todos *Todos) edit(index int, title string) {
    err := todos.ValidateIndex(index)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    (*todos)[index].Title = title
    fmt.Println("Successfully changed the title!")
}
