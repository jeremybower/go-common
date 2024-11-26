package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeremybower/go-common/env"
	"github.com/stretchr/testify/require"
)

func databasePoolForTesting(t *testing.T) *pgxpool.Pool {
	// Parse the database url.
	url, err := url.Parse(env.Required("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}

	// Create the database.
	dbName := "test-" + uuid.NewString()
	exitCode := createDatabase(t, url, dbName)
	require.Equal(t, 0, exitCode)

	// Replace the path with a unique database name.
	url.Path = "/" + dbName

	// Create a new db.
	dbPool, err := pgxpool.New(context.Background(), url.String())
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}

	// Migrate the database.
	ctx := context.Background()
	dbPool.Exec(ctx, "CREATE TABLE values (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL, value TEXT NOT NULL);")

	// Success.
	return dbPool
}

func createDatabase(t *testing.T, url *url.URL, dbName string) int {
	// Execute psql command to create the database.
	cmd := exec.Command("psql", url.String(), "-c", fmt.Sprintf(`CREATE DATABASE "%s";`, dbName))
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	writer := &testWriter{t: t}
	cmd.Stderr = writer
	cmd.Stdout = writer

	// Set the working dir and run.
	cmd.Dir = "/workspace"
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}

		t.Fatal(err)
	}

	return 0
}

type testWriter struct {
	t *testing.T
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.t.Log(string(p))
	return len(p), nil
}
