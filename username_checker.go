package main

import (
    "bufio"
    "bytes"
    "flag"
    "fmt"
    "net/http"
    "os"
)

// checkUsername sends a POST request with the given username as JSON payload
// to the target URL and prints whether the username exists based on the response.
func checkUsername(username, url string) {
    // Create the JSON payload. You may need to adapt this to match the expected API format.
    payload := []byte(fmt.Sprintf(`{"username": "%s"}`, username))

    // POST request to the target URL
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err != nil {
        fmt.Printf("Error checking username '%s': %v\n", username, err)
        return
    }
    defer resp.Body.Close()

    // Interpret the HTTP response status
    if resp.StatusCode == http.StatusOK {
        fmt.Printf("Username '%s' exists!\n", username)
    } else {
        fmt.Printf("Username '%s' does not exist.\n", username)
    }
}

func main() {
    // Define command-line flags for the wordlist file and target URL
    wordlistPath := flag.String("wordlist", "", "Path to the wordlist file")
    targetURL := flag.String("url", "", "Target URL for username check (e.g., https://example.com/check-username)")
    flag.Parse()

    // Validate that a target URL was provided
    if *targetURL == "" {
        fmt.Println("Error: You must provide a target URL using the -url flag.")
        os.Exit(1)
    }

    // Open the wordlist file
    file, err := os.Open(*wordlistPath)
    if err != nil {
        fmt.Printf("Error opening wordlist file '%s': %v\n", *wordlistPath, err)
        os.Exit(1)
    }
    defer file.Close()

    // Read the file line by line and check each username
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        username := scanner.Text()
        if username == "" {
            continue // Skip empty lines
        }
        checkUsername(username, *targetURL)
    }

    // Check for any scanning errors
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading wordlist file: %v\n", err)
    }
}
