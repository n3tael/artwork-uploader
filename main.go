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

const (
	uploadServer string = "https://api.imgur.com/3/image"
)

type SuccessResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
}

func main() {
	var key = flag.String("key", "", "Imgur API key")
	var resizeImage = flag.Bool("resize", true, "Resize the image to 200px wide")
	flag.Parse()
	if *key == "" {
		flag.Usage()
		os.Exit(2)
	}
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "failed to get path from stdin: %v", scanner.Err())
		os.Exit(2)
	}
	imagePath := scanner.Text()
	
	image, err := processImage(imagePath, *resizeImage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to process image: %v", err)
		os.Exit(1)
	}

	link, err := upload(image, *key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to upload image: %v", err)
		os.Exit(1)
	}
	
	fmt.Println(link)
}

func processImage(imagePath string, resizeImage bool) (bytes.Buffer, error) {
	var image bytes.Buffer
	
	fileBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return bytes.Buffer{}, err
	}
	
	image.Write(fileBytes)

	if resizeImage == true {
		decodedImage, err := imaging.Decode(&image)
		if err != nil {
			return bytes.Buffer{}, err
		}
		
		resizedImage := imaging.Resize(decodedImage, 200, 0, imaging.Lanczos)
		
		if err := imaging.Encode(&image, resizedImage, 1); err != nil {
			return bytes.Buffer{}, err
		}
	}
	
	return image, nil
}

func upload(file bytes.Buffer, key string) (string, error) {
	client := &http.Client{}
	
	req, err := http.NewRequest("POST", uploadServer, &file)
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Authorization", "Client-ID " + key)
	
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
	
	successResponse := SuccessResponse{}
	if err := json.Unmarshal([]byte(content), &successResponse); err != nil {
		return "", err
	}
	
	return successResponse.Data.Link, nil
}