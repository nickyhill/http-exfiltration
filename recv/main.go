package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/netip"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Get preferred outbound ip of this machine
// Function is from stackoverflow at this link:
// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func createFileWrite(s string) {
	f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		log.Fatalln(err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Print the received data (you can process it as needed)
	fmt.Printf("Received POST request\n")
	createFileWrite(string(body) + "\n")
	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST request received"))
}

func processOutput() {
	f, err := os.ReadFile("output.txt")
	if err != nil {
		log.Fatalln(err)
	}

	decode_string := string(f)

	trimed_string := strings.Replace(decode_string, "\n", "", -1)
	ascii_string, err := hex.DecodeString(trimed_string)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("decoded.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	file.Write(ascii_string)

}

func init() {
	host_IP, _ := netip.AddrFromSlice(GetOutboundIP())
	fmt.Println("Server is active at:", host_IP)
}

func main() {

	// Register the postHandler function as the handler for "/post" route
	http.HandleFunc("/post", postHandler)

	// Create channel for interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Start the HTTP server on port 8080
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for an interrupts
	<-signalChan
	processOutput()
	fmt.Println("Received an interrupt, stopping server...")
}
