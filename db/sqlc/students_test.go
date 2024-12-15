package db

import (
	"assistant_teacher/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	class := createRandomClass(t)
	args := CreateStudentParams{
		ClassID: class.ID,
		Name:    util.RandomOwner(),
		Surname: util.RandomString(7),
	}

	student, err := testQueries.CreateStudent(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, args.Name, student.Name)
	require.Equal(t, args.Surname, student.Surname)

	require.NotZero(t, student.ID)
	require.NotZero(t, student.CreatedAt)

	return student
}
func createRandomStudentsInSameClass(t *testing.T, classID int64) Student {

	args := CreateStudentParams{
		ClassID: classID,
		Name:    util.RandomOwner(),
		Surname: util.RandomString(7),
	}

	student, err := testQueries.CreateStudent(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, args.Name, student.Name)
	require.Equal(t, args.Surname, student.Surname)

	require.NotZero(t, student.ID)
	require.NotZero(t, student.CreatedAt)

	return student
}
func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

func TestDeleteStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	err := testQueries.DeleteStudent(context.Background(), student1.ID)
	require.NoError(t, err)
	args := GetStudentsParams{
		Name:    student1.Name,
		Surname: student1.Surname,
		ClassID: student1.ClassID,
	}
	student2, err := testQueries.GetStudents(context.Background(), args)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, student2)

}

func TestGetStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	args := GetStudentsParams{
		Name:    student1.Name,
		Surname: student1.Surname,
		ClassID: student1.ClassID,
	}

	student2, err := testQueries.GetStudents(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.Name, student2.Name)
	require.Equal(t, student1.Surname, student2.Surname)
	require.WithinDuration(t, student1.CreatedAt.Time, student2.CreatedAt.Time, time.Second)
}

func TestUpdateStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	args := UpdateStudentParams{
		Name:    util.RandomString(6),
		Surname: util.RandomString(7),
		ID:      student1.ID,
		ClassID: student1.ClassID,
	}

	student2, err := testQueries.UpdateStudent(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, args.Name, student2.Name)
	require.Equal(t, args.Surname, student2.Surname)
	require.Equal(t, student1.ID, student2.ID)
	require.WithinDuration(t, student1.CreatedAt.Time, student2.CreatedAt.Time, time.Second)
}

func TestGetStudentsByClass(t *testing.T) {
	class := createRandomClass(t)

	var lastStudent Student
	for i := 0; i < 10; i++ {
		lastStudent = createRandomStudentsInSameClass(t, class.ID)
	}
	arg := GetStudentsByClassParams{
		ClassID: class.ID,
		Limit:   10,
		Offset:  0,
	}

	students, err := testQueries.GetStudentsByClass(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, students)

	for _, student := range students {
		require.NotEmpty(t, student)
		require.Equal(t, lastStudent.ClassID, student.ClassID)
	}
}
