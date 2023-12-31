package main

import (
	"Dictionnaire/dictionary"
	"fmt"
	"net/http"
)

func main() {

    dict := dictionary.NewDictionary()

    dict.Add("estiam", "école")
    dict.Add("iles", "étudiant")
    dict.Add("aziz", "prof")

    fmt.Println("**** Définition du mot estiam :")
    fmt.Println(dict.Get("estiam"))

    // Supprime Estiam
    dict.Remove("estiam")

    fmt.Println("**** Liste des mots : ")
    for _, word := range dict.List() {
        fmt.Println(word) 
    }

    // Sauvegarder les données
    err := dict.Save()
    if err != nil {
        fmt.Println(err)
    }

    // Charger les données
    err = dict.Load()
    if err != nil {
        fmt.Println(err)
    }

    // Démarrer le serveur HTTP
    http.HandleFunc("/", dict.ServeHTTP)
    http.ListenAndServe(":8090", nil)

}
