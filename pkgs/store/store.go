package store

import (
	"camera-server-backend/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Store ...
type Store struct {
	db *gorm.DB
}

// NewStore ...
func NewStore() (*Store, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1234 dbname=camera host=localhost port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	s := Store{db: db}
	s.db.AutoMigrate(&model.Event{})
	return &s, nil
}

// InsertEvents ...
func (s *Store) InsertEvents(events []model.Event) error {
	s.db.Create(events)
	return nil
}

// GetAllEvents ...
func (s *Store) GetAllEvents() (events []model.Event, err error) {
	s.db.Find(&events)
	return events, nil
}
