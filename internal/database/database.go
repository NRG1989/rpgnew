package database

import (
	"context"
	"go-aut-registration-user-otp/internal/config"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Client struct {
	*sqlx.DB
	logger     logrus.FieldLogger
	schemaName string
}

func NewClient(cfg *config.Config, logger logrus.FieldLogger) (*Client, error) {
	db, err := sqlx.Open("pgx", cfg.DB.URL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	return &Client{
		db,
		logger,
		cfg.DB.SchemaName,
	}, nil
}

const statusCheckQuery = `SELECT true`

func (db *Client) StatusCheck(ctx context.Context) error {
	var tmp bool

	return db.QueryRowContext(ctx, statusCheckQuery).Scan(&tmp)
}
