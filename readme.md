# SCT System

This project is a simplified social cash transfer (SCT) disbursement process system. It includes a backend built with Go and a frontend built with Svelte. The backend interacts with a MySQL database to store beneficiary information, transfer details, and transaction history.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/chikuBanda/go-svelte-sct-system
cd go-svelte-sct-system
```
### Run the application
```sh
docker-compose up --build
```

## Services
The following services will be running:

### MySQL Database
* URL: localhost:3306
* User: sct-user
* Password: insecure_password
* Database: sct_system_db

### Backend
* URL: http://localhost:8080
* Description: The backend service is responsible for handling API requests related to beneficiaries, transfers, and transaction history.

### Frontend
* URL: http://localhost:9000
* Description: The frontend service provides a simple UI to interact with the backend APIs.

## API Endpoints
The backend exposes the following API endpoints:

#### 1. Register a New Beneficiary
* URL: POST /register_beneficiary
* Description: Registers a new beneficiary.
* Example Request Body:
    ```shell
    {
      "name": "Sibo",
      "account_number": "515151528"
    }
    ```

#### 2. Initiate a Transfer
* URL: POST /initiate_transfer
* Description: Initiates a new transfer.
* Example Request Body:
    ```shell
    {
      "beneficiary_id": 1,
      "amount": 1000
    }
    ```

#### 3. Check Transfer Status
* URL: GET /check_transfer_status/[transfer_id]
* Description: Checks the status of a specific transfer.

#### 4. Get Transaction History for a Beneficiary
* URL: GET /get_transaction_history
* Description: Checks the status of a specific transfer.
* Example Request Body:
    ```shell
    {
      "beneficiary_id": 1
    }
    ```

#### 4. List of beneficiaries
* URL: GET /get_beneficiaries
* Description: Returns a list of all beneficiaries.

## Architecture of the Go Application
The Go application for the SCT (Social Cash Transfer) system is designed with a clear separation of concerns, leveraging Go's strengths in simplicity. The architecture consists of several key components:

### Components
1. Models
2. Database Layer
3. Repository Layer
4. Middleware
5. Router and Handlers
6. Utility Functions
7. Main Application

### 1. Models
The models define the structure of the data stored in the database. They are represented as Go structs and are mapped to database tables using GORM (Go Object-Relational Mapping).

* models/models.go
  * Contains the struct definitions for Beneficiary, Transfer, and TransactionHistory.

### 2. Database Layer
The database layer is responsible for establishing a connection to the MySQL database. It uses GORM to handle the connection and database operations.

* database/connection.go
    * Contains the logic to initialize and manage the database connection.

### 3. Repository Layer
The repository layer abstracts the data access logic and provides methods to perform CRUD operations on the models. This layer interacts directly with the database.

* repositories/beneficiary.go
  * Contains functions to manage beneficiary data.

* repositories/transactionHistory.go 
  * Contains functions to manage transaction history data.
  
* repositories/transfer.go 
  * Contains functions to manage transfer data.

### 4. Middleware
Middleware functions provide a way to filter and process incoming requests before they reach the handlers. They are typically used for tasks like logging, authentication, and CORS (Cross-Origin Resource Sharing) handling.

* middleware/cors.go
  * Contains the logic to handle CORS, allowing the backend to respond to requests from different origins.

### 5. Router and Handlers
The router is responsible for routing HTTP requests to the appropriate handlers. Handlers define the HTTP endpoints and map them to the corresponding service functions, handling request validation, response formatting, and error handling.

* router/router.go 
  * Sets up the HTTP routes and initializes the router.
  
* router/handlers.go 
  * Defines the HTTP handlers for different endpoints.

### 6. Utility Functions
Utility functions provide common functionality that can be reused across different parts of the application.

* utils/functions.go 
  * Contains utility functions that support various parts of the application.

### 7. Main Application 
The main application initializes all components and starts the HTTP server. It performs the following tasks:

* Initialize the database connection 
* Set up the router and define routes 
* Start the HTTP server 
* cmd/main.go 
  * The entry point of the application. It initializes the database connection, sets up the router, and starts the HTTP server.

### Folder structure
```plaintext
sct-system/
├── Dockerfile
├── go.mod
├── go.sum
├── structure.txt
├── .idea/
│   └── workspace.xml
├── cmd/
│   └── main.go
├── database/
│   └── connection.go
├── middleware/
│   └── cors.go
├── models/
│   └── models.go
├── repositories/
│   ├── beneficiary.go
│   ├── transactionHistory.go
│   └── transfer.go
├── router/
│   ├── handlers.go
│   └── router.go
└── utils/
└── functions.go
```