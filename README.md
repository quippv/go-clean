# go-clean

*go-clean* is a well-structured Go bootstrap project that embraces [Clean Architecture principles](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). It provides modular maintanable and scalable framework for building Go applications, emphasizing code organization, separation of concerns, and testability. This repository serves as a strong foundation for developers looking to build robust applications with Go, while maintaining clarity and simplicity in the codebase.

## Features
- [*Clean Architecture*](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html): Ensures separation of concerns, promoting decoupling between business logic and external interfaces, making the codebase easy to maintain and scale.

- [*Dockerized Setup*](https://docs.docker.com/compose/): A `docker-compose.yml` file is provided to simplify the setup of the application environment, ensuring concistency across different systems.

- [*Task Automation*](https://taskfile.dev/): Using `Taskfile.yml`, repetitive tasks like start the project or running tests are automated, improving development efficiency

- [*Air Reload for Development*](https://github.com/air-verse/air): Configured with `.air.toml`, enabling live reloading of the Go application during development. Every change to the code triggers a rebuild and reload, ensuring a smoother development experience.

- [*Goose for Migration*](https://github.com/pressly/goose): Goose is a powerful tool for managing database schema changes in Go applications, allowing incremental migrations written in either SQL or Go. By tracking migration versions, it ensures that changes are applied in sequence without errors, enabling smooth transitions between schema updates across development, staging, and production environments. This makes managing evolving database structures straightforward and reliable, especially when collaborating in teams or deploying to different environments.

- [*Swag for API Documentation*](https://github.com/swaggo/swag): Swag automates the generation of OpenAPI (Swagger) documentation for Go projects by converting code annotations into complete API documentation. It simplifies the process of keeping your API documentation in sync with the codebase, ensuring that developers and external consumers always have access to up-to-date specs. Swag’s generated documentation can be served alongside the application, allowing interactive exploration of the API endpoints, making testing, debugging, and integration much easier.

## Project Structure

- `cmd/`: Contains the entry points of the application. This where the `main` function resides, defining how the application is initiated, whether it's a web server, CLI tool or othe services.

- `internal/`: This is the core of the apllication. The `internal` directory encapsulates the domain logic and contains several submodules:
    - `entity/`: Defines the fundamental business entities, domains, rules, and interfaces. The entitites here repesent the core objects that drive business logic, isolated from the infrastructure layer.
    - `usecase/`: Contains use cases or aplication services that coordinate domain objects and implement the business logic. The services handle workflows like creating, retrieving, or updating entities.
    - `repository/`: Implements the repository pattern, where persistence logic is abstracted. It handles interactions with databases or external APIs while keeping business logic isolated.
    - `handler/`: Defines the HTTP, HTML File, gRPC or CLI handlers, acting as the entry pointto the system's interface layer. These handlers interact with the use case layer and translate incoming request into application actions.
    - `dto/`: Stands for "Data Transfer Objects" and is used to define simple objects that are passed between layers (tipycally between the handlers and use cases). DTOs encapsulate and structure the data exchanged between the application layers, ensuring separation of concerns and preventing direct exposure of domain entities in external interfaces.

- `configs/`: This directory contains configuration files and environment settings that drive the behavior of the application across various environments (development, staging, production).

- `migrations/`: Contains migration files for evolving the database schema in a version-controlled manner. It ensures that changes to the database structure are systematically tracked and reproducible across different environments.

- `utils/`: Provides various utility functions or helpers that can be reused across different parts of the application. These include common operations like id generating, string manipulations, error handling, or date-time processing.


## Getting Started

TO begin working with *go-clean*, follow these simple steps:

1. *Clone the Repository*:
```bash
git clone https://github.com/quippv/go-clean.git
```
2. *Install Goose, Air, Swag, Taskfile*:
```bash
go install github.com/air-verse/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/swaggo/swag/cmd/swag@latest
```
    - Taskfile: [https://taskfile.dev/installation/](https://taskfile.dev/installation/)
3. *Setup the Aplication with Docker*:
Ensure Docker is installed and running, then use:
```bash
task set:up
```
4. *For Development*:
To enable live reloading during development, you can use the `air` tool:
```bash
task start
```
5. *For Test All*:
To run test all go code you can use:
```bash
task test:all
```
6. *For Migration*:
You can see the migration script on `Taskfile.yml`

## Why Clean Architecture?

Clean Architecture separates code into layers, which insolates the business logic from the technical concerns. This approach brings numerous advantages:
- *Testability*: Since each layer is independent, unit testing becomes easier and more comprehensive.
- *Maintanability*: Changes in the external layers (like database or UI) don't affect the core business logic.
- *Scalability*: As the application grows, the architecture remains well-structured and modular, making it easier to manage.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests if you find bugs or have ideas to enhance the project. All forms of contribution are appreciated as we aim to build a clean, efficient Go application architecture.

