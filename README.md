# neo-chat

Natural language query interface for NASA Near-Earth Object data, powered by ClickHouse and Claude.

Ask questions in plain English like "Which asteroids came closest to Earth in the last 10 years?" and get answers backed by real NASA data.

## Data Sources

All data comes from NASA's [SSD/CNEOS API](https://ssd-api.jpl.nasa.gov/):

| Dataset | Description | Rows |
|---------|-------------|------|
| [Close Approaches](https://ssd-api.jpl.nasa.gov/doc/cad.html) | Predicted asteroid/comet flybys of Earth (1900–2100) | ~210K |
| [Small-Body Database](https://ssd-api.jpl.nasa.gov/doc/sbdb_query.html) | Orbital & physical data for all known asteroids and comets | ~1.5M |
| [Sentry](https://ssd-api.jpl.nasa.gov/doc/sentry.html) | Objects with non-zero Earth impact probability | ~2K |

## Prerequisites

- [Go 1.24+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/) (for ClickHouse)

## Usage

```bash
# Build
go build -o neo-chat .

# Fetch NASA data into CSV files
./neo-chat fetch

# Load CSVs into ClickHouse
./neo-chat load

# Start the API server
./neo-chat serve
