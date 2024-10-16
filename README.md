# Time-Series Data Service

A Golang-based time-series data service that fetches data from an external API, stores it in time series DB, and exposes it via a gRPC API. The service support Prometheus for time-series data storage.

## Features

- Fetches time-series data from a provided API.
- Stores data in Prometheus.
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

```bash
docker-compose up --build
```

### Configuration
Configure your API endpoint and other settings in the config.yml file.

### Usage
The gRPC server will be available on the specified port (default: 5050).
Metrics can be accessed at /metrics on the specified port (default: 2112).

### API Endpoints
- gRPC: The service provides a gRPC interface for querying time-series data.
- Metrics: The service exposes metrics for monitoring at the /metrics endpoint.
