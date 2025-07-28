# Simple Iot System

This is a backend application for managing a simplified IoT system of devices and measurements using Go and MongoDB. It is composed by two microservices, one for the collection of data and the other implements a trivial data aggregation.

## Prerequisites

- Docker
- Docker Compose

## Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/AndreaTrasacco/simple-iot-system.git
   ```
   ```bash
   cd simple-iot-system
   ```

2. **Run the Application**

Use Docker Compose to build and run the application:

```bash
docker-compose build
```
```bash
docker-compose up
```

This will start the Go services and MongoDB server.

## API Endpoints

### Register a New Device

```http
POST http://localhost:8888/devices
```

Sample Request Body:

```json
{
    "deviceId": "dev-123",
    "type": "thermometer",
    "location": "43.717992, 10.946594"
}
```
### List All Devices


```http
GET http://localhost:8888/devices
```

### Upload Measurements


```http
POST http://localhost:8888/measurements
```

Sample Request Body:


```json
[
    {
        "deviceId": "dev-123",  
        "timestamp": "2025-07-26T18:40:00Z",
        "metric": "temperature",
        "value": 25.6
    }
]
```

### Get Measurements for a Device


```http
GET http://localhost:8888/measurements?deviceId=dev-123&from=2025-07-26T18:33:29.03Z&to=2025-07-26T20:33:29.03Z
```

### Get Statistics


```http
GET http://localhost:8889/stats?deviceId=dev-123&metric=temperature
```

### Running Tests
Run unit tests using Go's testing package (Go required):


```bash
cd devicedatacollector
```
```bash
go test ./...
```

This will execute all tests in the project and display the results.

## Additional Notes

- Ensure Docker is running on your machine.
- The application services are configured to run on ports 8888 and 8889.
