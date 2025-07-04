# 🧺 Laundry API

A modern, scalable API for managing laundry services, built with **Golang**, **Ent**, **Wire**, and **Docker**.

---

## ✨ Features

- ✅ RESTful & gRPC endpoints
- ✅ Authentication with JWT
- ✅ Role-based permissions
- ✅ Merchant & user management
- ✅ Seeding for features, users, merchants
- ✅ Clean DI with **Wire**
- ✅ Configurable with environment variables
- ✅ Docker-ready

---

## 🚀 Getting Started

### 📦 Requirements

- Go 1.21+
- [Ent](https://entgo.io/) ORM
- Docker & Docker Compose
- Wire for dependency injection

---

### ⚙️ Installation

Clone the repository:

```bash
git clone https://github.com/umardev500/go-laundry.git
cd go-laundry
```

Install dependencies:

```bash
go mod tidy
```

Generate Ent & Wire code:

```bash
go generate ./...
```

Run database migrations and seeds:

```bash
go run cmd/seed/main.go
```

---

### ▶️ Running Locally

Run the HTTP server:

```bash
go run cmd/laundry/main.go
```

Or run with Docker Compose:

```bash
docker-compose up --build
```

The API will be available at **http://localhost:3000**

---

## 📚 Project Structure

```plaintext
.
├── cmd
│   ├── laundry/           # API entry point
│   └── seed/              # Seed scripts
├── internal
│   ├── app/               # HTTP server setup
│   ├── config/            # Database & app configs
│   ├── constants/         # Features & permissions
│   ├── di/                # Dependency injection with Wire
│   ├── domain/            # Core domain logic
│   ├── handler/           # gRPC & HTTP handlers
│   ├── repository/        # Repositories
│   ├── seeds/             # Seed data
│   ├── service/           # Services
│   └── usecase/           # Business use cases
├── pkg/                   # Shared packages (JWT, context, response)
├── docker-compose.yml     # Docker services
├── Makefile               # Common tasks
├── go.mod                 # Go modules
├── go.sum                 # Dependencies checksum
└── tmp/                   # Temp files, build logs
```

---

## 📚 Environment Variables

Create a `.env` file in your project root.  
Here’s an example `.env.example`:

```env
# 🔐 JWT Configuration
JWT_SECRET=supersecretkey
JWT_EXPIRATION_HOURS=24

# 🌐 Server Configuration
PORT=3000

# 🗄️ Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=laundry_db
DB_SSLMODE=disable
```

---

---

## 🧩 Makefile

Use `make` commands for convenience:

```bash
make generate   # Run Ent and Wire generators
make run        # Run the API server
make seed       # Seed initial data
make build      # Build the project
```

---

## 📝 License

This project is licensed under the **MIT License**.  
Use it freely, improve it, share it.

---

## 🤝 Contributing

Pull requests are welcome!  
Please open an issue first to discuss major changes.

---

## 📫 Contact

Maintained by **[Your Name](https://github.com/yourusername)**.

---

**Happy Washing! 🧼✨**
