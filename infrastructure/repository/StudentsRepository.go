package repository

import (
	"university/domain/dto"
	"university/domain/entity"
)

type StudentsRepository interface {
	GetStudentWithNotes() (students *[]entity.Students, err error)
	CreateNote(note dto.CreateNote, student string) error
	DeleteNote(note dto.CreateNote, student string) error
}
