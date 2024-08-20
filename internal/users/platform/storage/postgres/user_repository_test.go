package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage"
	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UserRepository_Save_RepositoryError(t *testing.T) {
	userID, userName := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test User"
	user, err := usersdomain.NewUser(userID, userName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name) VALUES ($1, $2)").
		WithArgs(userID, userName).
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(&connection, &config)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_UserRepository_Save_Succeed(t *testing.T) {
	userID, userName := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test User"

	user, err := usersdomain.NewUser(userID, userName)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO users (id, name) VALUES ($1, $2)").
		WithArgs(userID, userName).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(&connection, &config)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_UserRepository_Exists_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT 1 FROM users WHERE id = $1").
		WithArgs(id).
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(&connection, &config)

	userID, err := usersdomain.NewUserID(id)
	require.NoError(t, err)
	exists, err := repo.Exists(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.False(t, exists)
}

func Test_UserRepository_Exists_Succeed(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT 1 FROM users WHERE id = $1").
		WithArgs(id).
		WillReturnRows(sqlMock.NewRows([]string{"1"}).AddRow(1))

	repo := NewUserRepository(&connection, &config)

	userID, err := usersdomain.NewUserID(id)
	require.NoError(t, err)
	exists, err := repo.Exists(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.True(t, exists)
}

func Test_UserRepository_Get_RepositoryError(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT id, name FROM users WHERE id = $1").
		WithArgs(id).
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(&connection, &config)

	userID, err := usersdomain.NewUserID(id)
	require.NoError(t, err)
	user, err := repo.Get(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Nil(t, user)
}

func Test_UserRepository_Get_NotFound(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT id, name FROM users WHERE id = $1").
		WithArgs(id).
		WillReturnRows(sqlMock.NewRows([]string{"id", "name"}))

	repo := NewUserRepository(&connection, &config)

	userID, err := usersdomain.NewUserID(id)
	require.NoError(t, err)
	user, err := repo.Get(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Nil(t, user)
}

func Test_UserRepository_Get_Succeed(t *testing.T) {
	id := "37a0f027-15e6-47cc-a5d2-64183281087e"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	connection := storage.Connection{
		Db: db,
	}
	config := storage.Dbconfig{
		Timeout: 1 * time.Millisecond,
	}
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT id, name FROM users WHERE id = $1").
		WithArgs(id).
		WillReturnRows(sqlMock.NewRows([]string{"id", "name"}).AddRow(id, "Test User"))

	repo := NewUserRepository(&connection, &config)

	userID, err := usersdomain.NewUserID(id)
	require.NoError(t, err)
	user, err := repo.Get(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
