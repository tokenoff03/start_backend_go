package repositories

import (
	"fmt"
	"gin/internal/config"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	urlPath := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password)
	db, err := sqlx.Open("postgres", urlPath)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
