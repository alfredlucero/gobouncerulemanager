# gobouncerulemanager

My Go MySQL REST API version of a bounce rule manager

## Getting Started

Install MySQL locally and start it up like so while in the project root directory

```bash
mysql -u <user> -p

# Type in password and hit enter to enter MySQL CLI

# Load up initial database data
source db/db.sql

```

OR start up the Docker container with MySQL running like so

```bash
# TODO: fill these out once we have the Docker container
```

Make sure to export these environment variables so the server knows how to connect to the local MySQL database.

```bash
export MYSQL_USER=username
export MYSQL_PASSWORD=password
export MYSQL_DATABASE=database
```

To start up the server you may go to `cmd/server/main.go` and do

```bash
go run main.go
```
