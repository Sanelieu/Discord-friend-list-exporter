package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("Enter Discord Token")
	var token string
	fmt.Scanln(&token)

	relationships, err := getRelationships(token)
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()
	fileName := currentTime.Format("2006-01-02_15-04-05.txt")

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	log.Println("File", fileName, "created successfully.")

	for i, relationship := range relationships {
		rLog := fmt.Sprintf("Tag: %s:%s | ID: %s", relationship.User.Username, relationship.User.Discriminator, relationship.ID)
		log.Printf("[%d] Exporting relationship | %s ", i, rLog)

		// Write rLog to the file
		_, err := file.WriteString(rLog + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
}
