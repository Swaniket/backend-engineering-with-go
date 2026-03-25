## Folder Structure

###### `code/bin`
- This is where the compiled binary files lives.

###### `code/cmd`
- This is the entry point of the application

###### `code/cmd/api`
- Here it will have anything related to our transpot layer / HTTP / handlers / server related things

###### `code/cmd/migrate`
- Here we will have the migration folder if we have our custom migrations

###### `code/cmd/migrate/migrations`
- Will have the migrations here

###### `code/internal`
- All of the internal packages, means not to be exported to our API server
- This package doesn't know anything about the outside.
- This will include the storage layer that we will use for postgres
- Data validation / Sending emails / package implementation / Rate limiter

###### `code/internal/env`
- This will have the functions to access the keys from .envrc file

###### `code/internal/store`
- This is where the repositories / data will stay (Database methods)

###### `code/docs`
- Auto generated docs from swagger will live here.

###### `code/scripts`
- We can have some scripts for setting up our server

## Design Principles

### Separation of Concerns
- Each level in your program should be separate by a clear barrier, the `transport layer`, the `service layer`, the `storage layer` etc.

### Dependency Inversion Principle (DIP)
- You're injecting the dependencies in your layer. You don't directly call them.
- Why? Because it promotes loose coupling and makes it easier to test your program.

### Adaptability to change
- By organizing your code in a modular and flexible way, you can more easily introduce new features, refactor existing code, and respond to evolving business requirements.

- Your systems should be easy to change, if you have to change a lot of existing code to add a new feature, you are doing it wrong.