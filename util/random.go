package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomClass generates a random Class name
func RandomClass() string {
	return RandomString(10)
}

// RandomId generates a random ID
func RandomId() int64 {
	return RandomInt(10, 100)
}

// RandomHeader generates a random header
func RandomHeader() pgtype.Text {
	return pgtype.Text{String: RandomString(10), Valid: true}
}

// RandomParagraph generates a random paragraph
func RandomParagraph() pgtype.Text {
	return pgtype.Text{String: RandomString(150), Valid: true}
}
