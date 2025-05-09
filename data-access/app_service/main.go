package main

//reference: https://go.dev/doc/tutorial/database-access
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {

	// Get a databse handle
	//fmt.Println(cfg.FormatDSN())

	cfgStr, cfgError := getConnectionProperties("properties.env")

	if cfgError != nil {
		log.Fatal(cfgError)
	}

	var err error
	db, err = sql.Open("mysql", cfgStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connected!")

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumById(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}

		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func getConnectionProperties(str string) (string, error) {
	envErr := godotenv.Load(str)

	if envErr == nil {
		// Capture connection properties.
		cfg := mysql.Config{
			User:   os.Getenv("MYSQL_USER"),
			Passwd: os.Getenv("MYSQL_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("MYSQL_ADDRESS"),
			DBName: os.Getenv("MYSQL_DATABASE"),
		}

		return cfg.FormatDSN(), nil
	} else {
		return "", envErr
	}
}
