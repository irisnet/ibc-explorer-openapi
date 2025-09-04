# ibc-explorer-openapi
IBC Network Explorer Open API is a service for a blockchain explorer that provides fundamental data visualization for cross-chain ecosystems, focusing on IBC (Inter-Blockchain Communication) network information queries.

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.

## Getting Started

### Building from Source

First, build the application:

```bash
make build
```

### Running the Application

Start the application with default configuration:

```bash
./ibc-explorer-openapi start
```

Or start with a custom configuration file:

```bash
./ibc-explorer-openapi start test -c configFilePath
```


### Running with Docker

1. Build the Docker image:

```bash
docker build -t ibc-explorer-openapi .
```

2. Run the container:

```bash
docker run --name ibc-explorer-openapi -p 8000:8000 ibc-explorer-openapi
```

## Configuration

### Environment Variables
- CONFIG_FILE_PATH: `option` `string` Path to the configuration file
