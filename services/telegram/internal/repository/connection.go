package repository

import "time"

func (p *Postgres) NewConnection(chatID int64, expires_at time.Time) error {
	sqlRaw := `
	INSERT INTO peer (chat_id, expires_at)
	VALUES ($1, $2);
	`
	_, err := p.conn.Exec(p.ctx, sqlRaw, chatID, expires_at)
	return err
}

func (p *Postgres) GetHostID(chatID int64) (int, error) {
	sqlRaw := `
	SELECT host_id FROM peer
	WHERE chat_id = $1;
	`
	var hostID int
	err := p.conn.QueryRow(p.ctx, sqlRaw, chatID).Scan(&hostID)
	return hostID, err
}
