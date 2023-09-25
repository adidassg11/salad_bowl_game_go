package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	fmt.Println("hello, establishing DB connection...")

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "salad_bowl_db",
	}

	// Get a db handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// log.Fatal(err)
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		// log.Fatal(pingErr)
		panic(pingErr.Error())
	}

	fmt.Println("Database Connected!")

	fmt.Println("Starting web server...")

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	router.StaticFile("/", "index.html")
	router.POST("/games", createGame)
	router.POST("/games/:id/words", createWord)

	router.Run("localhost:8080")
}

// Creates a new game in the database
func createGame(c *gin.Context) {
	res, err := db.Exec("INSERT INTO games (current_team, turn_end_time) VALUES (?, ?)", 1, time.Now().Add(1*time.Minute)) // TODO - make the time better
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	message := fmt.Sprintf("Game created with id {%s}", strconv.FormatInt(id, 10))
	c.JSON(200, gin.H{"message": message})
}

// Creates a new word
func createWord(c *gin.Context) {
	gameID := c.Param("id")

	// TODO - check here whether the timer has expired

	word := c.PostForm("word")
	team := c.PostForm("team")

	// TODO - need to validate that word and team are correct (ie. right now it's letting us add a blank word to DB despite the 'Allow Null' being FALSE)

	//gameIDInt, err := strconv.ParseInt(gameID, 10, 64)  // THIS is why we need an ORM * SMH *
	gameIDInt, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	teamInt, err := strconv.Atoi(team)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	// TODO - first convert game id
	res, err := db.Exec("INSERT INTO words (game_id, word, team) VALUES (?, ?, ?)", gameIDInt, word, teamInt)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	message := fmt.Sprintf("Word added with id {%s}", strconv.FormatInt(id, 10))
	c.JSON(200, gin.H{"message": message})
}