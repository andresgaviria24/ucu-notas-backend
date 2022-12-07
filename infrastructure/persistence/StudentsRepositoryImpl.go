package persistence

import (
	"university/domain/dto"
	"university/domain/entity"

	"gorm.io/gorm"
)

type StudentsRepositoryImpl struct {
	db *gorm.DB
}

func (s *StudentsRepositoryImpl) GetStudentWithNotes() (students *[]entity.Students, err error) {

	if err = s.db.Raw("select u.email ,u.full_name ,sc.nota,u.id  from users u join student_class sc on  u.id  = sc.id_student").Scan(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentsRepositoryImpl) CreateNote(note dto.CreateNote, student string) error {
	if err := s.db.Exec("INSERT into student_class (year,semester,id_student,id_class,nota) values (2022,1,?,1,?)", student, note.Note).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentsRepositoryImpl) DeleteNote(note dto.CreateNote, student string) error {
	if err := s.db.Exec("DELETE from student_class where id_student = ? and nota = ? limit 1", student, note.Note).Error; err != nil {
		return err
	}
	return nil
}
