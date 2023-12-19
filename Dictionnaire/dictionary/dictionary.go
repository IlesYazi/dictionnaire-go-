package dictionary

import (
	"fmt"
)

type Dictionary struct {
    entries map[string]string
}

func New()Dictionary{
    entries := make(map[string]string)
    return Dictionary{entries}
}


func (d *Dictionary) Add(word, definition string) {
    d.entries[word] = definition
}

func (d *Dictionary) Get(word string) string {
    return d.entries[word]
}

func (d *Dictionary) Remove(word string) {
    delete(d.entries, word)
}

func (d *Dictionary) List() {
    for word, definition := range d.entries {
        fmt.Printf("%s: %s\n", word, definition)
    }
}