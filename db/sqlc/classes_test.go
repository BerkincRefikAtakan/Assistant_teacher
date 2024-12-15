package db

import (
	"assistant_teacher/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomClass(t *testing.T) Class {
	teacher1 := createRandomTeacher(t)
	args := CreateClassParams{
		TeacherID: teacher1.ID,
		ClassName: util.RandomString(5),
	}

	class, err := testQueries.CreateClass(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, class)

	require.Equal(t, args.TeacherID, class.TeacherID)
	require.Equal(t, args.ClassName, class.ClassName)

	require.NotZero(t, class.ID)
	require.NotZero(t, class.TeacherID)
	require.NotZero(t, class.CreatedAt)

	return class
}

func TestCreateClass(t *testing.T) {
	createRandomClass(t)
}

func TestGetClass(t *testing.T) {
	class1 := createRandomClass(t)

	args := GetClassesParams{
		TeacherID: class1.TeacherID,
		ClassName: class1.ClassName,
	}

	class2, err := testQueries.GetClasses(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, class2)

	require.Equal(t, class1.TeacherID, class2.TeacherID)
	require.Equal(t, class1.ClassName, class2.ClassName)

	require.WithinDuration(t, class1.CreatedAt.Time, class2.CreatedAt.Time, time.Second)
}

func TestDeleteClass(t *testing.T) {
	Class1 := createRandomClass(t)

	err := testQueries.DeleteClass(context.Background(), Class1.ID)
	require.NoError(t, err)

	args := GetClassesParams{
		ClassName: Class1.ClassName,
		TeacherID: Class1.TeacherID,
	}
	Class2, err := testQueries.GetClasses(context.Background(), args)
	require.Error(t, err)

	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, Class2)

}

func TestUpdateClasses(t *testing.T) {
	class1 := createRandomClass(t)
	teacher1 := createRandomTeacher(t)
	args := UpdateClassParams{
		TeacherID: teacher1.ID,
		ID:        class1.ID,
	}

	class2, err := testQueries.UpdateClass(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, class2)

	require.Equal(t, args.TeacherID, class2.TeacherID)
	require.Equal(t, args.ID, class2.ID)

	require.WithinDuration(t, class1.CreatedAt.Time, class2.CreatedAt.Time, time.Second)
}
