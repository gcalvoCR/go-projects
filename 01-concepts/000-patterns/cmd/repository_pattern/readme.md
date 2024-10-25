### What is the repository pattern?

The repository pattern is a design pattern often used to separate the business logic from the data access logic. It provides an abstraction layer between the business layer and the data layer (like a database), making the code more modular, testable, and easier to maintain.

Classes focused on encapsulating the logic necessary to interact witho out data with the following benefits:
- Centralize common data access functionality
- Better code maintainability
- Decouples from our infrastructure layer
- Makes our code more testable

## What the Repository Pattern Solves:
- `Separation of Concerns`: Keeps business logic separate from data access.
- `Testability`: Allows you to mock the data layer in tests.
- `Flexibility`: Makes it easier to switch between different data sources (e.g., SQL vs. NoSQL).

## Structure of the Repository Pattern:
- `Model`: Represents the data structure.
- `Repository Interface`: Defines the methods to interact with the data.
- `Repository Implementation`: Actual implementation of the repository methods.
