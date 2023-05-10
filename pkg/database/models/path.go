package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sqlite "github.com/reagancn/filewatch/pkg/database"
)

/* Insert path to database */
func (p *Path) Insert(db *sql.DB) (string, error) {

	// Get added timestamp
	addedOn := time.Now().Unix()
	strAddedOn := fmt.Sprint(addedOn)
	p.AddedOn = strAddedOn
	p.LastModified = strAddedOn

	// Get file changes to watch
	// TODO: Make this configurable
	p.Events = []Event{{Type: "CREATE"}, {Type: "MODIFY"}, {Type: "DELETE"}}

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS paths (id INTEGER PRIMARY KEY, path TEXT, last_modified INTEGER, added_on INTEGER);")
	if err != nil {
		log.Println("Error creating files table")
		log.Fatal(err)
	}

	result, err := sqlite.Insert(db, "paths", []string{"path", "last_modified", "added_on"}, []string{p.Path, p.LastModified, p.AddedOn})

	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Path added to watchlist : ", id, " - ", p.Path)

	return fmt.Sprint(id), err
}

/* Get all paths from database */
func (p *Path) Get(db *sql.DB) ([]Path, error) {

	rows, err := sqlite.Get(db, "SELECT * FROM paths;")

	if err != nil {
		log.Println("Error getting paths")
		log.Fatal(err)
	}

	paths := []Path{}

	for rows.Next() {
		var id int
		var path string
		var last_modified int64
		var added_on int64
		err = rows.Scan(&id, &path, &last_modified, &added_on)
		if err != nil {
			log.Println("Error scanning rows")
			log.Fatal(err)
		}
		log.Printf("ID: %d, Path: %s, Last Modified: %d, Added On: %d\n", id, path, last_modified, added_on)
		paths = append(paths, Path{ID: id, Path: path, LastModified: fmt.Sprint(last_modified), AddedOn: fmt.Sprint(added_on)})
	}

	return paths, err
}
