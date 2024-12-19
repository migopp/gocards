package db

import "database/sql"

// I'm aware that this is somewhat scuffed, but I feel
// like it is more scuffed to have the db directly associated
// with the server
//
// They are technically separate entities...
//
// For now, let's just make the db some global entity
var db *sql.DB
