package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var SECRET string = os.Getenv("SECRET")

func hasValidSignature(body []byte, headerSignature string) bool {
	fmt.Printf("received signature: " + headerSignature + "\n")

	signature, err := hex.DecodeString(headerSignature)
	if err != nil {
		return false
	}

	var secret = []byte(SECRET)

	hasher := hmac.New(sha256.New, secret)
	hasher.Write(body)
	expected := hasher.Sum(nil)
	expectedHex := hex.EncodeToString(expected)
	fmt.Printf("expected signature: " + expectedHex)

	if !hmac.Equal(expected, signature) {
		fmt.Printf("invalid signature")
		return false
	}

	return true
}

func main() {
	if SECRET == "" {
		fmt.Println("Secret variable not set or empty.")
		panic("Secret variable not set or empty.")
	}

	eventMessages := map[string]map[string]interface{}{}

	// Dummy webhook message handler
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("could not read body")
			w.WriteHeader(400)
			return
		}

		println("received body:")
		fmt.Println(body)

		signature := r.Header.Get("X-IFood-Signature")

		if !hasValidSignature(body, signature) {
			fmt.Printf("invalid signature")
			// fmt.Printf("got: " + signature + "\n")
			// fmt.Printf("expected: " + expected + "\n")
			w.WriteHeader(400)
			return
		}

		var genericMsg map[string]interface{}
		err = json.Unmarshal(body, &genericMsg)
		if err != nil {
			fmt.Printf("failed to parse body:\n%s", body)
			w.WriteHeader(400)
			return
		}

		fmt.Printf("received message: %v\n", genericMsg)

		eventMessages["eventId"] = genericMsg

		w.Header().Add("Content-Type", "application/json")
		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
