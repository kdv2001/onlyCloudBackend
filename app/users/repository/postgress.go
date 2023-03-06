package repository

import "github.com/jackc/pgx/v5"

type postgresUsersRepository struct {
	connection *pgx.Conn
}

func NewPostgresUsersRepository(connection *pgx.Conn) *postgresUsersRepository {
	return &postgresUsersRepository{connection: connection}
}

func (uc *postgresUsersRepository) SignIn() error {
	return nil
}
