package main

import (
	"fmt"

	notescli "github.com/alirezazahiri/gonotes/internal/cli/notes"
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

	notesCLI := notescli.NewNotesCLI(notesManager)
	err = notesCLI.Run()
	if err != nil {
		fmt.Println("error running notes CLI:", err)
		return
	}
}
