# Cats Social

Cats Social is an app that allows cat owners to match their cats with each other.

![It's a cat](https://media.giphy.com/media/jJjb9AUHOiP3nJJMdy/giphy.gif)

## 📜Table of Contents

- [Cats Socials - Tree Planting Backend](#cats-social)
  - [📜Table of Contents](#table-of-contents)
  - [🔍Requirements](#requirements)
  - [🛠️Installation](#️installation)
  - [🌟Features](#features)
  - [🚀Usage](#usage)
  - [⚙️Configuration](#️configuration)
  - [💾Database Migration](#database-migration)

## 🔍Requirements

This application requires the following:

- Golang: https://go.dev/dl/
- PostgreSQL: https://www.postgresql.org/download/
- Golang Migrate: https://github.com/golang-migrate/migrate

## 🛠️Installation

To install the boilerplate, follow these steps:

1. Make sure you have Golang, PostgreSQL, and Golang Migrate installed and configured on your system.

2. Clone this repository:

   ```bash
   git clone https://github.com/BennoAlif/ps-cats-social.git
   ```

3. Navigate to the project directory:

   ```bash
   cd ps-cats-social
   ```

4. Run the following command to install dependencies:
   ```bash
   go mod download
   ```

## 🌟Features

Cats Social offers the following features:

- **Authentication**:
  - User registration
  - User login
- **Cat Management (CRUD)**:
  - Create new cat profiles
  - View existing cat profiles
  - Update cat profiles
  - Delete cat profiles
- **Matching**:
  - Match your cat with other cats
  - View matching cats
  - Approve or reject matches
  - Delete matches

## 🚀Usage

1. **Setting Up Environment Variables**

   Before starting the application, you need to set up the following environment variables:

   - `DB_NAME`: Name of your PostgreSQL database
   - `DB_PORT`: Port of your PostgreSQL database (default: 5432)
   - `DB_HOST`: Hostname or IP address of your PostgreSQL server
   - `DB_USERNAME`: Username for your PostgreSQL database
   - `DB_PASSWORD`: Password for your PostgreSQL database
   - `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)
   - `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)
   - `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)

2. **Database Migrations**

   Cats Social uses Golang Migrate to manage database schema changes. Here's how to work with migrations:

   - Apply migrations to the database:

     ```bash
     migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
     ```

3. **Running the Application**

   Once you have set up the environment variables, you can start the application by running:

   ```bash
   go run src/main.go
   ```

   This will start the Cats Social application on the default port (usually 8080). You can access the application in your web browser at http://localhost:8080

## ⚙️Configuration

The application uses environment variables for configuration. You can configure the database connection, JWT secret, and bcrypt salt by setting the following environment variables:

- Refer to the [Usage](#usage) section for a detailed explanation of each environment variable.

## 💾Database Migration

Database migration must use golang-migrate as a tool to manage database migration

1. Direct your terminal to your project folder first
2. Initiate folder
   ```bash
   mkdir db/migrations
   ```
3. Create migration

   ```bash
   migrate create -ext sql -dir db/migrations add_user_table
   ```

   This command will create two new files named `add_user_table.up.sql` and `add_user_table.down.sql` inside the `db/migrations` folder

   - `.up.sql` can be filled with database queries to create / delete / change the table
   - `.down.sql` can be filled with database queries to perform a `rollback` or return to the state before the table from `.up.sql` was created

4. Execute migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
   ```

5. Rollback migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations dow
   ```

6. View the current migration state
   ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" version
   ```
