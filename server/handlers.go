package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/migopp/gocards/db"
)

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{})
}

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl.html", gin.H{})
}

func postSignup(c *gin.Context) {
	// Get signup details
	email := c.PostForm("email")
	pass := c.PostForm("password")

	// Encrypt password
	epass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		log.Printf("Failed to encrypt password: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to encrypt password",
		})
		return
	}

	// Create a user with credentials
	u := db.User{
		Email:    email,
		Password: string(epass),
	}
	if err := db.GCDB.CreateUser(&u); err != nil {
		log.Printf("Failed to create user: %s\n", err)
		errs := fmt.Sprintf("Failed to create user: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": errs,
		})
		return
	}

	// Issue JWT
	if err := issueJWT(u, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Auth generation failed",
		})
		return
	}

	// Redirect to index
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
	return

}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", gin.H{})
}

func postLogin(c *gin.Context) {
	// Get login details
	email := c.PostForm("email")
	pass := c.PostForm("password")

	// Fetch details from DB and verify credentials
	u, err := db.GCDB.FetchUserWithEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Unknown user",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad password",
		})
		return
	}

	// Issue JWT
	if err := issueJWT(u, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Auth generation failed",
		})
		return
	}

	// Redirect to index
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
	return
}

func getCards(c *gin.Context) {
	c.HTML(http.StatusOK, "cards.tmpl.html", gin.H{})
}

func getDecks(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{})
}

func postDecks(c *gin.Context) {
	// Parse the uploaded file
	const fs = 10 << 20 // ~10mb
	if err := c.Request.ParseMultipartForm(fs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to parse deck upload form",
		})
		return
	}

	// Open the file as a multipart
	file, header, err := c.Request.FormFile("deck-file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to open deck from form data",
		})
		return
	}
	defer file.Close()

	// Parse file into `LDeck`
	ld, err := db.YMLToDeck(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to parse deck from `.yml`",
		})
		return
	}

	// Upload the deck to the DB
	//
	// This requires us to know which user we are working
	// on behalf of, so we fetch it from the JWT first
	u, err := getSessionUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to fetch session user",
		})
		return
	}
	if err := db.GCDB.LoadDeck(&ld, u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed load deck into DB",
		})
		return
	}

	// Serve updates
	//
	// First, we will need to fetch all the available decks
	// for the session user
	decks, err := db.GCDB.FetchDecksForUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed fetch decks for session user",
		})
		return
	}
	fmt.Printf("Decks = %v\n", decks)
	c.HTML(http.StatusOK, "decks.comp.html", gin.H{
		"Decks": decks,
	})
}
