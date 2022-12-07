package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"university/domain/dto"
	"university/infrastructure/persistence"
	"university/infrastructure/repository"
)

type ServiceImpl struct {
	studentsRepository repository.StudentsRepository
}

func InitServiceImpl() *ServiceImpl {
	dbOrdersHelper, err := persistence.InitOrdersHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ServiceImpl{
		studentsRepository: dbOrdersHelper.StudentsRepository,
	}
}

func (cd *ServiceImpl) FindCurrentOrdersDetailByIdUser() (dto.Response, []dto.Students) {
	studentsNotes := []dto.Students{}

	student, _ := cd.studentsRepository.GetStudentWithNotes()

	for _, s := range *student {
		for i, n := range studentsNotes {
			if s.Email == n.Email {
				studentsNotes[i].Nota = append(studentsNotes[i].Nota, s.Nota)
				studentsNotes[i].Average = CalulateAverage(studentsNotes[i].Nota)
				break
			}

			if len(studentsNotes) == i+1 {
				studentsNotes = append(studentsNotes, dto.Students{
					Id:       s.Id,
					Email:    s.Email,
					FullName: s.FullName,
					Nota:     []string{s.Nota},
				})
			}
		}

		if len(studentsNotes) == 0 {
			studentsNotes = append(studentsNotes, dto.Students{
				Id:       s.Id,
				Email:    s.Email,
				FullName: s.FullName,
				Nota:     []string{s.Nota},
			})
		}
	}
	return dto.Response{}, studentsNotes
}

func CalulateAverage(notes []string) string {
	var xs []float64

	for _, n := range notes {
		if s, err := strconv.ParseFloat(n, 64); err == nil {
			xs = append(xs, s)
		}
	}
	return average(xs)
}

func average(xs []float64) string {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return fmt.Sprintf("%v", (total / float64(len(xs))))
}

func (cd *ServiceImpl) CreateNote(note dto.CreateNote, idStudent string) dto.Response {

	_ = cd.studentsRepository.CreateNote(note, idStudent)

	return dto.Response{
		Status: http.StatusOK,
	}
}

func (cd *ServiceImpl) DeleteNote(note dto.CreateNote, idStudent string) dto.Response {

	_ = cd.studentsRepository.DeleteNote(note, idStudent)

	return dto.Response{
		Status: http.StatusOK,
	}
}
