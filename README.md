# Iranian National Code Generator (NCG)

A simple CLI tool written in Go to generate valid Iranian national codes. This tool is built using the [Cobra](https://github.com/spf13/cobra) library and provides commands to generate random national codes, including rounded random codes.

## Features
- Generate a random valid Iranian national code.
- Generate a rounded random valid Iranian national code.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/national-code-generator.git

Navigate to the project directory:

cd national-code-generator

Build the binary:

go build -o ncg

Run the tool:

    ./ncg

## Commands
# ncg generate

Generates a random valid Iranian national code.

# Usage:

./ncg generate

# Example Output:

Generated National Code: 1234567890

# ncg generate round

Generates a rounded random valid Iranian national code (e.g., numbers ending with 0).

# Usage:

./ncg generate round

# Example Output:

Generated Rounded National Code: 1234567800

# Iranian National Code Generator (NCG)

A simple CLI tool written in Go to generate valid Iranian national codes. This tool is built using the [Cobra](https://github.com/spf13/cobra) library and provides commands to generate random national codes, including rounded random codes.

## Features
- Generate a random valid Iranian national code.
- Generate a rounded random valid Iranian national code.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/national-code-generator.git

    Navigate to the project directory:

cd national-code-generator

Build the binary:

go build -o ncg

Run the tool:

    ./ncg

Commands
ncg generate

Generates a random valid Iranian national code.

Usage:

./ncg generate

Example Output:

Generated National Code: 1234567890

ncg generate round

Generates a rounded random valid Iranian national code (e.g., numbers ending with 0).

Usage:

./ncg generate round

Example Output:

Generated Rounded National Code: 1234567800

Example

$ ./ncg generate
Generated National Code: 4851027394

$ ./ncg generate round
Generated Rounded National Code: 8543200000

# How It Works

    The CLI uses the Cobra library to define commands and subcommands.
    National codes are generated with valid checksums based on the Iranian national code algorithm.
    Rounded codes are derived to end with predictable digits (e.g., 0).

# Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the tool.

# License

This project is licensed under the MIT License.
