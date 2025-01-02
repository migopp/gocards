package db_test

import (
	"log"
	"os"
	"testing"

	"github.com/migopp/gocards/db"
	"github.com/migopp/gocards/test"
)

var TDB *db.DB

func TestMain(m *testing.M) {
	// Open DB
	TDB = db.New(db.SQLite, "test.db")
	if err := TDB.Connect(); err != nil {
		log.Fatalf("Failed to connect to test DB: %v", err)
	}
	if err := TDB.Migrate(); err != nil {
		log.Fatalf("Failed to migrate test DB: %v", err)
	}

	// Run tests
	code := m.Run()

	// Clean up, clean up, everybody everywhere.
	// Clean up, clean up, everybody do your share.
	//
	// We may actually want to _not_ clean this up
	// for manual DB inspection, but if that's the case,
	// I'll remove this code manually.
	_ = os.Remove("test.db")

	// Leave
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	// Create a sample user
	su := db.User{
		Email:    "testA@test.mailprov.com",
		Password: "someencryptedpassword",
	}
	if err := TDB.CreateUser(&su); err != nil {
		t.Errorf("Failed to create user `su`: %v", err)
	}
}

func TestCreateAndIDFetchUser(t *testing.T) {
	// Create a sample user
	su := db.User{
		Email:    "testB@test.mailprov.com",
		Password: "someencryptedpassword",
	}
	if err := TDB.CreateUser(&su); err != nil {
		t.Errorf("Failed to create user `su`: %v", err)
	}

	// Fetch based on email and ensure fields match
	fu, err := TDB.FetchUserWithID(su.ID)
	test.AssertEq(err, nil, t)
	test.AssertEq(fu.ID, su.ID, t)
	test.AssertEq(fu.Email, su.Email, t)
	test.AssertEq(fu.Password, su.Password, t)
}

func TestCreateAndEmailFetchUser(t *testing.T) {
	// Create a sample user
	su := db.User{
		Email:    "testC@test.mailprov.com",
		Password: "someencryptedpassword",
	}
	if err := TDB.CreateUser(&su); err != nil {
		t.Errorf("Failed to create user `su`: %v", err)
	}

	// Fetch based on email and ensure fields match
	fu, err := TDB.FetchUserWithEmail("testC@test.mailprov.com")
	test.AssertEq(err, nil, t)
	test.AssertEq(fu.ID, su.ID, t)
	test.AssertEq(fu.Email, su.Email, t)
	test.AssertEq(fu.Password, su.Password, t)
}
