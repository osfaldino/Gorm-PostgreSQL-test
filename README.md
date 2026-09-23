# Gorm-PostgreSQL-test

A simple Go CLI application for managing students and academic groups using GORM and PostgreSQL.

## Features
- CRUD operations for Students (View, Add, Edit, Delete).
- CRUD operations for Groups (View, Add, Edit, Delete).
- Interactive terminal menu.

## Database Schema

### Group
- `id` (Primary Key, Auto-increment)
- `group_name` (string)

### Student
- `id` (Primary Key, Auto-increment)
- `first_name` (string, not null)
- `last_name` (string, not null)
- `email` (string, unique)
- `group_id` (integer)

## How to Run

1. **Configure DB**: Check and update your PostgreSQL connection string in `main.go`:
   ```go
   dsn := "host=localhost user=postgres password=1234567890 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
   ```
2. **Install dependencies**:
   ```bash
   go mod tidy
   ```
3. **Run the app**:
   ```bash
   go run main.go
   ```

## Menu Options
Enter a number from 0 to 8 when prompted:
- `1` / `2` — Show students / groups
- `3` / `4` / `5` — Add / Delete / Edit student
- `6` / `7` / `8` — Add / Delete / Edit group
- `0` — Exit
