package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Session holds a GORM session
type Session struct {
	db     *gorm.DB
	config *Config
}

// Database holds the dbname and credentials
type Database struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Config represents a json configuration file used for creating session against PostgreSQL
type Config struct {
	Environment          string   `json:"environment"`
	Database             Database `json:"database"`
	AccessControlEnabled bool     `json:"accessControlEnabled"`
	Host                 string   `json:"host"`
	Port                 string   `json:"port"`
	SSLMode              string   `json:"sslmode"`
}

// NewSession creates a new session with PostgreSQL using GORM
func NewSession(config *Config) (*Session, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.Database.Username, config.Database.Password,
		config.Database.Name, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Session{
		db:     db,
		config: config,
	}, nil
}

// Copy creates a new connection session
func (s *Session) Copy() (*Session, error) {
	return &Session{db: s.db.Session(&gorm.Session{NewDB: true}), config: s.config}, nil
}

// GetDB returns the GORM database instance
func (s *Session) GetDB() *gorm.DB {
	return s.db
}

// GetConfig returns a pointer to the session configuration
func (s *Session) GetConfig() *Config {
	return s.config
}

// Close terminates the session
func (s *Session) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// DropDatabase deletes the database from PostgreSQL
func (s *Session) DropDatabase() error {
	err := s.db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", s.config.Database.Name)).Error
	return err
}

// SetSyncTime sets the statement timeout for queries
func (s *Session) SetSyncTime(t time.Duration) error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(t)
	return nil
}

// GetTableData retrieves data from a given table
func (s *Session) GetTableData(table string, out interface{}) error {
	return s.db.Table(table).Find(out).Error
}
