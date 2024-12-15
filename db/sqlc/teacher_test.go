package db

import (
	"assistant_teacher/util"
	"context"
	"errors"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

var ErrRecordNotFound = errors.New("no rows in result set")

func createRandomTeacher(t *testing.T) Teacher {
	args := CreateTeacherParams{
		Name:    util.RandomOwner(),
		Surname: util.RandomString(7),
	}

	teacher, err := testQueries.CreateTeacher(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, teacher)

	require.Equal(t, args.Name, teacher.Name)
	require.Equal(t, args.Surname, teacher.Surname)

	require.NotZero(t, teacher.ID)
	require.NotZero(t, teacher.CreatedAt)

	return teacher
}

func TestCreateTeacher(t *testing.T) {
	createRandomTeacher(t)
}

func TestGetTeacherByNameAndSurname(t *testing.T) {
	teacher1 := createRandomTeacher(t)

	args := GetTeacherByNameAndSurnameParams{
		Name:    teacher1.Name,
		Surname: teacher1.Surname,
	}

	teacher2, err := testQueries.GetTeacherByNameAndSurname(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	require.Equal(t, teacher1.Name, teacher2.Name)
	require.Equal(t, teacher1.Surname, teacher2.Surname)
	require.WithinDuration(t, teacher1.CreatedAt.Time, teacher2.CreatedAt.Time, time.Second)
}

func TestUpdateTeacher(t *testing.T) {
	teacher1 := createRandomTeacher(t)

	args := UpdateTeacherParams{
		Name:    util.RandomString(6),
		Surname: util.RandomString(7),
		ID:      teacher1.ID,
	}

	teacher2, err := testQueries.UpdateTeacher(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	require.Equal(t, args.Name, teacher2.Name)
	require.Equal(t, args.Surname, teacher2.Surname)
	require.Equal(t, teacher1.ID, teacher2.ID)
	require.WithinDuration(t, teacher1.CreatedAt.Time, teacher2.CreatedAt.Time, time.Second)
}

func TestDeleteTeacher(t *testing.T) {
	teacher1 := createRandomTeacher(t)
	err := testQueries.DeleteTeacher(context.Background(), teacher1.ID)
	require.NoError(t, err)
	args := GetTeacherByNameAndSurnameParams{
		Name:    teacher1.Name,
		Surname: teacher1.Surname,
	}
	teacher2, err := testQueries.GetTeacherByNameAndSurname(context.Background(), args)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, teacher2)

}
