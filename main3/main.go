package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := &http.Client{Timeout: time.Second * 20}
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/stream", strings.NewReader(""))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(resp.Body)
	defer resp.Body.Close()

	for {
		rawLine, err := reader.ReadBytes('\n')
		if errors.Is(err, io.EOF) {
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bytes.TrimRight(rawLine, "\n")))
	}
}