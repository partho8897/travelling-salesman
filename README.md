# Travelling Salesman API

This is a sample project demonstrating a RESTful API in Go for managing travelling salesmen. The API allows for the registration, login, logout, retrieval of salesman details, and tracking unique destinations travelled by salesmen. It uses in-built memory storage and does not include authentication for simplicity.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- Register a salesman
- Login/logout functionality
- Get salesman details
- Retrieve active salesmen
- Track unique destinations travelled by salesmen

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22 or higher

### Installing

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/travelling-salesman-api.git
    cd travelling-salesman-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## API Endpoints

### Register a Salesman

- **URL:** `/sales/person/register`
- **Method:** `POST`

### Login

- **URL:** `/sales/person/:personId/login`
- **Method:** `POST`

### Logout

- **URL:** `/sales/person/:personId/logout`
- **Method:** `POST`

### Get Salesman Details

- **URL:** `/sales/person/:personId`
- **Method:** `GET`

### Get Active Salesmen

- **URL:** `/sales/analytics/get_active_sales_person`
- **Method:** `GET`

### Get Total Unique Destinations

- **URL:** `/sales/analytics/get_unique_destinations`
- **Method:** `GET`

## Usage

To interact with the API, you can use tools like `curl`, Postman, or any HTTP client of your choice.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.
