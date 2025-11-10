package notes

import (
	"fmt"

	note "github.com/alirezazahiri/gonotes/internal/entities/note"
	notesrepo "github.com/alirezazahiri/gonotes/internal/repository/notes"
	idutils "github.com/alirezazahiri/gonotes/pkg/id"
)

type NoteManager struct {
	Notes map[int64]*note.Note `json:"notes"`
	repo  notesrepo.INoteRepository
}

func NewNoteManager(repo notesrepo.INoteRepository) *NoteManager {
	return &NoteManager{
		Notes: make(map[int64]*note.Note),
		repo:  repo,
	}
}

func (nm *NoteManager) CreateNote(text string) (*note.Note, error) {
	note := &note.Note{
		ID:   idutils.GenerateIDInt64(),
		Text: text,
	}

	err := nm.repo.Save(*note)

	if err != nil {
		return nil, err
	}

	nm.Notes[note.ID] = note

	return note, nil
}

func (nm *NoteManager) GetNoteById(noteId int64) (*note.Note, error) {
	note, exists := nm.Notes[noteId]

	if exists {
		return note, nil
	}

	return nil, fmt.Errorf("could not found note with id %d", noteId)
}

func (nm *NoteManager) UpdateNote(newNote note.Note) (*note.Note, error) {
	note, exists := nm.Notes[newNote.ID]

	if exists {
		err := nm.repo.Update(newNote)

		if err != nil {
			return nil, err
		}

		note.Text = newNote.Text

		return note, nil
	}

	return nil, fmt.Errorf("note with the given id '%d' does NOT exist", newNote.ID)
}

func (nm *NoteManager) DeleteNote(noteId int64) error {
	_, exists := nm.Notes[noteId]

	if !exists {
		return fmt.Errorf("could not found note with id %d", noteId)
	}

	err := nm.repo.Delete(noteId)

	if err != nil {
		return err
	}

	delete(nm.Notes, noteId)

	return nil
}

func (nm *NoteManager) LoadStoredNotes() error {
	notes, err := nm.repo.GetAll()

	if err != nil {
		return err
	}

	notesMap := make(map[int64]*note.Note)
	for _, note := range notes {
		notesMap[note.ID] = &note
	}

	nm.Notes = notesMap

	return nil
}
