<div align="center">
  <h1 style="font-size: 4rem; font-weight: bold;">UCheck</h1>
  <p>A simple and fast command-line tool for checking the status of multiple URLs concurrently.</p>
</div>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Badge">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License Badge">
</p>

---

## 🚀 About UCheck

UCheck is a lightweight and efficient CLI tool built with Go that allows you to add multiple URLs to a queue and check their status codes and latency in parallel. It provides a simple and intuitive interface for managing your URL queue and executing checks.

## ✨ Features

-   🌐 **Add URLs:** Easily add URLs to a persistent queue.
-   📋 **List URLs:** View all the URLs currently in your queue.
-   ⚡️ **Concurrent Checks:** Utilizes Go's concurrency to check multiple URLs at once, providing fast results.
-   📊 **Detailed Report:** Get a clear summary of the status (success or failure), status code, and latency for each URL.
-   🗑️ **Auto-Clean:** The queue is automatically cleared after a successful run.
-   🧹 **Manual Clean:** Option to manually clear the entire queue.

## 🛠️ Tech Stack

| Category      | Technology                                                                                           |
| :------------ | :--------------------------------------------------------------------------------------------------- |
| **CLI**       | [Go](https://go.dev/)                                                                                |
| **Core Libs** | `net/http`, `sync`, `encoding/json`                                                                  |

## 📂 Project Structure

```
ucheck/
├── cmd/                # Command-line interface logic
│   ├── add.go
│   ├── clean.go
│   ├── execute.go
│   └── list.go
├── internal/           # Core application logic
│   ├── client.go
│   └── utils.go
├── go.mod              # Go module definition
├── main.go             # Application entry point
└── README.md           # This file
```

## 🏁 Getting Started

### Prerequisites

-   [Go](https://go.dev/doc/install) (version 1.22 or higher)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/tristnaja/Ucheck.git
    cd Ucheck
    ```

2.  **Build the binary:**
    ```sh
    go build
    ```
    This will create a `ucheck` (or `ucheck.exe`) executable in the project directory.

3.  **(Optional) Install globally:**
    You can move the built binary to a directory in your system's `PATH` to make it accessible from anywhere.
    ```sh
    # For Linux/macOS
    sudo mv ucheck /usr/local/bin/

    # For Windows (using PowerShell)
    Move-Item -Path .
    ```

## 🕹️ Usage

UCheck uses a simple command structure. The URL queue is stored in `~/.config/ucheck/db.json`.

### 1. Add a URL

Add a new URL to the queue. You can add multiple URLs by running this command several times.

```sh
ucheck add -u <your-url>
# Or using the shorthand
ucheck add --url <your-url>
```
**Example:**
```sh
ucheck add -u https://www.google.com
ucheck add -u "https://github.com"
```

### 2. List URLs

View all the URLs currently in the queue.

```sh
ucheck list
```
**Example Output:**
```
Here are the list of your queue:
| Size | 2
1. https://www.google.com
2. https://github.com
```

### 3. Run the Checks

Execute the status check for all URLs in the queue. The queue will be cleared after the run.

```sh
ucheck run
```
**Example Output:**
```
Worker 1: finished processing https://www.google.com with Job ID of 0
Worker 2: finished processing https://github.com with Job ID of 1

-+-+-+-Final Report-+-+-+-
Job ID: 0 | URL: https://www.google.com | (SUCCESS) Status Code: 200
Job ID: 1 | URL: https://github.com | (SUCCESS) Status Code: 200
```

### 4. Clean the Queue

Manually remove all URLs from the queue.

```sh
ucheck clean
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have ideas for improvements or find any bugs.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for details.
