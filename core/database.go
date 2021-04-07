package core

import (
	"fmt"
	"github.com/pressly/goose"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const (
	MigrationFolder = "core/db/migrations/"
)

// Base model which als models that should be persisted extend
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Base repository which all repositories extend
type BaseRepository struct {
}

// Returns the singleton database handle
func (receiver BaseRepository) Db() *gorm.DB {
	return databaseHandle
}

// Finds or panics
func (receiver BaseRepository) FindFirstOrPanic(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	result := receiver.Db().First(dest, conds...)
	if err := result.Error; err != nil {
		panic(err)
	}

	return result
}

// Singleton - only accessible through the repository
var databaseHandle *gorm.DB

// Package init
func init() {
	_db, err := gorm.Open(sqlite.Open(DatabaseFile), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", DatabaseFile))
	}

	databaseHandle = _db
}

// Performs migrations
func Migrate() {
	db, _ := goose.OpenDBWithDriver(DatabaseDriver, DatabaseFile)
	goose.SetVerbose(true)
	if err := goose.Up(db, MigrationFolder); err != nil {
		panic(fmt.Sprintf("Migration failed: %s", err))
	}
}
