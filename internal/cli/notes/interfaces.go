package notescli

type INotesCLI interface {
	PrintNotesList() error
	CreateNote(text string) error
	UpdateNote(id int64, text string) error
	DeleteNote(id int64) error
	GetNoteById(id int64) error
}