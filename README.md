# Organization API

![Go](TODO:get-badge-in-the-future-for-this-repo)
[![Go Report Card](TODO:get-report-card)](TODO:probably-from-goreportcard)

`gun-organization` is a modern, professional, REST API service developed with Go (Golang) that provides a unified interface to extract and manipulate organizational data from GitHub. This service is specifically focused on returning lists of users, emails, and teams based on the given GitHub organization.

## 📚 Table of Contents

- [Organization API](#organization-api)
  - [📚 Table of Contents](#-table-of-contents)
  - [💡 Features](#-features)
  - [🔧 Requirements](#-requirements)
  - [🚀 Setup](#-setup)
  - [💻 Usage](#-usage)
  - [🧪 Testing](#-testing)
  - [🤝 Contributing](#-contributing)
  - [📜 License](#-license)

## 💡 Features

- Retrieves the organization data from GitHub via their API.
- Offers various methods that manipulate the retrieved data.
- Returns a list of users, their corresponding emails, and the teams they are part of.
- Utilizes Go's powerful features to offer efficient, reliable, and concurrent operations.

## 🔧 Requirements

- Go 1.16 or higher.
- A GitHub Personal Access Token.

## 🚀 Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/forjadev/gun-organization.git
   ```
2. Install the dependencies:
   ```sh
   go mod download
   ```
3. Create a `.env` file in the root directory of the project by copying the `.env.example` file:
   ```sh
   cp .env .env
   ```
4. Update the environment variables in the .env file with your own values.
5. Build the project:
   ```sh
   go build
   ```
6. Run the server:
   ```sh
   ./main
   ```
The server will start running on http://localhost:8080.

## 💻 Usage

TODO

## 🧪 Testing

TODO

## 🤝 Contributing

TODO

## 📜 License

This project is [MIT](LICENSE) licensed.
