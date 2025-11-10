package notes

import "github.com/alirezazahiri/gonotes/internal/entities/note"

type INoteManager interface {
	CreateNote(text string) (*note.Note, error)
	GetNoteById(id int64) (*note.Note, error)
	UpdateNote(newNote note.Note) (*note.Note, error)
	DeleteNote(id int64) error
}