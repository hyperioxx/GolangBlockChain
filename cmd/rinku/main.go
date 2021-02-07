package main

import (
	"time"
	"log"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"Rinku/webserver"
	"Rinku/blockchain"
)


func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := blockchain.Block{0, t.String(), "", "", ""}
		spew.Dump(genesisBlock)
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
	}()


	log.Fatal(webserver.Run())

}