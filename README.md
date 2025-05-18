# Ambient Weather Prometheus

This project exposes a webhook endpoint for ambient weather stations to send weather data and exports this data using a Prometheus client.

## Project Structure

```
ambient-weather-prometheus
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── webhook.go   # Handles incoming webhook requests
│   ├── metrics
│   │   └── prometheus.go # Prometheus metrics setup and exposure
│   └── models
│       └── weather_data.go # Defines the structure of weather data
├── go.mod                # Module definition file
├── go.sum                # Module dependency checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd ambient-weather-prometheus
   ```

2. Install the necessary dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

- The application will start an HTTP server that listens for incoming POST requests on the `/webhook` endpoint.
- Ambient weather stations can send weather data in JSON format to this endpoint.
- The application will process the incoming data and update Prometheus metrics accordingly.

## Prometheus Metrics

- The metrics can be accessed at the `/metrics` endpoint.
- Ensure that Prometheus is configured to scrape this endpoint to collect the weather data metrics.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.