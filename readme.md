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