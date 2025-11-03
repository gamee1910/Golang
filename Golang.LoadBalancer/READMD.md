# Golang Load Balancer Implementation

A simple load balancer implementation in Go using the Round Robin algorithm.

## Overview

This project demonstrates a basic HTTP load balancer that distributes incoming requests across multiple backend servers using the Round Robin scheduling algorithm.

## Features

- Round Robin load balancing algorithm
- Health checks for backend servers
- Simple HTTP server implementation
- Configurable backend servers
- Lightweight and efficient

## Usage

1. Run the load balancer:

```bash
go run main.go --backends=https://grafana.com:443,http://info.cern.ch:80
```

## How It Works

The load balancer:
1. Receives incoming HTTP requests
2. Selects the next available backend server using Round Robin
3. Forwards the request to the selected server
4. Returns the response to the client

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

MIT License