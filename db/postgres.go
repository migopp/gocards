package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/migopp/gocards/debug"
)

// This doesn't connect to DB provider, so have to
// fix at some point.
//
// For now, development can continue on a local SQLite DB,
// so not the end of the world. The actual deployment DB
// should be something like PostgreSQL though.

type PostgreDB struct {
	dsn string
	gdb *gorm.DB
}

func (p *PostgreDB) connect() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err == nil {
		p.gdb = db
	}
	return err
}

func (p *PostgreDB) migrate() error {
	debug.Assert(p.gdb != nil, "p.gdb == nil")
	return nil
	// return p.db.AutoMigrate(&User{}, &Deck{}, &Card{})
}

func (p *PostgreDB) db() *gorm.DB {
	debug.Assert(p.gdb != nil, "p.gdb == nil")
	return p.gdb
}
