package repository

import (
	"context"
	"fmt"
	"telegram-service/internal/dto"
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

func (p *Postgres) Close() error {
	return p.conn.Close(p.ctx)
}

// ping postgres server
func (p *Postgres) Ping() error {
	return p.conn.Ping(p.ctx)
}

// add user`s identity in client table
func (p *Postgres) AddUser(username string, telegramID int64, expiresAt time.Time) error {
	// Начинаем транзакцию для атомарности операции
	tx, err := p.conn.Begin(p.ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(p.ctx) // Откатываем, если не будет Commit

	// Проверяем, есть ли уже клиент с таким telegram_id (с блокировкой строки)
	var exists bool
	err = tx.QueryRow(p.ctx, `
		SELECT EXISTS (SELECT 1 FROM client WHERE telegram_id = $1 FOR UPDATE)
	`, telegramID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if client exists: %w", err)
	}

	if exists {
		// Клиент уже есть — возвращаем свою ошибку
		return dto.ErrUserExist
	}

	// Только если нет — вставляем нового
	// Находим минимальный свободный host_id (заполняем "дырки" от удаленных записей)
	_, err = tx.Exec(p.ctx, `
		INSERT INTO client (host_id, telegram_username, telegram_id, status, expires_at)
		VALUES (
			COALESCE(
				(SELECT t1.host_id + 1
				 FROM client t1
				 LEFT JOIN client t2 ON t1.host_id + 1 = t2.host_id
				 WHERE t2.host_id IS NULL AND t1.host_id + 1 > 1
				 ORDER BY t1.host_id
				 LIMIT 1),
				COALESCE((SELECT MAX(host_id) + 1 FROM client), 2)
			),
			$1, $2, false, $3
		)
	`, username, telegramID, expiresAt)
	if err != nil {
		return fmt.Errorf("failed to insert client: %w", err)
	}

	// Коммитим транзакцию
	if err = tx.Commit(p.ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
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
