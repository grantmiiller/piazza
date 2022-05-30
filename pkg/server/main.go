package server

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// Verbose logging
func vLog(s string, v bool) {
	if v {
		fmt.Println(s)
	}
}

func getListener(q string, v bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		rawValue := query.Get(q)
		// Query removes the "+" from the string, making it impossible to base64 decode
		rawValue = strings.Replace(rawValue, " ", "+", -1)
		fmt.Printf("Connection from: %s\n\n", r.RemoteAddr)
		if rawValue != "" {
			decodedBytes, err := base64.StdEncoding.DecodeString(rawValue)
			if err != nil {
				vLog("====Printing Raw String====", v)
				fmt.Println(rawValue)
				vLog("==== END Raw String====", v)
			} else {
				vLog("====Printing Decoded String====", v)
				fmt.Println(string(decodedBytes))
				vLog("==== END Decoded String====", v)
			}
		} else {
			fmt.Println("No value received")
		}
		w.WriteHeader(201)
	}
}

// Pass configuration file instead of individual values
func StartServer(port string, query string, verbose bool) {
	http.HandleFunc("/", getListener(query, verbose))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
