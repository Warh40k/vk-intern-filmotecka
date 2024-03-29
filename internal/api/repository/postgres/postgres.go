package postgres

import (
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	actorsTable      = "actors"
	filmsTable       = "films"
	filmsActorsTable = "films_actors"
)

var (
	ErrUnique   = errors.New("unique costraint violation")
	ErrNoRows   = errors.New("no rows in relation")
	ErrInternal = errors.New("internal error")
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	constring := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("pgx", constring)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
