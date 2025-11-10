package notesrepo

import (
	"github.com/alirezazahiri/gonotes/internal/entities/note"
)

type INoteRepository interface {
	Save(note note.Note) error
	Delete(id int64) error
	Update(newNote note.Note) error
	GetAll() (map[int64]note.Note, error)
}
