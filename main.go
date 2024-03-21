package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thisisthemurph/gob/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	hostURL := os.Getenv("HOST_URL")
	token := os.Getenv("TOKEN")

	clusterAPI := api.NewClusterAPI(hostURL, token)

	clusters, err := clusterAPI.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, cluster := range clusters {
		fmt.Println(cluster.Name)
		c, err := clusterAPI.Get(cluster.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.String())
	}
}
