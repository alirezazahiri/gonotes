package notescli

import (
	"flag"
	"fmt"

	"github.com/alirezazahiri/gonotes/internal/entities/note"
	"github.com/alirezazahiri/gonotes/internal/notes"
)

type NotesCLI struct {
	noteManager *notes.NoteManager
}

func NewNotesCLI(noteManager *notes.NoteManager) *NotesCLI {
	return &NotesCLI{
		noteManager: noteManager,
	}
}

func (n *NotesCLI) Run() error {
	// --list
	listFlag := flag.Bool("list", false, "list all notes")
	// --create --text="note text"
	createFlag := flag.Bool("create", false, "create a new note")
	textFlag := flag.String("text", "", "note text")
	// --update --id=1 --text="note text"
	updateFlag := flag.Bool("update", false, "update an existing note")
	// shared flag for commands needing an id
	idFlag := flag.Int64("id", 0, "note id")
	// --delete --id=1
	deleteFlag := flag.Bool("delete", false, "delete a note")
	// --get --id=1
	getFlag := flag.Bool("get", false, "get a note")

	flag.Parse()

	switch {
	case *listFlag:
		n.PrintNotesList()
	case *createFlag && *textFlag != "":
		n.CreateNote(*textFlag)
	case *updateFlag && *idFlag != 0 && *textFlag != "":
		n.UpdateNote(*idFlag, *textFlag)
	case *deleteFlag && *idFlag != 0:
		n.DeleteNote(*idFlag)
	case *getFlag && *idFlag != 0:
		n.GetNoteById(*idFlag)
	}
	return nil
}

func (n *NotesCLI) PrintNotesList() error {
	notesMap := n.noteManager.Notes

	for id, note := range notesMap {
		fmt.Printf("%d: %s\n", id, note.Text)
	}

	return nil
}

func (n *NotesCLI) CreateNote(text string) error {
	note, err := n.noteManager.CreateNote(text)
	if err != nil {
		return err
	}

	fmt.Printf("Note created with id: %d\n", note.ID)
	return nil
}

func (n *NotesCLI) UpdateNote(id int64, text string) error {
	note := note.Note{
		ID:   id,
		Text: text,
	}
	_, err := n.noteManager.UpdateNote(note)
	if err != nil {
		return err
	}

	fmt.Printf("Note updated with id: %d\n", id)
	return nil
}

func (n *NotesCLI) DeleteNote(id int64) error {
	err := n.noteManager.DeleteNote(id)
	if err != nil {
		return err
	}

	fmt.Printf("Note deleted with id: %d\n", id)
	return nil
}

func (n *NotesCLI) GetNoteById(id int64) error {
	note, err := n.noteManager.GetNoteById(id)
	if err != nil {
		return err
	}

	fmt.Printf("Note with id: %d\n", note.ID)
	return nil
}
