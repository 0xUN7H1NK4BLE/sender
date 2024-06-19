package main

import (
    "bytes"
    "flag"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

func uploadFile(webhookURL, filePath, message string) error {
    var requestBody bytes.Buffer
    writer := multipart.NewWriter(&requestBody)

    // Add the file
    if filePath != "" {
        file, err := os.Open(filePath)
        if err != nil {
            return fmt.Errorf("could not open file: %v", err)
        }
        defer file.Close()

        part, err := writer.CreateFormFile("file", filepath.Base(filePath))
        if err != nil {
            return fmt.Errorf("could not create form file: %v", err)
        }

        _, err = io.Copy(part, file)
        if err != nil {
            return fmt.Errorf("could not copy file content: %v", err)
        }
    }

    // Add the message
    if message != "" {
        _ = writer.WriteField("content", message)
    }

    writer.Close()

    req, err := http.NewRequest("POST", webhookURL, &requestBody)
    if err != nil {
        return fmt.Errorf("could not create request: %v", err)
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("could not send request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    return nil
}

func main() {
    webhookURL := flag.String("w", "", "Discord webhook URL")
    filePath := flag.String("f", "", "Path to the file to upload")
    message := flag.String("m", "", "Message to send with the file")
    flag.Parse()

    if *webhookURL == "" {
        fmt.Println("\033[31mError: Discord webhook URL must be provided\033[0m") // Red color for error
        return
    }

    if *filePath == "" && *message == "" {
        fmt.Println("\033[31mError: Either file path or message must be provided\033[0m") // Red color for error
        return
    }

    err := uploadFile(*webhookURL, *filePath, *message)
    if err != nil {
        fmt.Printf("\033[31mError uploading file: %v\033[0m\n", err) // Red color for error
        return
    }

    fmt.Println("\033[32mFile uploaded successfully!\033[0m") // Green color for success
}
