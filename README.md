# Time-Series Data Service

A Golang-based time-series data service that fetches data from an external API, stores it in a time-series database (`InfluxDB`), and exposes it via a `gRPC API`. This service transitioned from using Prometheus in version v1.0.0.

## Features
- Fetches time-series data from a provided API.
- Stores data in `InfluxDB`.
- Updates data at a dynamic interval (default: every 5 minutes) using a scheduler.
- Exposes a gRPC API for querying time-series data.

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Installation
Clone the repository:
```bash
git clone https://github.com/Saradorri/timeseries.git
cd timeseries
```
Create your own `.env` file using the provided `.env.example` file and configure the necessary environment variables. Build and run the services using Docker Compose:
```bash
docker-compose --env-file path/to/your/.env up --build
```
### Configuration
Configure your API endpoint and other settings in the `config.yml` file.
### Usage
The gRPC server will be available on the specified port (default: `5050`).
### API Endpoints
**gRPC**: The service provides a gRPC interface for querying time-series data. You can use any gRPC client to connect and make requests.
### Sample gRPC Request and Response
You can use Postman or a gRPC client to make requests to the gRPC API. Here's an example:
```json
{
    "aggregation": "AVG",
    "end": "1729105300",
    "start": "1719105200",
    "window": "5m"
}
```

### Sample gRPC Response
```json
{
    "meta": {
    "window": "5m",
    "aggregation": "avg",
    "status": "SUCCESS",
    "message": "Query executed successfully"
    },
    "data": [
    {
            "time": 1697731200000,
            "value": 12.5
        },
        {
            "time": 1697731500000,
            "value": 13.0
        },
        {
            "time": 1697731800000,
            "value": 14.2
        }
    ]
}

```