package db

import (
	"gorm.io/gorm"

	"github.com/migopp/gocards/debug"
)

type dbi interface {
	connect() error
	migrate() error
	db() *gorm.DB
}

type DB struct {
	Type DBType
	dbi  dbi
}

func New(t DBType, s any) *DB {
	switch t {
	case SQLite:
		f, ok := s.(string)
		debug.Assert(ok, "SQLite DB name is non-string")
		return &DB{
			Type: SQLite,
			dbi: &SQLiteDB{
				name: f,
			},
		}
	case PostgreSQL:
		conn, ok := s.(string)
		debug.Assert(ok, "Postgre DB DSN is non-string")
		return &DB{
			Type: PostgreSQL,
			dbi: &PostgreDB{
				dsn: conn,
			},
		}
	default:
		panic("Attempt to create DB with undefined type")
	}
}

func (db *DB) Connect() error {
	return db.dbi.connect()
}

func (db *DB) Migrate() error {
	return db.dbi.migrate()
}

func (db *DB) CreateUser(u *User) error {
	r := db.dbi.db().Create(u)
	return r.Error
}

func (db *DB) CreateDeck(d *Deck) error {
	r := db.dbi.db().Create(d)
	return r.Error
}

func (db *DB) CreateCard(c *Card) error {
	r := db.dbi.db().Create(c)
	return r.Error
}

func (db *DB) FetchUserWithID(id uint) (User, error) {
	var u User
	r := db.dbi.db().Where("id = ?", id).First(&u)
	return u, r.Error
}

func (db *DB) FetchUserWithEmail(e string) (User, error) {
	var u User
	r := db.dbi.db().Where("email = ?", e).First(&u)
	return u, r.Error
}

func (db *DB) FetchDecksForUser(u User) ([]Deck, error) {
	var decks []Deck
	r := db.dbi.db().Where("user_id = ?", u.ID).Find(&decks)
	return decks, r.Error
}
