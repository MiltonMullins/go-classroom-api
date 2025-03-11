package repositories

import (
	"database/sql"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
)

type ClassroomRepository interface {
	GetClassrooms() ([]*models.Classroom, error)
	GetClassroomByID(classroomID int) (*models.Classroom, error)
	CreateClassroom(classroom *models.Classroom) error
	UpdateClassroom(classroomID int, classroom *models.Classroom) error
	DeleteClassroom(classroomID int) error
}

type classroomRepository struct {
	db *sql.DB
}

func NewClassroomRepository(db *sql.DB) ClassroomRepository {
	return &classroomRepository{
		db: db,
	}
}

func (r *classroomRepository) GetClassrooms() ([]*models.Classroom, error) {
	rows, err := r.db.Query("SELECT * FROM classrooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classrooms := []*models.Classroom{}
	for rows.Next() {
		classroom := &models.Classroom{}
		err := rows.Scan(&classroom.ID, &classroom.Name, &classroom.Description, &classroom.TeacherID)
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, classroom)
	}

	return classrooms, nil
}

func (r *classroomRepository) GetClassroomByID(classroomID int) (*models.Classroom, error) {
	row := r.db.QueryRow("SELECT * FROM classrooms WHERE id = $1", classroomID)
	classroom := &models.Classroom{}
	err := row.Scan(&classroom.ID, &classroom.Name, &classroom.Description, &classroom.TeacherID)
	if err != nil {
		return nil, err
	}
	return classroom, nil
}

func (r *classroomRepository) CreateClassroom(classroom *models.Classroom) error {
	_, err := r.db.Exec("INSERT INTO classrooms (name, description, teacher_id) VALUES ($1, $2, $3)", classroom.Name, classroom.Description, classroom.TeacherID)
	if err != nil {
		return err
	}

	return nil
}

func (r *classroomRepository) UpdateClassroom(classroomID int, classroom *models.Classroom) error {
	_, err := r.db.Exec("UPDATE classrooms SET name = $1, description = $2, teacher_id = $3 WHERE id = $4", classroom.Name, classroom.Description, classroom.TeacherID, classroomID)
	if err != nil {
		return err
	}

	return nil
}

func (r *classroomRepository) DeleteClassroom(classroomID int) error {
	_, err := r.db.Exec("DELETE FROM classrooms WHERE id = $1", classroomID)
	if err != nil {
		return err
	}

	return nil
}
