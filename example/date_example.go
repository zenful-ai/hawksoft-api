package main

import (
	"fmt"
	"time"

	hawksoft "github.com/zenful-ai/hawksoft-api/client"
)

// ExampleDateHandling demonstrates how to work with date fields in the HawkSoft API
func ExampleDateHandling() {
	fmt.Println("\n=== Date Handling Examples ===\n")

	// Example 1: Parsing dates from API responses
	fmt.Println("1. Parsing dates from API responses:")

	// Simulate a policy from the API with date strings
	policy := hawksoft.Policy{
		Carrier:        "Example Insurance Co",
		EffectiveDate:  hawksoft.StringPtr("2024-10-08T00:00:00"),
		ExpirationDate: hawksoft.StringPtr("2025-10-08T00:00:00"),
		LOBs:           []hawksoft.LOB{{Name: hawksoft.StringPtr("Auto")}},
		Status:         "Active",
		Type:           "Personal",
	}

	// Parse the dates to time.Time
	effectiveDate := hawksoft.DateTimePtrToTime(policy.EffectiveDate)
	expirationDate := hawksoft.DateTimePtrToTime(policy.ExpirationDate)

	fmt.Printf("   Effective Date: %s\n", effectiveDate.Format("January 2, 2006"))
	fmt.Printf("   Expiration Date: %s\n", expirationDate.Format("January 2, 2006"))
	fmt.Printf("   Days until expiration: %d\n", int(time.Until(expirationDate).Hours()/24))

	// Example 2: Creating dates for API requests
	fmt.Println("\n2. Creating dates for API requests:")

	now := time.Now()
	futureDate := now.AddDate(1, 0, 0) // One year from now

	// Convert to API format
	effectiveDateStr := hawksoft.FormatDateTime(now)
	expirationDateStr := hawksoft.FormatDateTime(futureDate)

	fmt.Printf("   Current time as API format: %s\n", effectiveDateStr)
	fmt.Printf("   One year from now: %s\n", expirationDateStr)

	// Example 3: Working with pointer fields
	fmt.Println("\n3. Creating policy with dates:")

	newPolicy := hawksoft.Policy{
		Carrier:        "Another Insurance Co",
		EffectiveDate:  hawksoft.TimeToDateTimePtr(now),
		ExpirationDate: hawksoft.TimeToDateTimePtr(futureDate),
		LOBs:           []hawksoft.LOB{{Name: hawksoft.StringPtr("Home")}},
		Status:         "Active",
		Type:           "Personal",
		Premium:        floatPtr(1200.50),
	}

	fmt.Printf("   Policy effective: %s\n", *newPolicy.EffectiveDate)
	fmt.Printf("   Policy expires: %s\n", *newPolicy.ExpirationDate)

	// Example 4: Handling nil/empty dates
	fmt.Println("\n4. Handling nil/empty dates:")

	var nilPolicy hawksoft.Policy
	nilPolicy.Carrier = "Test"
	nilPolicy.LOBs = []hawksoft.LOB{}
	nilPolicy.Status = "Pending"
	nilPolicy.Type = "Commercial"

	// Safe to call with nil pointers
	effectiveNil := hawksoft.DateTimePtrToTime(nilPolicy.EffectiveDate)
	fmt.Printf("   Nil date is zero: %v\n", effectiveNil.IsZero())

	// Example 5: Invoice dates (date-only format)
	// Note: Invoice dates use openapi_types.Date which handles the date-only format automatically
	fmt.Println("\n5. Invoice dates (date-only format):")
	fmt.Println("   Invoice dates use openapi_types.Date type")
	fmt.Println("   This type automatically handles 'YYYY-MM-DD' format serialization")

	fmt.Println("\nNote: The HawkSoft API uses two date formats:")
	fmt.Println("  - DateTime fields: '2024-10-08T00:00:00' (no timezone)")
	fmt.Println("  - Date fields: '2024-10-08' (invoices)")
	fmt.Println("  - RFC3339 fields: '2024-10-08T00:00:00Z' (timestamps in requests)")
}

func floatPtr(f float64) *float64 {
	return &f
}
