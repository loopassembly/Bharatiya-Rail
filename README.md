# Railway Skills Development Platform

## Overview

The Railway Skills Development Platform is a comprehensive learning management system designed specifically for Indian Railways employees in lower grades. This platform facilitates skill development, knowledge assessment, and career advancement through structured learning modules and mock tests.

## Purpose

This platform addresses the need for accessible, targeted training for railway personnel who may not have extensive technical backgrounds but need to acquire specialized skills for railway operations and maintenance. The system helps users:

1. Learn essential railway technical concepts
2. Practice with realistic mock tests
3. Track learning progress
4. Prepare for job advancement opportunities

## Key Features

### User Authentication System
- Secure registration and login
- Password encryption using bcrypt
- JWT-based authentication

### Learning Management
- Structured learning modules for railway personnel
- Technical content specific to railway operations
- Progressive skill development tracks

### Assessment System
- Mock tests simulating real job scenarios
- Railway-specific technical assessments
- Performance tracking and improvement analytics

### Record Management
- Track training progress and completion
- Store assessment results
- Manage technical certifications and qualifications

## Technology Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: Fiber (Go web framework)
- **Database**: GORM with SQLite
- **Authentication**: JWT (JSON Web Tokens)
- **Password Security**: bcrypt encryption

### Deployment
- Containerized with Docker
- Deployable to cloud platforms

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/railway-skills-platform.git
cd railway-skills-platform
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (create a .env file):
```
DB_DSN=railway.db
JWT_SECRET=your_jwt_secret
```

4. Run the application:
```bash
go run main.go
```

The application will be available at `http://localhost:8080`.

## API Endpoints

### Authentication
- `POST /api/auth/register` - Create new user account
- `POST /api/auth/login` - Authenticate user and get token

### User Records
- `POST /api/users/records` - Create a new training record (requires authentication)
- `GET /api/users/records` - Get all training records for authenticated user

## Project Structure

```
railway-bac/
├── controllers/        # Request handlers
├── initializers/       # Application setup
├── middleware/         # Request processing
├── models/             # Data structures
├── routes/             # API endpoint definitions
├── testing/            # Mock tests
├── utils/              # Helper functions
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Contributing

This project was developed as part of an internship at Indian Railways. Contributions to enhance the platform's functionality are welcome.

## License

[Include your license information here]

## Acknowledgments

- Indian Railways for the internship opportunity
- All contributors to the open-source libraries used in this project