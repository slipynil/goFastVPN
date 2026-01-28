package repository

// создает сущность клиента
func (p *Postgres) AddClient(username string, chatID int64) error {
	sqlRaw := `
	INSERT INTO client (username, chat_id, status)
	VALUES ($1, $2, false)
	`
	_, err := p.conn.Exec(p.ctx, sqlRaw, username, chatID)
	return err
}

func (p *Postgres) StatusTrue(chatID int64) error {
	sqlRaw := `
	UPDATE client
	SET status = true
	WHERE chat_id = $1;
	`
	_, err := p.conn.Exec(p.ctx, sqlRaw, chatID)
	return err
}

func (p *Postgres) StatusFalse(chatID int64) error {
	sqlRaw := `
	UPDATE client
	SET status = false
	WHERE chat_id = $1;
	`
	_, err := p.conn.Exec(p.ctx, sqlRaw, chatID)
	return err
}

func (p *Postgres) CheckStatus(chatID int64) bool {
	sqlRaw := `
	SELECT status
	FROM client
	WHERE chat_id = $1;
	`
	var status bool
	p.conn.QueryRow(p.ctx, sqlRaw, chatID).Scan(&status)
	return status
}

func (p *Postgres) IsTested(chatID int64) error {
	sqlRaw := `
	UPDATE client
	SET is_tested = true
	WHERE chat_id = $1;
	`
	_, err := p.conn.Exec(p.ctx, sqlRaw, chatID)
	return err
}
