package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	hawksoft "github.com/zenful-ai/hawksoft-api/client"
)

func main() {
	// Get credentials from environment variables
	username := os.Getenv("HAWKSOFT_USERNAME")
	password := os.Getenv("HAWKSOFT_PASSWORD")

	if username == "" || password == "" {
		fmt.Println("HAWKSOFT_USERNAME and HAWKSOFT_PASSWORD environment variables not set.")
		fmt.Println("This example will show the API structure without making actual calls.")
		fmt.Println("\nTo test with real credentials, set these environment variables and run again.")
		demonstrateAPI()
		return
	}

	// Create a new client
	apiClient, err := hawksoft.NewClient("https://partner.hawksoft.app/v3")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a request editor for Basic Auth
	basicAuth := func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}

	ctx := context.Background()

	// Example 1: Get list of agencies
	fmt.Println("Fetching agencies...")
	resp, err := apiClient.GetAgencies(ctx, &hawksoft.GetAgenciesParams{
		Version: "3.0",
	}, basicAuth)
	if err != nil {
		log.Printf("Error fetching agencies: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode == 200 {
			var agencies []int
			if err := json.Unmarshal(body, &agencies); err != nil {
				log.Printf("Error parsing response: %v\n", err)
			} else {
				fmt.Printf("Found %d agencies: %v\n", len(agencies), agencies)
			}
		} else {
			fmt.Printf("API returned status %d: %s\n", resp.StatusCode, string(body))
		}
	}

	fmt.Println("\nFor more examples, see the code comments in this file.")
}

func demonstrateAPI() {
	fmt.Println("\n=== HawkSoft Partner API Client Example ===\n")

	fmt.Println("Available Operations:")
	fmt.Println("1. GetAgencies() - Get list of subscribed agencies")
	fmt.Println("2. GetAgencyOffices(agencyId) - Get offices for an agency")
	fmt.Println("3. GetChangedClients(agencyId) - Get clients changed since a date")
	fmt.Println("4. GetClient(agencyId, clientId) - Get specific client details")
	fmt.Println("5. GetClientList(agencyId, clientNumbers) - Get multiple clients")
	fmt.Println("6. SearchClients(agencyId, policyNumber) - Search for clients")
	fmt.Println("7. CreateLogNote(agencyId, clientId, note) - Create a log note")
	fmt.Println("8. UploadAttachment(agencyId, clientId, file) - Upload attachment")
	fmt.Println("9. CreateReceipts(agencyId, clientId, receipts) - Record payments")

	fmt.Println("\nExample Usage:")
	fmt.Println(`
	// Create client
	client, _ := hawksoft.NewClient("https://partner.hawksoft.app/v3")

	// Set up authentication
	basicAuth := func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}

	// Make API call
	resp, err := client.GetAgencies(ctx, &hawksoft.GetAgenciesParams{
		Version: "3.0",
	}, basicAuth)

	// Parse response
	body, _ := io.ReadAll(resp.Body)
	var agencies []int
	json.Unmarshal(body, &agencies)
	`)

	fmt.Println("\nData Models Available:")
	fmt.Println("- ClientData: Complete client information")
	fmt.Println("- ClientDetails: Client business details")
	fmt.Println("- Policy: Insurance policy information")
	fmt.Println("- Person: Person associated with client")
	fmt.Println("- Contact: Contact information")
	fmt.Println("- Claim: Insurance claim")
	fmt.Println("- Invoice: Billing invoice")
	fmt.Println("- Office: Agency office location")
}
