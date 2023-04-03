package notesd

import "github.com/VanOODev/education_web_server/models/storage"

type NotesD struct {
	storage storage.Storage
}

func (d *NotesD) Close() {
}

func (d *NotesD) Add(data int64) (index int64) {
	//TODO implement me
	panic("implement me")
}

func (d *NotesD) Get(index int64) (data int64) {
	//TODO implement me
	panic("implement me")
}

func (d *NotesD) Delete(index int64) {
	//TODO implement me
	panic("implement me")
}

func NewNotesD(storage storage.Storage) *NotesD {
	nd := &NotesD{
		storage: storage,
	}
	return nd
}
