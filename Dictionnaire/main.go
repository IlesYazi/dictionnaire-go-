package main

import (
    "fmt"
)


func main() {

    entries := map[string]string{}

    entries["iles"] = "étudiant"
    entries["aziz"] = "prof"
    entries["estiam"] = "école"

    fmt.Println(entries["iles"])

    delete(entries, "estiam")

    for word, definition := range entries {
        fmt.Printf("%s: %s\n", word, definition)
    }
}
