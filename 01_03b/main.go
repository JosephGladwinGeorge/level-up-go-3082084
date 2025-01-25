package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	// TODO: Fill in definition
	Id  string  `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var entries []raffleEntry
	f,err:=os.ReadFile(path)
		if err!=nil{
			log.Fatal(err)
		}
	//var file []byte
	//f.Read(file)
	err=json.Unmarshal(f,&entries)
	if err!=nil{
		log.Fatal(err)
	}

	println(entries)

	return entries
		
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}