package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
    Word string
    Definition string
}

type Dictionary struct {
    filename string
    entries []*Entry
}

func New(filename string) *Dictionary {
    return &Dictionary{
        filename: filename,
        entries: []*Entry{},
    }
}

func (d *Dictionary) Add(word, definition string) {
    entry := &Entry{
        Word: word,
        Definition: definition,
    }

    d.entries = append(d.entries, entry)

    // Écrit le dictionnaire mis à jour dans le fichier (conversion objet en octets)
    data, err := json.Marshal(d.entries)
    if err != nil {
        fmt.Println("Erreur lors de la sérialisation des entrées :", err)
        return
    }

    err = os.WriteFile(d.filename, data, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'écriture du fichier :", err)
        return
    }
}

func (d *Dictionary) Get(word string) string {
    for _, entry := range d.entries {
        if entry.Word == word {
            return entry.Definition
        }
    }

    return ""
}

func (d *Dictionary) Remove(word string) {
    for i, entry := range d.entries {
        if entry.Word == word {
            // Supprime l'entrée
            d.entries = append(d.entries[:i], d.entries[i+1:]...)

            // Écrit le dictionnaire mis à jour dans le fichier
            data, err := json.Marshal(d.entries)
            if err != nil {
                fmt.Println("Erreur lors de la sérialisation des entrées :", err)
                return
            }

            err = os.WriteFile(d.filename, data, 0644)
            if err != nil {
                fmt.Println("Erreur lors de l'écriture du fichier :", err)
                return
            }

            return
        }
    }
}

func (d *Dictionary) List() {
    // Parcours les entrées du dictionnaire
    for _, entry := range d.entries {
        fmt.Println(entry.Word, entry.Definition)
    }
}
