package postgres

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage"
	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
)

type UserRepository struct {
	connection *storage.Connection
	dbconfig   *storage.Dbconfig
}

func NewUserRepository(connection *storage.Connection, dbconfig *storage.Dbconfig) *UserRepository {
	return &UserRepository{
		connection: connection,
		dbconfig:   dbconfig,
	}
}

func (r *UserRepository) Save(ctx context.Context, user usersdomain.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser)).For(sqlbuilder.PostgreSQL)
	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:   user.Id.String(),
		Name: user.Name.String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbconfig.Timeout)
	defer cancel()

	_, err := r.connection.Db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist user on database: %v", err)
	}

	return nil
}

func (r *UserRepository) Exists(ctx context.Context, id usersdomain.UserID) (bool, error) {
	sb := sqlbuilder.Select("1").From(sqlUserTable)
	sb.Where(sb.Equal("id", id.String()))
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbconfig.Timeout)
	defer cancel()

	rows, err := r.connection.Db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return false, fmt.Errorf("error trying to check if user exists on database: %v", err)
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (r *UserRepository) Get(ctx context.Context, id usersdomain.UserID) (*usersdomain.User, error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	sb := sqlbuilder.Select("id", "name").From(sqlUserTable)
	sb.Where(sb.Equal("id", id.String()))
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbconfig.Timeout)
	defer cancel()

	row := r.connection.Db.QueryRowContext(ctxTimeout, query, args...)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error trying to get user from database: %v", err)
	}
	user := new(sqlUser)
	err := row.Scan(userSQLStruct.Addr(user)...)
	if err != nil {
		return nil, nil
	}

	userVO, err := usersdomain.NewUser(user.ID, user.Name)
	return &userVO, err
}
