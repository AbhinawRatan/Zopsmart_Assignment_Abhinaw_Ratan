# Vinyl Records Management API

A simple HTTP (REST) API for managing a vinyl record collection using the GoFr framework.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Unit Tests](#unit-tests)
- [Bonus Features](#bonus-features)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project implements a CRUD API for managing a collection of vinyl records. It is built using the GoFr framework, providing a clean and structured approach to develop HTTP APIs.

## Features

- **List Records:** Retrieve the list of vinyl records.
- **Add Record:** Add a new record to the collection.
- **Update Record Condition:** Update the condition of a record in the collection.
- **Remove Record:** Remove a record from the collection.

## Getting Started

To get started with this project:

1. Clone the repository:

2. Navigate to the project directory:
3. Run the application:

4. Test the API using your preferred method (e.g., Postman, curl).

## API Endpoints

- **List Records**
- `GET /records`
- Retrieves a list of all records in the collection.

- **Add Record**
- `POST /records`
- Adds a new record to the collection.
- Request body example:
 ```json
 {
   "AlbumTitle": "Kind of Blue",
   "Artist": "Miles Davis",
   "ReleaseYear": 1959,
   "Genre": "Jazz",
   "Condition": "New",
   "InStock": true
 }
 ```

- **Update Record Condition**
- `PUT /records/:id`
- Updates the specified record's condition.
- Request body example:
 ```json
 {
   "Condition": "Good"
 }
 ```

- **Remove Record**
- `DELETE /records/:id`
- Removes the specified record from the collection.

## Unit Tests

The project includes a set of unit tests to ensure that each API endpoint operates correctly. To run the tests, execute the following command in the terminal:

 go run main.go

The tests cover the following:

- **List Records Test:** Verifies that the `GET /records` endpoint correctly retrieves the list of records.
- **Add Record Test:** Checks that the `POST /records` endpoint successfully adds a new record and returns the correct JSON response with the new record details.
- **Update Record Test:** Ensures that the `PUT /records/:id` endpoint accurately updates a record's condition based on the provided ID.
- **Remove Record Test:** Confirms that the `DELETE /records/:id` endpoint removes the correct record from the collection.

These tests mock the data layer and focus on the HTTP layer to validate responses, status codes, and error handling.

## Bonus Features

- **Search Functionality:** Ability to filter the records list by artist, album title, genre, or year.
- **Pagination:** Support for paginating the list of records to manage large collections efficiently.
- **Record History:** An additional feature to track the history of each record's condition over time.

## Contributing

Feel free to contribute to this project. Fork the repository, make your changes, and submit a pull request for review.

## License

This project is licensed under the MIT License. See the LICENSE file for full details.

