package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// === ЗДЕСЬ БУДЕТ ДОБАВЛЕНИЕ ПОЛЬЗОВАТЕЛЯ В POSTGRES ===

type Postgres struct {
	conn *pgx.Conn
	ctx  context.Context
}

func New(ctx context.Context, dbConn string) (*Postgres, error) {

	conn, err := pgx.Connect(ctx, dbConn)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		conn: conn,
		ctx:  ctx,
	}, nil
}

func (p *Postgres) Close() error {
	return p.conn.Close(p.ctx)
}
