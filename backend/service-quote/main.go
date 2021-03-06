package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jsonFile, err := os.Open("quotes.json")
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	quotes := []quote{}
	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] host=%s", r.Host)

		ret := []quote{}
		for i := 0; i < 6; i++ {
			ret = append(ret, quotes[rand.Intn(len(quotes))])
		}
		json.NewEncoder(w).Encode(ret)
	})

	log.Println("[INFO] Server is running on port:" + os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}
