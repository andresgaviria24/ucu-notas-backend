package service

import (
	"university/domain/dto"
)

type Service interface {
	FindCurrentOrdersDetailByIdUser() (dto.Response, []dto.Students)
	CreateNote(note dto.CreateNote, idStudent string) dto.Response
	DeleteNote(note dto.CreateNote, idStudent string) dto.Response
}
