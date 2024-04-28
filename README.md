# clean-json

`clean-json` is a Go utility designed to extract the largest valid JSON object from a stream of mixed content. It is particularly useful for parsing logs, data streams, or files where JSON data is embedded within other text.

## Features

- **Robust JSON Extraction**: Efficiently parses and extracts embedded JSON from mixed text inputs.
- **Handles Large Inputs**: Optimized to process large volumes of data without performance degradation.
- **Flexible**: Can handle various JSON formats including objects and arrays.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

```bash
go version go1.15.2 linux/amd64
```

### Installing

A step-by-step series of examples that tell you how to get a development environment running:

Clone the repository:

```bash
git clone https://github.com/garyblankenship/clean-json.git
```

Navigate to the cloned directory:

```bash
cd clean-json
```

Build the project:

```bash
go build
```

Run the utility:

```bash
echo '{"example": "data"}' | ./clean-json
```

## Running the tests

Explain how to run the automated tests for this system:

```bash
go test
```

## Usage

Here's a quick example to show you how to use `clean-json`:

```bash
echo 'Some random text {"json": true} and more text' | ./clean-json
```

Expected output:

```json
{"json": true}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

