package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Generator 
	rand.Seed(time.Now().UnixNano())

	// Post ke url
	url := "https://jsonplaceholder.typicode.com/posts"

	
	genRandVal := func() int {
		return rand.Intn(100) + 1
	}

	// Perulangan
	getStatus := func(val int) string {
		switch {
		case val < 5:
			return "aman"
		case val < 9:
			return "siaga"
		default:
			return "bahaya"
		}
	}

	
	sendPOST := func(water, wind int) {
	
		data := map[string]int{"water": water, "wind": wind}

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))//Kirim Request
		if err != nil {
			panic(err)
		}

	
		fmt.Println("Response:", resp.Status)
	}

	
	sendPOST(genRandVal(), genRandVal())

	
	for range time.Tick(15 * time.Second) {
		water := genRandVal()
		wind := genRandVal()

		fmt.Printf("water = %d\n ", water)
		fmt.Printf("wind = %d\n", wind)
		fmt.Printf("status water = %s\n", getStatus(water))
		fmt.Printf("status wind = %s\n",getStatus(wind) )

		sendPOST(water, wind)
	}
}