package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Album struct {
	id     int
	title  string
	artist string
	price  float64
}

func main() {
	connStr := "user=gopsql dbname=gopsql host=localhost password=gopsql sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	albumsByColtrane, err := AlbumsByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	for _, album := range albumsByColtrane {
		fmt.Println(album.id, album.title, album.price)
	}
}

func AlbumsByArtist(db *sql.DB, name string) ([]Album, error) {
	rows, err := db.Query("SELECT id, title, artist, price FROM album WHERE artist = $1x", name)
	if err != nil {
		log.Fatal(err)
	}
	albums := make([]Album, 0)
	for rows.Next() {
		var id int
		var title string
		var artist string
		var price float64
		err := rows.Scan(&id, &title, &artist, &price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, Album{id, title, artist, price})
	}
	return albums, nil
}
