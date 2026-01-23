package repository

import (
	"context"
	"time"

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

// ping postgres server
func (p *Postgres) Ping() error {
	return p.conn.Ping(p.ctx)
}

// add user`s identity in client table
func (p *Postgres) AddUser(username string, telegramID int64, expiresAt time.Time) error {
	sqlRow := `
	INSERT INTO client (telegram_username, telegram_id, status, expires_at)
	VALUES ($1, $2, false, $3);
	`
	_, err := p.conn.Exec(p.ctx, sqlRow, username, telegramID, expiresAt)
	return err
}

// close postgres server connection
func (p *Postgres) Close() error {
	return p.conn.Close(p.ctx)
}

// get hostID from client table where telegram_username = $1
func (p *Postgres) GetHostID(telegramID int64) (int, error) {
	sqlRow := `
	SELECT host_id FROM client WHERE telegram_id = $1;
	`
	var hostID int
	err := p.conn.QueryRow(p.ctx, sqlRow, telegramID).Scan(&hostID)
	if err != nil {
		return hostID, err
	}
	return hostID, nil
}

func (p *Postgres) UpdateStatusTrue(telegramID int64) error {
	sqlRow := `
	UPDATE client
	SET status = TRUE
	WHERE telegram_id = $1;
	`
	_, err := p.conn.Exec(p.ctx, sqlRow, telegramID)
	return err
}

func (p *Postgres) DeleteClient(telegramID int64) error {
	sqlRow := `
	DELETE FROM client
	WHERE  telegram_id = $1;
	`
	_, err := p.conn.Exec(p.ctx, sqlRow, telegramID)
	return err
}
