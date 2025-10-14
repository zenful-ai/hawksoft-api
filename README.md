# HawkSoft Partner API

OpenAPI 3.0 specification for the HawkSoft Partner API v3.0.

## API Documentation

- **API Reference**: https://partner.hawksoft.app/v3/api.html
- **Data Model Documentation**: https://partner.hawksoft.app/v3/model.html
- **Base URL**: https://partner.hawksoft.app/v3

## Authentication

The API uses HTTP Basic Authentication. All requests require:
- Basic authentication credentials
- `version=3.0` query parameter

## OpenAPI Specification

The OpenAPI specification is available in `openapi.yaml` and can be used with:
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [Redoc](https://redocly.com/redoc/)
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- Code generation tools like [OpenAPI Generator](https://openapi-generator.tech/)

## API Endpoints

### Agencies
- `GET /vendor/agencies` - List subscribed agencies
- `GET /vendor/agency/{agencyId}/offices` - List agency offices

### Clients
- `GET /vendor/agency/{agencyId}/clients` - Get changed clients
- `POST /vendor/agency/{agencyId}/clients` - Get multiple clients by client numbers
- `GET /vendor/agency/{agencyId}/client/{clientId}` - Get specific client details
- `GET /vendor/agency/{agencyId}/clients/search` - Search clients (in development)

### Client Actions
- `POST /vendor/agency/{agencyId}/client/{clientId}/log` - Create log note
- `POST /vendor/agency/{agencyId}/client/{clientId}/attachment` - Upload attachment
- `POST /vendor/agency/{agencyId}/client/{clientId}/receipts` - Record payment (HS6 only)

## Client Includes

When retrieving client data, you can optionally include additional information using the `include` query parameter:

- `Details` - Client details (agency number, business type, client type, etc.)
- `People` - People associated with the client
- `Contacts` - Contact information (phone, email, etc.)
- `Claims` - Claims associated with the client
- `Policies` - Policy information
- `Invoices` - Invoice data

Example: `GET /vendor/agency/123/client/456?version=3.0&include=Details&include=Policies`

## Generated Go Client

A Go client library has been generated from the OpenAPI specification using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen).

### Installation

```bash
go get github.com/zenful-ai/hawksoft-api/client
```

### Quick Start

```go
package main

import (
    "context"
    "net/http"

    hawksoft "github.com/zenful-ai/hawksoft-api/client"
)

func main() {
    // Create client
    client, err := hawksoft.NewClient("https://partner.hawksoft.app/v3")
    if err != nil {
        panic(err)
    }

    // Set up Basic Authentication
    basicAuth := func(ctx context.Context, req *http.Request) error {
        req.SetBasicAuth(username, password)
        return nil
    }

    // Make API call
    resp, err := client.GetAgencies(context.Background(),
        &hawksoft.GetAgenciesParams{Version: "3.0"},
        basicAuth)
}
```

### Running the Example

```bash
# Run without credentials (demonstration mode)
go run example/main.go

# Run with credentials
export HAWKSOFT_USERNAME=your_username
export HAWKSOFT_PASSWORD=your_password
go run example/main.go
```

### Regenerating the Client

If you update the OpenAPI spec, regenerate the client with:

```bash
oapi-codegen --config=config.yaml openapi.yaml
```

## Support

For HawkSoft Partner API support, visit: https://partner.hawksoft.app
