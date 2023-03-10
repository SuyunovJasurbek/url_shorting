package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/config"
	"github.com/SuyunovJasurbek/url_shorting/src/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cast"
)

type Storage struct {
	db   *sqlx.DB
	user storage.UserI
	url  storage.UrlI
}

func NewPostgres(cfg *config.Config) (storage.StorageI, error) {
	psqlConnString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
	)
	fmt.Println(psqlConnString)
	db, err := sqlx.Open("postgres", psqlConnString)
	if err != nil {
		log.Fatalf("cannot connect to postgresql db: %s", err.Error())
	}

	// setting some settings for db
	db.SetConnMaxLifetime(time.Duration(time.Duration(cast.ToInt(cfg.PostgresConnMaxIdleTime)).Minutes()))
	db.SetMaxOpenConns(cast.ToInt(cfg.PostgresMaxConnections))

	// checking for Ping&Pong
	if err := db.Ping(); err != nil {
		log.Fatalf("ping error %s", err.Error())
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) User() storage.UserI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}

func (s *Storage) Url() storage.UrlI {
	if s.url == nil {
		s.url = NewUrlRepo(s.db)
	}
	return s.url
}
