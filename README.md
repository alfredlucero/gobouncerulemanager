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
export SERVER_PORT=8000

# OR you can change the values in local.conf and do
source local.conf
```

To start up the server you may go to `cmd/server/main.go` and do

```bash
go run main.go
```

### Using Prometheus and Grafana

We have a /metrics endpoint that gives us back some Prometheus metrics that we can visualize with Grafana.

Assuming you have Docker Desktop for Mac installed and running, you can start up the bounce rule server and the Grafana server like this.

```bash
# In one terminal start up the web server
go run cmd/bounceruleserver/main.go

# In another terminal instance, start up Grafana
docker run -d -p 3000:3000 grafana/grafana
```

Log into the Grafana server with admin/admin and connect to the Prometheus data source on http://localhost:8000.

### Sample CURLs for Bounce Rule Manager

`Getting all bounce rules`

```bash
curl -X GET localhost:8000/bounce_rules
```

`Getting a specific bounce rule`

```bash
curl -X GET localhost:8000/bounce_rules/1
```

`Creating a bounce rule`

```bash
curl -d '{ "response_code": 450, "enhanced_code": "4.7.1", "regex": "someregex", "priority": 0, "description": "some description", "bounce_action": "no_action"}' -H 'Content-Type: application/json' localhost:8000/bounce_rules
```

`Updating a bounce rule`

```bash
curl -X PUT -d '{ "response_code": 451, "enhanced_code": "4.8.1", "regex": "somenewregex", "priority": 0, "description": "some description", "bounce_action": "no_action"}' -H 'Content-Type: application/json' localhost:8000/bounce_rules/5
```

`Deleting a bounce rule`

```bash
curl -X DELETE localhost:8000/bounce_rules/5
```

`Getting all bounce rule changes`

```bash
curl -X GET localhost:8000/bounce_rule_changes
```

`Getting a bounce rule's changes`

```bash
curl -X GET localhost:8000/bounce_rule_changes/3
```
