# Rental Property API

A RESTful Rental Property API built using **Go** and **Beego Framework**.

The API loads rental property data from a JSON file, stores it in memory, and provides endpoints to retrieve properties with filtering support.

---

# Setup Instructions

## 1. Clone Repository

```bash
git clone https://github.com/oneakash/rental-property-api.git

cd rental-property-api
````

---

## 2. Install Dependencies

```bash
go mod tidy
```

---

## 3. Install Beego CLI

```bash
go install github.com/beego/bee/v2@latest
```

Verify:

```bash
bee version
```

---

# Configuration

Application configuration is available at:

```
conf/app.conf
```

Example:

```conf
appname = rental-property-api
httpport = 8080
runmode = dev
copyrequestbody = true
autorender = false
```

---

# Run Application

Start the Beego server:

```bash
bee run
```

Server will start:

```
http://localhost:8080
```

---

# API Endpoints

## 1. Get All Properties

### Request

```http
GET /v1/properties
```

Example:

```bash
curl http://localhost:8080/v1/properties
```

---

## 2. Get Property By ID

### Request

```http
GET /v1/properties/:id
```

Example:

```bash
curl http://localhost:8080/v1/properties/BC-1000001
```

---

# Query Parameters

## Limit

Return maximum number of properties.

Example:

```bash
curl "http://localhost:8080/v1/properties?limit=5"
```

---

# Filtering Examples

## Price Range

```bash
curl "http://localhost:8080/v1/properties?min_price=50&max_price=200"
```

---

## Feed Filter

```bash
curl "http://localhost:8080/v1/properties?feed=11"
```

---

## Published Filter

```bash
curl "http://localhost:8080/v1/properties?published=true"
```

---

## Property Type Filter

Supported values:

```
Hotel
House
Apartment
Villa
Resort
Hostel
```

Example:

```bash
curl "http://localhost:8080/v1/properties?property_type=Hotel"
```

---

## Bedroom Filter

```bash
curl "http://localhost:8080/v1/properties?min_bedroom=3"
```

---

## Rating Filters

Minimum star rating:

```bash
curl "http://localhost:8080/v1/properties?min_star_rating=4"
```

Minimum review score:

```bash
curl "http://localhost:8080/v1/properties?min_review_score=8"
```

Minimum reviews:

```bash
curl "http://localhost:8080/v1/properties?min_reviews=100"
```

---

## Amenities Filter

Amenities use OR logic.

Example:

```bash
curl "http://localhost:8080/v1/properties?amenities=Internet,Parking"
```

Returns properties having:

```
Internet OR Parking
```

---

# Error Responses

## Invalid Parameter

Example:

```bash
curl "http://localhost:8080/v1/properties?min_price=abc"
```

Response:

```json
{
    "Error": "invalid min_price"
}
```

HTTP Status:

```
400 Bad Request
```

---

## Property Not Found

Example:

```bash
curl http://localhost:8080/v1/properties/invalid-id
```

Response:

```json
{
    "Error": "Property not found"
}
```

HTTP Status:

```
404 Not Found
```

---

# Run Tests

Execute:

```bash
go test ./... -v
```

---

# Static Analysis

Run:

```bash
go vet ./...
```
