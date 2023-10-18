package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var recpt = flag.String("h", "", "host server IP address")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	url := "http://" + *recpt + ":8080/post" // Replace with your server's URL

	// Create a client and send the request
	client := &http.Client{}
	// The data you want to send in the POST request

	// Open the local file
	f, err := os.Open("textfile.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Read file in 60 byte chunks
	dataBytes := make([]byte, 60)

	for {
		dataBytes = dataBytes[:cap(dataBytes)]
		bytesRead, err := f.Read(dataBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		dataBytes = dataBytes[:bytesRead]
		hexString := hex.EncodeToString(dataBytes)

		postData := []byte(hexString)

		// Create a request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// Check the response
		fmt.Println("Response Status:", resp.Status)
	}
}
