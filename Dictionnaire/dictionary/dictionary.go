package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Dictionary struct {
    data []DictionaryEntry
}

type DictionaryEntry struct {
    Word string `json:"word"`
    Definition string `json:"definition"`
}

func NewDictionary() Dictionary {
    return Dictionary{
        data: []DictionaryEntry{},
    }
}

func (d *Dictionary) Add(word, definition string) {
    d.data = append(d.data, DictionaryEntry{
        Word: word,
        Definition: definition,
    })
}

func (d *Dictionary) Get(word string) string {
    for _, entry := range d.data {
        if entry.Word == word {
            return entry.Definition
        }
    }
    return ""
}

func (d *Dictionary) Remove(word string) {
    for i, entry := range d.data {
        if entry.Word == word {
            d.data[i] = d.data[len(d.data)-1]
            d.data = d.data[:len(d.data)-1]
            return
        }
    }
}

func (d *Dictionary) List() []string {
    words := make([]string, 0, len(d.data))
    for _, entry := range d.data {
        words = append(words, entry.Word)
    }
    return words
}

func (d *Dictionary) Save() error {
    data, err := json.Marshal(d.data)
    if err != nil {
        return err
    }
    err = os.WriteFile("dictionary.json", data, 0644)
    return err
}

func (d *Dictionary) Load() error {
    data, err := os.ReadFile("dictionary.json")
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &d.data)
    return err
}

func (d *Dictionary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        d.Add(r.FormValue("word"), r.FormValue("definition"))
        d.Save()
    } else if r.Method == http.MethodGet {
        if r.FormValue("word") != "" {
            w.Write([]byte(d.Get(r.FormValue("word"))))
        } else {
            w.Write([]byte(fmt.Sprintf("Liste des mots : %v", d.List())))
        }
    } else if r.Method == http.MethodDelete {
        d.Remove(r.FormValue("word"))
        d.Save()
    }
}