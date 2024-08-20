package postgres

const (
	sqlUserTable = "users"
)

type sqlUser struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
