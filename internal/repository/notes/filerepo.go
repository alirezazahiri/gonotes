package notesrepo

import (
	"fmt"
	"strconv"

	"github.com/alirezazahiri/gonotes/internal/entities/note"
	jsonutils "github.com/alirezazahiri/gonotes/pkg/json"
)

type NotesFileRepo struct {
	filePath string
}

func NewFileRepo(filePath string) *NotesFileRepo {
	return &NotesFileRepo{
		filePath: filePath,
	}
}

func (n NotesFileRepo) Save(note note.Note) error {
	notes, err := n.GetAll()

	if err != nil {
		return err
	}

	notes[note.ID] = note

	return jsonutils.WriteJSONFile(notes, n.filePath)
}

func (n NotesFileRepo) Delete(id int64) error {
	notes, err := n.GetAll()

	if err != nil {
		return err
	}

	_, ok := notes[id]
	if !ok {
		return fmt.Errorf("could not found note with id %d", id)
	}

	delete(notes, id)

	return jsonutils.WriteJSONFile(notes, n.filePath)
}

func (n NotesFileRepo) Update(newNote note.Note) error {
	notes, err := n.GetAll()

	if err != nil {
		return err
	}

	notes[newNote.ID] = newNote

	err = jsonutils.WriteJSONFile(notes, n.filePath)
	if err != nil {
		return err
	}

	return err
}

func (n NotesFileRepo) GetAll() (map[int64]note.Note, error) {
	notes, err := jsonutils.ReadJSONFile(n.filePath)
	if err != nil {
		return nil, err
	}

	notesMap, ok := notes.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid notes file format")
	}

	noteStructsMap := make(map[int64]note.Note)

	for idStr, rawNote := range notesMap {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid note id received from file")
		}

		noteObj, ok := rawNote.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid notes file format")
		}

		text, ok := noteObj["text"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid note text received from file")
		}

		noteStructsMap[id] = note.Note{
			ID:   id,
			Text: text,
		}
	}

	return noteStructsMap, nil
}
