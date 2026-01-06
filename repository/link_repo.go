package repository

import "database/sql"

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(shortCode, originalURL string) error {
	query := `
		INSERT INTO dbo.Link (id, shortCode, originalUrl, clicks, createdAt)
		VALUES (NEWID(), @p1, @p2, 0, GETDATE())
	`
	_, err := r.db.Exec(query, shortCode, originalURL)
	return err
}

func (r *LinkRepository) FindByShortCode(shortCode string) (string, error) {
	var originalURL string
	query := `
		SELECT originalUrl
		FROM dbo.Link
		WHERE shortCode = @p1
	`
	err := r.db.QueryRow(query, shortCode).Scan(&originalURL)
	return originalURL, err
}

func (r *LinkRepository) IncreaseClick(shortCode string) {
	r.db.Exec(
		"UPDATE dbo.Link SET clicks = clicks + 1 WHERE shortCode = @p1",
		shortCode,
	)
}
