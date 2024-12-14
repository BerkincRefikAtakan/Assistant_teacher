// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateClass(ctx context.Context, arg CreateClassParams) (Class, error)
	CreateParagraph(ctx context.Context, arg CreateParagraphParams) (Paragraph, error)
	CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error)
	CreateTeacher(ctx context.Context, arg CreateTeacherParams) (Teacher, error)
	DeleteClass(ctx context.Context, id int64) error
	DeleteParagraph(ctx context.Context, id int64) error
	DeleteStudent(ctx context.Context, id int64) error
	DeleteTeacher(ctx context.Context, id int64) error
	GetClassByID(ctx context.Context, id int64) (Class, error)
	GetClasses(ctx context.Context) ([]Class, error)
	GetClassesByTeacher(ctx context.Context, teacherID int64) ([]Class, error)
	GetParagraphByID(ctx context.Context, id int64) (Paragraph, error)
	GetParagraphs(ctx context.Context) ([]Paragraph, error)
	GetParagraphsByTeacher(ctx context.Context, teacherID int64) ([]Paragraph, error)
	GetStudentByID(ctx context.Context, id int64) (Student, error)
	GetStudents(ctx context.Context) ([]Student, error)
	GetStudentsByClass(ctx context.Context, classID int64) ([]Student, error)
	GetTeacherByID(ctx context.Context, id int64) (Teacher, error)
	GetTeachers(ctx context.Context) ([]Teacher, error)
	UpdateClass(ctx context.Context, arg UpdateClassParams) (Class, error)
	UpdateParagraph(ctx context.Context, arg UpdateParagraphParams) (Paragraph, error)
	UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error)
	UpdateTeacher(ctx context.Context, arg UpdateTeacherParams) (Teacher, error)
}

var _ Querier = (*Queries)(nil)