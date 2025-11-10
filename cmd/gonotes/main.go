package main

import (
	"fmt"

	"github.com/alirezazahiri/gonotes/internal/notes"
	notesrepo "github.com/alirezazahiri/gonotes/internal/repository/notes"
)

const notesStorageFilePath = "data/notes.json"

func main() {
	notesRepository := notesrepo.NewFileRepo(notesStorageFilePath)
	notesManager := notes.NewNoteManager(*notesRepository)
	
	err := notesManager.LoadStoredNotes()
	if err != nil {
		fmt.Println("error loading stored notes:", err)
		return
	}
}
