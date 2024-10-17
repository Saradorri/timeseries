# Time-Series Data Service

A Golang-based time-series data service that fetches data from an external API, stores it in time series DB, and exposes it via a gRPC API. The service support InfluxDB for time-series data storage.

## Features

- Fetches time-series data from a provided API.
- Stores data in InfluxDB.
- Updates data in dynamic interval (default: every 5 min) using a scheduler.
- Exposes a gRPC API for querying time-series data.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Saradorri/timeseries.git
cd timeseries
```

2. Build and run the services using Docker Compose:
(use .env.example file to create your own .env file)

```bash
docker-compose --env-file `path/to/your/.env` up --build
```

### Configuration
Configure your API endpoint and other settings in the `config.yml` file.

### Usage
- The gRPC server will be available on the specified port (default: **5050**).

### API Endpoints
- gRPC: The service provides a gRPC interface for querying time-series data.
