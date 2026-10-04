package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// healthcheckURL is the Hub's own health endpoint, reached over loopback.
const healthcheckURL = "http://127.0.0.1:" + httpPort + "/api/v1/health"

// runHealthcheck queries the health endpoint and returns the process exit code:
// 0 when the Hub answers 200, 1 otherwise. It backs `watchflare-hub healthcheck`,
// which the container image needs because it has no shell, curl or wget.
func runHealthcheck(url string) int {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "healthcheck failed: %v\n", err)
		return 1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "healthcheck failed: status %d\n", resp.StatusCode)
		return 1
	}
	return 0
}
