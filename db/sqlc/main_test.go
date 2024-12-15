package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://root:secret@localhost:5433/assistant_teacher?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot create pool:", err)
	}

	// Bağlantının sağlıklı olduğunu doğrulamak için ping atın
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("cannot ping db:", err)
	}

	testDB = pool
	testQueries = New(testDB)

	code := m.Run()

	pool.Close()
	os.Exit(code)
}
