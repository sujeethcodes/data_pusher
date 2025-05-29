# Data Pusher (Golang + Echo)

A lightweight API service that manages Accounts and Destinations, and forwards incoming JSON data to dynamic destinations based on account secret tokens. Built using **Golang**, **Echo**, and **MySQL** with modular architecture and singleton DB connection.

## ✨ Features

- Create, retrieve, and delete accounts
- Add, list, and delete destinations for each account
- Forward incoming data to all destination URLs associated with an account (based on secret token)
- Custom headers and HTTP method support for each destination
- Modular codebase (MVC-style) with clear separation of concerns

---

## 📁 Folder Structure

├── cmd/
│ └── main.go # Entry point
├── controller/ # HTTP handlers
│ ├── account_controller.go
│ ├── destination_controller.go
│ └── data_controller.go
├── usecase/ # Business logic layer
│ ├── account_usecase.go
│ ├── destination_usecase.go
│ └── data_forwarder.go
├── repository/
│ └── mysql.go # Singleton DB connection
├── entity/ # Data models (structs)
│ ├── account.go
│ └── destination.go
├── utils/ # Helper functions
│ └── parse_headers.go
├── .env # Environment variables (DB config)
└── go.mod



---

## 🧪 API Endpoints

### Account APIs

| Method | Endpoint           | Description           |
|--------|--------------------|-----------------------|
| POST   | `/accounts`        | Create an account     |
| GET    | `/accounts/:account_id`    | Get account by ID     |
| DELETE | `/accounts/`    | Delete account        |

### Destination APIs

| Method | Endpoint                          | Description                |
|--------|-----------------------------------|----------------------------|
| POST   | `/destinations`                   | Add a new destination      |
| GET    | `/destinations/:account_id`      | List destinations by account ID |

### Data Forwarding API

| Method | Endpoint              | Description                  |
|--------|-----------------------|------------------------------|
| POST   | `/server/incoming_data` | Forward data to destinations |

**Headers:**  
`CL-X-TOKEN`: Secret token of the account.

---

## 🛠️ Setup Instructions

1. **Clone the repo**

git clone https://github.com/yourusername/data-pusher.git
cd data-pusher
go mod tidy
Configure .env
Run the app
go run cmd/main.go
App runs at: http://localhost:8080
`Auth & Security`
Token-based request authentication (CL-X-TOKEN)
`Data is only forwarded if the secret matches a registered account`

` Dependencies`
Echo
MySQL Driver
UUID
godotenv (optional for .env loading)



