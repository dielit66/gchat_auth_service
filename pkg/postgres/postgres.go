package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func New(user string, password string, host string, port string, dbname string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname))
	if err != nil {
		return nil, err
	}
	return conn, nil

}
