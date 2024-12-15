package db

import (
	"assistant_teacher/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomParagraph(t *testing.T) Paragraph {
	teacher := createRandomTeacher(t)
	args := CreateParagraphParams{
		TeacherID: teacher.ID,
		Header:    util.RandomHeader(),
		Paragraph: util.RandomParagraph(),
	}
	Paragraph, err := testQueries.CreateParagraph(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, Paragraph)

	require.Equal(t, args.TeacherID, Paragraph.TeacherID)
	require.Equal(t, args.Header, Paragraph.Header)
	require.Equal(t, args.Paragraph, Paragraph.Paragraph)

	require.NotZero(t, Paragraph.ID)
	require.NotZero(t, Paragraph.CreatedAt)

	return Paragraph
}
func createRandomParagraphsCreatedBySameTeacher(t *testing.T, teacherID int64) Paragraph {

	args := CreateParagraphParams{
		TeacherID: teacherID,
		Header:    util.RandomHeader(),
		Paragraph: util.RandomParagraph(),
	}

	Paragraph, err := testQueries.CreateParagraph(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, Paragraph)

	require.Equal(t, args.TeacherID, Paragraph.TeacherID)
	require.Equal(t, args.Header, Paragraph.Header)
	require.Equal(t, args.Paragraph, Paragraph.Paragraph)

	require.NotZero(t, Paragraph.ID)
	require.NotZero(t, Paragraph.CreatedAt)

	return Paragraph
}
func TestCreateParagraph(t *testing.T) {
	createRandomParagraph(t)
}

func TestDeleteParagraph(t *testing.T) {
	paragraph1 := createRandomParagraph(t)
	arg := DeleteParagraphParams{
		TeacherID: paragraph1.TeacherID,
		Header:    paragraph1.Header,
	}
	err := testQueries.DeleteParagraph(context.Background(), arg)
	require.NoError(t, err)
	args := GetParagraphsParams{
		TeacherID: paragraph1.TeacherID,
		Header:    paragraph1.Header,
	}
	paragraph2, err := testQueries.GetParagraphs(context.Background(), args)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, paragraph2)

}

func TestGetParagraph(t *testing.T) {
	paragraph1 := createRandomParagraph(t)

	args := GetParagraphsParams{
		TeacherID: paragraph1.TeacherID,
		Header:    paragraph1.Header,
	}
	paragraph2, err := testQueries.GetParagraphs(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, paragraph2)

	require.Equal(t, paragraph1.TeacherID, paragraph2.TeacherID)
	require.Equal(t, paragraph1.Header, paragraph2.Header)
	require.Equal(t, paragraph1.ID, paragraph2.ID)
	require.WithinDuration(t, paragraph1.CreatedAt.Time, paragraph2.CreatedAt.Time, time.Second)
}

func TestUpdateParagraphOrAndHeader(t *testing.T) {
	paragraph1 := createRandomParagraph(t)
	args := UpdateParagraphOrAndHeaderParams{
		Header:    util.RandomHeader(),
		Paragraph: util.RandomParagraph(),
		ID:        paragraph1.ID,
	}
	paragraph2, err := testQueries.UpdateParagraphOrAndHeader(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, paragraph2)

	require.Equal(t, args.Header, paragraph2.Header)
	require.Equal(t, args.Paragraph, paragraph2.Paragraph)
	require.Equal(t, args.ID, paragraph2.ID)
	require.WithinDuration(t, paragraph1.CreatedAt.Time, paragraph2.CreatedAt.Time, time.Second)
}

func TestGetParagraphByTeacherID(t *testing.T) {
	teacher := createRandomTeacher(t)

	var lastParagraph Paragraph
	for i := 0; i < 10; i++ {
		lastParagraph = createRandomParagraphsCreatedBySameTeacher(t, teacher.ID)
	}
	arg := GetParagraphsByTeacherParams{
		TeacherID: teacher.ID,
		Limit:     10,
		Offset:    0,
	}

	Paragraphs, err := testQueries.GetParagraphsByTeacher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Paragraphs)

	for _, paragraph := range Paragraphs {
		require.NotEmpty(t, paragraph)
		require.Equal(t, lastParagraph.TeacherID, paragraph.TeacherID)
	}
}
