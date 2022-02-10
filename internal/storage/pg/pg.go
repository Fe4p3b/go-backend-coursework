package pg

import (
	"context"
	"database/sql"
	"time"

	"github.com/Fe4p3b/go-backend-coursework/internal/repositories"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type pg struct {
	db *sql.DB
}

var _ repositories.ShortenerRepository = &pg{}

func NewConnection(dsn string) (*pg, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &pg{db: conn}, nil
}

func (p *pg) Find(shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sql := `SELECT original_url FROM shortener.shortener WHERE short_url=$1`

	var URL string

	row := p.db.QueryRowContext(ctx, sql, shortURL)

	if err := row.Scan(&URL); err != nil {
		return "", err
	}

	return URL, nil
}

func (p *pg) Save(shortURL string, URL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sql := `INSERT INTO shortener.shortener(short_url, original_url) VALUES($1, $2)`

	_, err := p.db.ExecContext(ctx, sql, shortURL, URL)
	if err != nil {
		return err
	}

	return err
}

func (p *pg) AddCount(shortURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "UPDATE shortener.shortener SET visitors_counter = visitors_counter + 1 WHERE short_url=$1")
	if err != nil {
		return err
	}

	if _, err = stmt.ExecContext(ctx, shortURL); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *pg) GetVisitorCounter(shortURL string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sql := `SELECT visitors_counter FROM shortener.shortener WHERE short_url=$1`

	var counter int

	row := p.db.QueryRowContext(ctx, sql, shortURL)

	if err := row.Scan(&counter); err != nil {
		return 0, err
	}

	return counter, nil
}
