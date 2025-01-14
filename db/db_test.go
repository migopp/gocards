package db_test

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
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

func simulateMultipartFile(fp string) (multipart.File, *multipart.FileHeader, error) {
	// Open the local file
	of, err := os.Open(fp)
	if err != nil {
		return nil, nil, err
	}
	defer of.Close()

	// Create a buffer to write the multipart file
	//
	// Multipart files may not actually be backed on
	// the file system, and may just live in memory
	// -- just as we are going to force this one to live
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	ff, err := writer.CreateFormFile("fakempf", fp)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(ff, of)
	if err != nil {
		return nil, nil, err
	}
	writer.Close()

	// Create a reader and simulate parsing
	reader := multipart.NewReader(&buffer, writer.Boundary())
	const fs = 10 << 20
	form, err := reader.ReadForm(fs)
	if err != nil {
		return nil, nil, err
	}
	mpfh := form.File["fakempf"][0]
	mpf, err := mpfh.Open()
	if err != nil {
		return nil, nil, err
	}

	return mpf, mpfh, nil
}

/*
// TODO: Fix portability of this test
func TestTenYMLToLDeck(t *testing.T) {
	// Open a sample deck as a `multipart.File` : "../test/decks/ten.yml"
	file, header, err := simulateMultipartFile("../test/decks/ten.yml")
	if err != nil {
		t.Errorf("Failed to open `../test/decks/ten.yml` as multipart file: %v", err)
	}
	defer file.Close()

	// Convert to `LDeck`
	ld, err := db.YMLToLDeck(file, header)
	if err != nil {
		t.Errorf("Failed to convert `../test/decks/ten.yml` to `LDeck`: %v", err)
	}
	// Assert things about the deck.
	// In particular, check (the other fields should be loaded later):
	//	- Deck: Name
	//  - Card: Front, Back
	test.AssertEq(ld.DBDeck.Name, "some_deck", t)
	test.AssertEq(ld.DBCards[0].Front, "一", t)
	test.AssertEq(ld.DBCards[0].Back, "いち", t)

	// The owner of this deck will be a new user
	su := db.User{
		Email:    "TestTenYMLToLDeck@test.mailprov.com",
		Password: "someencryptedpassword",
	}
	if err := TDB.CreateUser(&su); err != nil {
		t.Errorf("Failed to create user `su`: %v", err)
	}

	// Load into DB
	if err = TDB.LoadDeck(&ld, su); err != nil {
		t.Errorf("Failed to load `LDeck` into DB: %v", err)
	}

	// Assert things about the metadata of the deck
	//
	// These are the things that we deferred filling earlier
	test.AssertEq(ld.DBDeck.UserID, su.ID, t)
	test.AssertEq(ld.DBDeck.User.ID, su.ID, t)
	test.AssertEq(ld.DBCards[0].DeckID, ld.DBDeck.ID, t)
	test.AssertEq(ld.DBCards[0].Deck, ld.DBDeck, t)

	// Fetch the deck we just loaded in
	decks, err := TDB.FetchDecksForUser(su)
	if err != nil {
		t.Errorf("Failed to fetch decks for user: %v", su)
	}
	test.AssertEq(decks[0].ID, ld.DBDeck.ID, t)
}
*/
