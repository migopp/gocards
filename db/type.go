package db

type DBType int

const (
	SQLite DBType = iota
	PostgreSQL
)
