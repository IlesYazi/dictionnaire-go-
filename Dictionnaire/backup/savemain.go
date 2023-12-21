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
    d.List()

    fmt.Println("**** Definition de estiam :")
    fmt.Println(d.Get("estiam"))

    fmt.Println("**** Supprimer estiam :")
    d.Remove("estiam")
    d.List()
    
}


