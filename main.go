package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

func main() {
	// For InmemoryTest
	// store := NewInMemoryPlayerStore()
	// server := NewPlayerServer(store)

	// log.Fatal(http.ListenAndServe(":5000", server))

	// For FileSystemTest
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		fmt.Printf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
