# GoServe 🌐

**GoServe** is a lightweight, custom implementation of an HTTP server built to master backend engineering in Go.

This project was built to explore **Network Programming** concepts and master the **Go Standard Library** (`net/http`, `encoding/json`, `context`). Instead of using heavy frameworks (like Gin or Fiber), GoServe relies entirely on the standard library to understand how Go handles concurrency, routing, and the HTTP lifecycle under the hood.

-----

## 📅 Project Roadmap (7-Day Sprint)

I am building this project over one week, adding new layers of complexity—from basic socket listening to middleware and JSON handling—daily.

  - [x] **Day 1: The Connection**
      - [x] Basic Server Setup (`ListenAndServe`)
      - [x] Handlers: `http.HandlerFunc`, `http.ResponseWriter`
      - [x] Endpoint: `/health`
  - [ ] **Day 2: Routing**
      - [ ] Concepts: `http.ServeMux`, Pattern matching
      - [ ] Implementation: Handling dynamic paths (e.g., `/users/{id}`)
  - [ ] **Day 3: Middleware**
      - [ ] Concepts: Decorator pattern, Higher-Order Functions
      - [ ] Implementation: Logging middleware, Panic recovery
  - [ ] **Day 4: Data Interchange (JSON)**
      - [ ] Concepts: `encoding/json`, Struct tags, Content-Type headers
      - [ ] Implementation: JSON Request/Response helpers
  - [ ] **Day 5: RESTful Methods**
      - [ ] Concepts: GET vs POST vs PUT vs DELETE, HTTP Status Codes
      - [ ] Implementation: CRUD operations on an in-memory store
  - [ ] **Day 6: Context & Concurrency**
      - [ ] Concepts: `context.Context`, Timeout management, Cancellation
      - [ ] Implementation: Simulating slow processes, request timeouts
  - [ ] **Day 7: Testing & Graceful Shutdown**
      - [ ] Concepts: `net/http/httptest`, `os.Signal`
      - [ ] Implementation: Unit testing handlers, Safe server shutdown

-----

## 🛠️ Installation & Usage

### Prerequisites

  * [Go 1.20+](https://go.dev/dl/) installed on your machine.
  * `curl` or a browser for testing endpoints.

### Quick Start

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/goserve.git
    cd goserve
    ```

2.  **Run the server:**

    ```bash
    go run main.go
    ```

3.  **Test the endpoints:**
    Open a new terminal and run:

    ```bash
    # Check server health
    $ curl -i http://localhost:8080/health
    HTTP/1.1 200 OK
    Date: Mon, 01 Jan 2024 10:00:00 GMT
    Content-Length: 13
    Content-Type: text/plain; charset=utf-8

    Server is UP!
    ```

-----

## 🧠 Technical Learnings

*Documenting the specific Go concepts and HTTP protocol challenges mastered.*

### **1. The `Handler` Interface**

I learned that in Go, everything is an interface. To handle a request, you don't necessarily need a function; you need a type that implements `ServeHTTP(w http.ResponseWriter, r *http.Request)`.

  * **Correction:** Initially, I thought `http.HandleFunc` was the only way. I learned that `http.Handle` is more powerful as it allows passing structs that hold dependencies (like database connections).

### **1.1. `ResponseWriter` & `io.Writer`**

A critical realization: `http.ResponseWriter` is an interface that implements `io.Writer` because it has a `Write([]byte) (int, error)` method. This is why functions like `fmt.Fprintf(w, ...)` work - they accept any `io.Writer`.

**Key insights:**
  * `io.Writer` is one of Go's most fundamental interfaces - any type with a `Write()` method satisfies it
  * This design allows the same writing functions to work with HTTP responses, files, network connections, and buffers
  * The first argument to `fmt.Fprintf` determines the destination (client response vs console vs file)
  * Interface composition is Go's superpower - small, focused interfaces that compose together

### **2. Multiplexers (`ServeMux`)**

*Upcoming: Notes on how the DefaultServeMux works and why defining a custom Mux is better for production safety to avoid exposing global state.*

### **3. Middleware Chaining**

*Upcoming: Notes on wrapping handlers to intercept requests before they reach the core logic (onion architecture).*

-----

## 📝 Dev Log

<details>
<summary><strong>Click to expand Daily Logs</strong></summary>

### Day 1: The Connection ✅

  * **Progress:** Successfully launched a server on port 8080 and served plain text responses.
  * **Endpoints Implemented:** `/`, `/health`
  * **Key Learnings:**
      * **`http.ResponseWriter` is an interface**, not a struct. It implements `io.Writer`, which means it has a `Write([]byte) (int, error)` method.
      * **`io.Writer` interface pattern**: Because `ResponseWriter` implements `io.Writer`, any function accepting `io.Writer` works with it (like `fmt.Fprintf`). This is the power of Go's interface composition.
      * **Why `fmt.Fprintf(w, ...)`?** It's cleaner than `w.Write([]byte(...))` because it handles string formatting and byte conversion automatically. The first argument determines WHERE it writes (client response vs console).
      * **Header order matters**: Must call `w.Header().Set()` *before* writing the body, or Go auto-sends headers with 200 OK and guessed Content-Type.
      * **Explicit headers are best practice**: Always set `Content-Type` explicitly rather than relying on Go's auto-detection.
  * **Challenges Solved:**
      * Understanding why `log.Fatal(http.ListenAndServe(...))` is the standard idiom (it blocks forever unless an error occurs).
      * Distinguishing between `http.HandleFunc` (convenience wrapper) and `http.Handle`.
      * Realizing that the "/" pattern in `ServeMux` acts as a catch-all for any unknown route.
      * Understanding that `fmt.Fprintf(w, ...)` writes to the HTTP response (client sees it), not the console.

### Day 2: Routing

  * *Pending...*

### Day 3: Middleware

  * *Pending...*

</details>

-----

## 📂 Project Structure

```text
.
├── main.go           # Entry point: Server configuration and startup
├── go.mod            # Go module definition
├── handlers.go       # (Future) HTTP handler logic
├── middleware.go     # (Future) Logging and Auth layers
├── README.md         # Documentation
└── ...               # (Future) Internal packages
```
