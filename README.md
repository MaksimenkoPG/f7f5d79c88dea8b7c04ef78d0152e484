# WTF?
Ruby + Go + RabbitMQ
- Ruby service, backend, calc a distance and route duration between two points(origin latitude/longitude and destination latitude/longitude).
- Go service, frontend, has a HTTP interface for end users for calc trip cost.
It gets origin latitude/longitude and destination latitude/longitude from user then gets the distance and route duration from Ruby service and then calcs trip cost.

# Setup
- [Install docker](https://docs.docker.com/install/#supported-platforms)
- Get [API Key](https://developers.google.com/maps/documentation/distance-matrix/get-api-key) for Google Distance Matrix API
- Clone this repo

# UP
- `RABBITMQ_DEFAULT_USER='user' RABBITMQ_DEFAULT_PASS='pass' GOOGLE_API_KEY='API_KEY' docker-compose up`

# Check
- Check `http://localhost:8080/trip_info` with params:
  - o_latitude
  - o_longitude
  - d_latitude
  - d_longitude

For example: `http://localhost:8080/trip_info?o_latitude=55.737547&o_longitude=37.408085&d_latitude=55.753564&d_longitude=37.621085`
