# SkillSync Protos

## Overview
This repository contains Protocol Buffer (protobuf) definitions for the SkillSync microservices architecture. It serves as the central source of truth for service interfaces, ensuring consistent communication between different services.

## Purpose
- Define service interfaces and message types using Protocol Buffers
- Generate client and server code for different services
- Ensure type safety and consistency across microservices
- Provide a single source of truth for API contracts

## Structure
- **Auth/**: Protocol definitions for the Auth Service
  - User registration and authentication
  - Profile management
  - Token validation

- **Job/**: Protocol definitions for the Job Service
  - Job posting and management
  - Job application handling
  - Job search functionality

- **chat/**: Protocol definitions for the Chat/Notification Service
  - Real-time messaging
  - Notifications

- **gen/**: Generated code from the protobuf definitions
  - Go client and server code

## Usage
1. Define or update service interfaces in `.proto` files
2. Generate code for your target language (Go, etc.)
3. Import the generated code in your services

## Important Notes
- All job_id fields use string type (not uint64) to ensure consistency between protobuf definitions and Go implementation
- Response structures are kept minimal, particularly for authentication responses
- Login responses include only ID and success message, not tokens
- Authentication tokens must be handled in the format "Bearer {token}"

## Generating Code
To regenerate code after modifying proto files:

```bash
# For Go code generation
protoc --go_out=. --go-grpc_out=. path/to/your.proto
```

## Service Integration
- Auth Service runs on port 50051
- Job Service runs on port 50052
- Services should be configured to connect to the correct ports for proper communication
