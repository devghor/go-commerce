# Go Commerce

A modern, high-performance e-commerce backend API built with **Go** and the **Fiber** framework, following clean architecture principles.  

⚠️ Currently in development.

---


## 🛠️ Tech Stack

- **Go** 1.24.2 - Programming language
- **Fiber v2** - Web framework
- **nodemon** - Development hot reload
- **godotenv** - Environment variable management

### Installation

1. Clone the repository:
```bash
git clone https://github.com/devghor/go-commerce.git
cd go-commerce
```

2. Install dependencies:
```bash
go mod tidy
npm install -g nodemon  # For hot reload during development
```

3. Set up environment variables:
```bash
cp .env.example .env  # Create your environment file
```

4. Configure your `.env` file:
```env
HTTP_PORT=3000
APP_ENV=dev
```

### Running the Application

#### Development Mode (with hot reload)
```bash
make server
```

#### Production Mode
```bash
go run main.go
```