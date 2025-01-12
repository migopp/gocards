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
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
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
	}

	// Redirect to index
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
	return

}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
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
	}

	// Redirect to index
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
	return
}

func getCards(c *gin.Context) {
	c.HTML(http.StatusOK, "cards.html", gin.H{})
}

func postCards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
