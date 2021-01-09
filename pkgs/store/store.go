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

const dnsStore = "host=postgres port=5432 user=dbo password=dbo dbname=camera sslmode=disable"

// NewStore ...
func NewStore() (*Store, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dnsStore}), &gorm.Config{})

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
	s.db.Order("id desc").Find(&events)
	return events, nil
}
