package main

func main () {
    todos := Todos{}
    
    storage := StoreFile[Todos]("todo.json")
    storage.Read(&todos)

    cmdFlags := NewCmdFlags()
    cmdFlags.Execute(&todos, storage)

    storage.Save(todos)
}
