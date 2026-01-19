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

func New(conn *pgx.Conn, ctx context.Context) *Postgres {
	return &Postgres{
		conn: conn,
		ctx:  ctx,
	}
}

func (p *Postgres) Ping() error {
	return p.conn.Ping(p.ctx)
}
