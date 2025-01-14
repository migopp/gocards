package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/migopp/gocards/debug"
)

type SQLiteDB struct {
	name string
	gdb  *gorm.DB
}

func (s *SQLiteDB) connect() error {
	db, err := gorm.Open(sqlite.Open(s.name), &gorm.Config{})
	if err == nil {
		s.gdb = db
	}
	return err
}

func (s *SQLiteDB) migrate() error {
	debug.Assert(s.gdb != nil, "s.gdb == nil")
	return s.gdb.AutoMigrate(&User{}, &Deck{}, &Card{})
}

func (s *SQLiteDB) db() *gorm.DB {
	debug.Assert(s.gdb != nil, "s.gdb == nil")
	return s.gdb
}
