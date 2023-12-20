package main

import (
	"Dictionnaire/dictionary"
	"fmt"
)


func main() {

    d := dictionary.New("testJsonFile")

    d.Add("estiam", "école")
    d.Add("iles", "étudiant")
    d.Add("aziz", "prof")

    fmt.Println(d.Get("estiam"))
    
    d.Remove("estiam")

    d.List()
    
}
