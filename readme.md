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

## MySQL Schema

### Tables and Relationships

#### Beneficiaries

| Column         | Type         | Description                   |
|----------------|--------------|-------------------------------|
| id             | INT          | Primary Key, Auto Increment   |
| name           | VARCHAR(255) | Name of the beneficiary        |
| account_number | VARCHAR(255) | Account number of the beneficiary |

#### Transfers

| Column         | Type         | Description                         |
|----------------|--------------|-------------------------------------|
| id             | INT          | Primary Key, Auto Increment         |
| beneficiary_id | INT          | Foreign Key referencing beneficiaries(id) |
| amount         | DECIMAL(10, 2)| Amount of the transfer              |
| status         | VARCHAR(50)  | Status of the transfer              |

#### TransactionHistories

| Column         | Type         | Description                             |
|----------------|--------------|-----------------------------------------|
| id             | INT          | Primary Key, Auto Increment             |
| transfer_id    | INT          | Foreign Key referencing transfers(id)   |
| timestamp      | DATETIME     | Timestamp of the transaction            |
| status         | VARCHAR(50)  | Status of the transaction               |

### Relationships

- **Beneficiaries to Transfers**
  - `beneficiaries.id` (Primary Key) is referenced by `transfers.beneficiary_id` (Foreign Key).

- **Transfers to TransactionHistories**
  - `transfers.id` (Primary Key) is referenced by `transaction_histories.transfer_id` (Foreign Key).

## Design Patterns and Decisions

### Repository Pattern

The Repository Pattern is used to separate the data access logic from the business logic. This pattern abstracts the data access layer, allowing the application to interact with the data through a set of well-defined interfaces. By using the Repository Pattern, we achieve the following benefits:

- **Separation of Concerns:** The data access logic is decoupled from the business logic, making the codebase more maintainable and easier to understand.
- **Testability:** It becomes easier to mock or stub the data access layer in unit tests, improving the testability of the business logic.
- **Flexibility:** Changes to the data source or data access logic do not affect the business logic. The repository can be modified or replaced without impacting other parts of the application.

In the application, repositories are defined for each entity:

- **`repositories/beneficiary.go`**: Manages operations related to the `Beneficiary` model.
- **`repositories/transactionHistory.go`**: Handles operations for the `TransactionHistory` model.
- **`repositories/transfer.go`**: Manages operations for the `Transfer` model.

### Separation of Concerns

The application architecture is designed with a clear separation of concerns to keep the code organized and structured. This separation improves maintainability, readability, and scalability. The key areas of separation are:

- **Database Initialization vs. Model Creation:**
  - **Database Initialization:** Managed in **`database/connection.go`**, responsible for setting up and managing the database connection.
  - **Model Creation:** Models are defined in **`models/models.go`**, separated from the database initialization logic. This ensures changes to the database schema or connection logic do not impact the model definitions.

- **Routes vs. Route Handlers:**
  - **Routes:** Configured in **`router/router.go`**, which sets up HTTP routes and initializes the router. It routes incoming requests to the appropriate handlers.
  - **Route Handlers:** Implemented in **`router/handlers.go`**, defining HTTP handlers that process requests and return responses. Handlers interact with repositories to perform business logic.

- **Route Handlers vs. Repositories:**
  - **Route Handlers:** Contain logic to handle HTTP requests, validate input, and format responses. They delegate data access operations to the repositories.
  - **Repositories:** Implemented in the **`repositories/`** folder, encapsulating the logic for interacting with the database. This layer provides a clean interface for route handlers to access and modify data.

### Benefits of Separation of Concerns

- **Organized Codebase:** Each component has a specific responsibility, making the codebase easier to navigate and understand.
- **Enhanced Maintainability:** Changes to one part of the system (e.g., database schema or route logic) can be made with minimal impact on other parts.
- **Improved Scalability:** The modular structure allows for easy expansion or modification of individual components without disrupting the entire system.

### Atomic Transactions

To ensure data consistency and integrity, especially when dealing with multiple related operations, atomic transactions are used. Atomic transactions guarantee that a series of database operations are executed as a single unit of work. If any part of the transaction fails, the entire transaction is rolled back, leaving the database in a consistent state.

In the context of creating a transfer, we use atomic transactions to:

- **Ensure Consistency:** When creating a transfer, both the `Transfer` record and the associated `TransactionHistory` record must be created successfully. If an error occurs while creating either record, the entire transaction is rolled back.
- **Maintain Integrity:** This approach ensures that the database is not left in an inconsistent state if any operation within the transaction fails.

The atomic transaction is implemented in the `CreateTransfer` function as follows:

```go
func CreateTransfer(newTransfer *models.Transfer) error {
	// Use an atomic transaction to create both the transfer and transaction history
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var beneficiary models.Beneficiary
		tx.Model(&models.Beneficiary{}).
			Where("id = ?", newTransfer.BeneficiaryId).
			First(&beneficiary)

		err := tx.Model(&beneficiary).Association("Transfers").Append(newTransfer)
		if err != nil {
			return err
		}

		var newTransactionHistory models.TransactionHistory
		newTransactionHistory.TransferID = newTransfer.ID
		newTransactionHistory.Timestamp = time.Now()
		newTransactionHistory.Status = "SENT"

		err = tx.Model(&models.TransactionHistory{}).Create(&newTransactionHistory).Error
		if err != nil {
			return err
		}

		newTransfer.TransactionHistory = newTransactionHistory

		// return nil will commit the whole transaction
		return nil
	})

	return err
}
