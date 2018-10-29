package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/yanmhlv/test_assignment/app"
)

var (
	flagURL      = flag.String("url", "http://localhost:3000/game", "")
	flagClientID = flag.Int("client-id", 0, "")
	flagFee      = flag.Float64("fee", 10, "")
)

func main() {
	flag.Parse()

	randomPair := make([]byte, 2)
	n, err := rand.Read(randomPair)
	if err != nil || n != 2 {
		panic(err)
	}

	req := app.GameRequest{
		ClientID: *flagClientID,
		Fee:      *flagFee,
		Pair:     app.Pair{randomPair[0], randomPair[1]},
	}

	fmt.Printf("send request %+v to %s\n", req, *flagURL)

	data, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(*flagURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result := app.GameResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
