# go-uni Prototype

## Overview
The `go-uni` prototype is a script designed to monitor newly created tokens on the Ethereum blockchain using the Uniswapv2 contract. It leverages the Go programming language and interacts with the Ethereum network through the Ethereum client and WebSocket client. The script is in its prototype stage, demonstrating basic functionality for tracking token creation events.

## Features
- **Token Monitoring**: The script continuously monitors new block headers for PairCreated events emitted by the Uniswapv2 contract.
- **Concurrency**: It utilizes goroutines to process PairCreated events concurrently, enhancing performance.
- **Graceful Shutdown**: Incorporates graceful shutdown handling to ensure clean termination upon receiving termination signals like SIGINT or SIGTERM.
- **Error Handling**: Includes error handling mechanisms for robustness during Ethereum client connection, contract instantiation, event filtering, and processing.

## Installation
- Clone the repository containing the script.
- Ensure that Go is installed on your system.
- Install the required dependencies using `go mod tidy`.

## Usage
- Configure the Ethereum client and WebSocket client endpoints as per your requirements.
- Compile and execute the script.
- The script will continuously monitor new block headers and print information about newly created tokens to the console.

## License
This project is licensed under the MIT License.

## Future Improvements
- Enhance error handling and logging for better debugging.
- Implement more robust event processing mechanisms.
- Expand functionality to include additional event types and data.

## Disclaimer
This script is provided as-is and is currently in a prototype stage. Use it at your own risk, and ensure appropriate testing and validation before deploying in a production environment.


## Acknowledgments
- The creators and maintainers of the Ethereum client libraries used in this project.
- The developers of Uniswap for their contributions to decentralized finance.
