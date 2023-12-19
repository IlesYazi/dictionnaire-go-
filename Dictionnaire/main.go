package main

import (
    "dictionary"
    "fmt"
)


func main() {

    entries := map[string]string{}

    entries["iles"] = "étudiant"
    entries["aziz"] = "prof"
    entries["estiam"] = "école"


    Remove(entries,"estiam")
	Get(entries,"iles")
	List(entries)

    for word, definition := range entries {
        fmt.Printf("%s: %s\n", word, definition)
    }
}
