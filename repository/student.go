package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	err := s.db.Where("id = ?", id).Updates(student).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	var student model.Student
	err := s.db.Where("id = ?", id).Delete(&student).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	// Function ini akan mengambil semua data mahasiswa yang ada di dalam tabel `students` dan table `classes` dengan melakukan JOIN pada kedua tabel tersebut. Kemudian, data tersebut akan di-scan dan dimasukkan ke dalam slice `[]model.StudentClass`.
	var studentClasses []model.StudentClass
	result := s.db.Table("students").
		Select("students.name, students.address, classes.id, classes.name as class_name, classes.professor, classes.room_number").
		Joins("JOIN classes ON students.class_id = classes.id").
		Scan(&studentClasses)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(studentClasses) == 0 {
		return &[]model.StudentClass{}, nil
	}
	return &studentClasses, nil
}
