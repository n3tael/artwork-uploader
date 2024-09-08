package main

import (
	"os"
	"io"
	"fmt"
	"flag"
	"bytes"
	"bufio"
	"net/http"
	"encoding/json"
	
	"github.com/disintegration/imaging"
)

var upload_server string = "https://api.imgur.com/3/image"
type SuccessResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
}

func main() {
	var apikey = flag.String("key", "", "Imgur API key")
	flag.Parse()
	if *apikey == "" {
		flag.Usage()
		os.Exit(2)
	}
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "failed to get path from stdin: %v", scanner.Err())
		os.Exit(2)
	}
	image_path := scanner.Text()
	
	image, err := process_image(image_path)
	if err != nil {
	    fmt.Fprintf(os.Stderr, "failed to process image: %v", err)
		os.Exit(1)
	}
	
	link, err := upload(image, *apikey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to upload image: %v", err)
		os.Exit(1)
	}
	
	fmt.Println(link)
}

func process_image(image_path string) (bytes.Buffer, error) {
	var encoded_image bytes.Buffer
	
	image, err := imaging.Open(image_path)
	if err != nil {
		return bytes.Buffer{}, err
	}
	
	resized_image := imaging.Resize(image, 200, 0, imaging.Lanczos)
	
	if err := imaging.Encode(&encoded_image, resized_image, 1); err != nil {
		return bytes.Buffer{}, err
	}
	
	return encoded_image, nil
}

func upload(file bytes.Buffer, key string) (string, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest("POST", upload_server, &file)
	if err != nil {
		return "", err
	}
	
	req.SetBasicAuth("Client-ID", key)
	
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
    contentBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
    content := string(contentBytes)
	
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("server responded not OK: %v", content)
	}
	
	success_response := SuccessResponse{}
	if err := json.Unmarshal([]byte(content), &success_response); err != nil {
		return "", err
	}
	
	return success_response.Data.Link, nil
}