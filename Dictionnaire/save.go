package save

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dictionary struct {
    filename string
}

func New(filename string) *Dictionary {
    return &Dictionary{
        filename: filename,
    }
}

func (d *Dictionary) Add(word, definition string) {
    file, err := os.OpenFile(d.filename, os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return
    }

    data, err := json.Marshal(map[string]string{
        word: definition,
    })
    if err != nil {
        return
    }

    _, err = file.Write(data)
    if err != nil {
        return
    }

    file.Close()
}

func (d *Dictionary) Get(word string) string {
    data, err := os.ReadFile(d.filename)
    if err != nil {
        return ""
    }

    var entries map[string]string
    err = json.Unmarshal(data, &entries)
    if err != nil {
        return ""
    }

    return entries[word]
}

func (d *Dictionary) Remove(word string) {
    entries := make(map[string]string)

    data, err := os.ReadFile(d.filename)
    if err != nil {
        return
    }

    err = json.Unmarshal(data, &entries)
    if err != nil {
        return
    }

    delete(entries, word)

    data, err = json.Marshal(entries)
    if err != nil {
        return
    }

    err = os.WriteFile(d.filename, data, 0644)
    if err != nil {
        return
    }
}

func (d *Dictionary) List() {
    data, err := os.ReadFile(d.filename)
    if err != nil {
        return
    }

    var entries map[string]string
    err = json.Unmarshal(data, &entries)
    if err != nil {
        return
    }

    for word, definition := range entries {
        fmt.Println(word, definition)
    }
}
