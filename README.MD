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
  - [x] **Day 2: Routing**
      - [x] Concepts: `http.ServeMux`, Pattern matching
      - [x] Implementation: Handling dynamic paths (e.g., `/users/{id}`)
  - [x] **Day 3: Middleware**
      - [x] Concepts: Decorator pattern, Higher-Order Functions
      - [x] Implementation: Logging middleware, Panic recovery
  - [x] **Day 4: Data Interchange (JSON)**
      - [x] Concepts: `encoding/json`, Struct tags, Content-Type headers
      - [x] Implementation: JSON Request/Response helpers
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

  * [Go 1.22+](https://go.dev/dl/) installed on your machine (required for wildcard routing features).
  * `curl` or a browser for testing endpoints.

### Quick Start

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/goserve.git
    cd goserve
    ```

2.  **Run the server:**

    ```bash
    go run .
    ```
    
    *Or specify all files:* `go run main.go handler.go middleware.go helper.go model.go`

3.  **Test the endpoints:**
    Open a new terminal and run:

    ```bash
    # Check server health
    $ curl http://localhost:8081/health
    {"status":"UP"}

    # List endpoints
    $ curl http://localhost:8081/users
    [{"id":1,"name":"Alice","email":"alice@example.com"},{"id":2,"name":"Bob","email":"bob@example.com"}]

    $ curl http://localhost:8081/posts
    [{"id":1,"title":"First Post","content":"Hello world","author":"Alice"},{"id":2,"title":"Go is great","content":"Concurrency is fun","author":"Bob"}]

    # Create a user (POST)
    $ curl -X POST http://localhost:8081/users -d '{"name": "John", "email": "john@example.com"}'
    {"id":101,"name":"John","email":"john@example.com"}

    # Dynamic routing with IDs
    $ curl http://localhost:8081/users/123
    {"id":123,"name":"User 123","email":"user123@example.com"}

    $ curl http://localhost:8081/posts/456
    {"id":456,"title":"Post 456","content":"Some content here...","author":"Unknown"}

    # Error handling - invalid ID
    $ curl http://localhost:8081/users/abc
    {"error":"invalid id"}
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

`ServeMux` is Go's HTTP request multiplexer (router). It matches incoming requests to handlers based on URL patterns.

**Key insights:**
  * **DefaultServeMux vs Custom Mux:** Using `http.HandleFunc("/path", handler)` registers routes on the **global** `DefaultServeMux`. This is risky because any imported package can also register routes on it.
  * **Production best practice:** Always create a custom `mux := http.NewServeMux()` and pass it to `ListenAndServe(":8080", mux)` to avoid global state pollution.
  * **Pattern matching rules:**
    - Exact match: `"/health"` only matches `/health`
    - Prefix match: `"/users/"` matches `/users/`, `/users/123`, `/users/123/profile`, etc.
    - Root pattern: `"/"` is a special case - it acts as a catch-all for any path that doesn't have a more specific handler
  * **Longest match wins:** If multiple patterns could match, ServeMux chooses the most specific (longest) one. E.g., `/api/v1/status` takes precedence over `/api/`

### **2.1. Go 1.22+ Enhanced Routing**

Go 1.22 introduced **wildcard patterns** in ServeMux, making dynamic routing much cleaner without third-party libraries.

**Key insights:**
  * **Wildcard syntax:** `/users/{id}` captures the dynamic segment as a named parameter
  * **`PathValue()` method:** `r.PathValue("id")` extracts the captured segment directly - no manual string parsing needed!
  * **Before Go 1.22:** You had to use prefix patterns (`/users/`) and manually parse with `strings.TrimPrefix()` or `strings.Split()`
  * **Pattern specificity:** `/users/{id}` is more specific than `/users/`, so it takes precedence
  * **Wildcard limitations:** Wildcards match any non-empty segment but not slashes (e.g., `{id}` matches "123" but not "123/profile")

**Example:**
```go
// Go 1.22+ way
mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")  // Clean extraction!
    fmt.Fprintf(w, "User ID: %s", id)
})

// Pre-1.22 way
mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/users/")  // Manual parsing
    fmt.Fprintf(w, "User ID: %s", id)
})
```

### **3. Middleware Chaining**

Middleware in Go implements the **decorator pattern** - wrapping handlers to add cross-cutting concerns without modifying the core business logic.

**Key insights:**
  * **Higher-Order Functions:** Middleware is a function that takes a handler and returns a new handler with added behavior
  * **Function signature:** `func(http.HandlerFunc) http.HandlerFunc` - takes a handler, returns a wrapped handler
  * **Closures:** The returned function "closes over" the original handler, maintaining access to it
  * **Execution flow (onion model):** Request flows through layers in order, then response flows back in reverse
  * **Order matters:** Outer middleware wraps inner middleware - panic recovery should be outermost to catch all errors

**Example pattern:**
```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Code BEFORE handler (on request)
        start := time.Now()
        
        next(w, r)  // Call the wrapped handler
        
        // Code AFTER handler (on response)
        duration := time.Since(start)
        log.Printf("Request: %s %s - %v", r.Method, r.URL.Path, duration)
    }
}
```

### **3.1. Panic Recovery with `defer` and `recover()`**

Go's panic recovery mechanism prevents a single handler crash from bringing down the entire server.

**Key insights:**
  * **`defer` keyword:** Deferred functions execute after the surrounding function returns (or panics)
  * **`recover()` built-in:** Only works inside `defer` - catches panics and returns the panic value
  * **Returns `nil` if no panic:** Always check `if err := recover(); err != nil` before acting
  * **LIFO execution:** Multiple defers execute in reverse order (last-in, first-out)
  * **Production safety:** Panic recovery is essential - prevents cascading failures from affecting other requests

**Example pattern:**
```go
func panicRecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("PANIC: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next(w, r)
    }
}
```

### **3.2. Middleware Composition**

Chaining multiple middlewares requires understanding execution order and proper composition.

**Key insights:**
  * **Nesting order:** `panicRecovery(logging(handler))` means panic recovery wraps logging, which wraps handler
  * **Execution flow:** Request → Panic Recovery → Logging → Handler → Logging → Panic Recovery → Response
  * **Composition helper:** Create a `WithDefaults()` function to apply standard middleware stack consistently
  * **Best practice:** Panic recovery should be outermost layer to catch errors from all inner layers
  * **Reusability:** Define middleware once, apply to all routes for consistent behavior

**Example pattern:**
```go
func WithDefaults(handler http.HandlerFunc) http.HandlerFunc {
    return panicRecoveryMiddleware(loggingMiddleware(handler))
}

// Usage
mux.HandleFunc("/health", WithDefaults(healthHandler))
```

### **4. JSON Data Interchange**

Handling JSON in Go requires understanding the `encoding/json` package and how Go's type system interacts with JSON data.

**Key insights:**
  * **Marshaling vs Encoding:**
    - `json.Marshal(v)`: Returns `[]byte`. Great for small objects or when you need the byte slice.
    - `json.NewEncoder(w).Encode(v)`: Writes directly to an `io.Writer` (stream). More memory efficient for HTTP responses but if header has sent and the encoding process fails halfway,
    the status code can't be updated anymore.
  * **Unmarshaling vs Decoding:**
    - `json.Unmarshal(data, &v)`: Parses `[]byte` into a struct.
    - `json.NewDecoder(r.Body).Decode(&v)`: Reads directly from `io.Reader` (stream). Best for HTTP request bodies.
  * **Struct Tags:** You control JSON output using tags like `json:"user_id,omitempty"`. Without tags, fields use their Go struct names.
  * **Exported Fields:** Only Capitalized fields are visible to the JSON encoder. Lowercase fields are ignored!
  * **Strict Decoding:** `decoder.DisallowUnknownFields()` is crucial for APIs to prevent clients from sending unexpected data silently.

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

### Day 2: Routing ✅

  * **Progress:** Implemented custom ServeMux with static and dynamic routing using Go 1.22+ wildcard patterns.
  * **Endpoints Implemented:** 
      * Static: `/users`, `/posts`
      * Dynamic: `/users/{id}`, `/posts/{id}`
  * **Key Learnings:**
      * **Custom ServeMux vs DefaultServeMux:** Created isolated router with `http.NewServeMux()` to avoid global state pollution. Production-safe pattern that prevents imported packages from registering unexpected routes.
      * **Go 1.22's wildcard routing:** Native support for `/path/{param}` syntax eliminates need for third-party routers. The `r.PathValue("param")` method extracts captured segments cleanly.
      * **Pattern matching precedence:** ServeMux uses "longest match wins" - more specific patterns take priority (e.g., `/users/{id}` beats `/users/`).
      * **DRY principles in action:** Created reusable `extractID()` helper function to avoid duplicating validation logic across handlers.
      * **Validation layers:** Implemented both empty-check and numeric validation for IDs. Returns appropriate HTTP status codes (400 Bad Request for invalid input).
      * **Type safety matters:** Helper returns `int` instead of `string` for IDs, catching non-numeric input early with `strconv.Atoi()`.
  * **Challenges Solved:**
      * Understanding when to use exact match (`/users`) vs prefix match (`/users/`) patterns.
      * Extracting and validating dynamic path segments without regex.
      * Building reusable validation logic that works across multiple endpoints.
      * Deciding between manual string parsing vs Go 1.22's native wildcards.
  * **Code Quality Improvements:**
      * Extracted common logic into helper functions
      * Consistent error handling across all dynamic routes
      * Proper HTTP status codes for different error scenarios

### Day 3: Middleware ✅

  * **Progress:** Implemented production-ready middleware system with logging and panic recovery. Refactored codebase into clean, modular structure.
  * **Middleware Implemented:**
      * Logging middleware (request method, path, duration)
      * Panic recovery middleware (catches crashes, prevents server shutdown)
      * `WithDefaults()` composition helper (applies standard middleware stack)
  * **Key Learnings:**
      * **Decorator pattern in action:** Middleware wraps handlers like onion layers, each adding specific behavior without modifying core logic.
      * **Higher-Order Functions:** Functions that take functions and return functions - the signature `func(http.HandlerFunc) http.HandlerFunc` is the key pattern.
      * **Closures are powerful:** The returned function maintains access to the original handler through closure scope.
      * **`defer` and `recover()` mechanism:** Deferred functions run after panics, and `recover()` catches them - but only works inside `defer`.
      * **Middleware order is critical:** Panic recovery must be outermost layer to catch errors from all inner layers. The pattern is `panicRecovery(logging(handler))`.
  * **Challenges Solved:**
      * Understanding function closure mechanics - how the returned function "remembers" the `next` handler.
      * Getting middleware execution order right - visualizing the onion model helped (request flows in, response flows out).
      * Realizing `recover()` only works inside `defer` - tried calling it directly at first.
      * Deciding on middleware composition strategy - created `WithDefaults()` for DRY principles.
  * **Code Quality Improvements:**
      * Refactored handlers into separate file for better organization
      * Created reusable middleware composition function
      * All routes now protected by panic recovery
      * Consistent logging across all endpoints
      * Enhanced panic logs with request context (method, path)

### Day 4: Data Interchange (JSON) ✅

  * **Progress:** Implemented full JSON support for requests and responses. Created reusable helpers and standardized error handling.
  * **New Features:**
      * `model.go` with struct definitions and JSON tags
      * `writeJSON` helper for consistent response formatting
      * `readJSON` helper for safe request parsing
      * Standardized error response format (`{"error": "message"}`)
      * Strict HTTP method matching (GET vs POST)
  * **Key Learnings:**
      * **Streaming JSON:** Using `json.NewEncoder` and `json.NewDecoder` is more efficient for HTTP than `Marshal/Unmarshal` because it works with streams (`io.Writer`/`io.Reader`).
      * **Content Negotiation:** Always setting `Content-Type: application/json` tells the client what to expect.
      * **Struct Tags:** Controlling how Go fields map to JSON keys (e.g., `Name` -> `"name"`).
      * **Strict Decoding:** `DisallowUnknownFields()` prevents clients from sending garbage data.
      * **Method Specificity:** Go 1.22 allows specifying methods in routes (`POST /users`), which simplifies handler logic significantly.
</details>

-----

## 📂 Project Structure

```text
.
├── main.go           # Entry point: Server configuration and route registration
├── handler.go        # HTTP handlers: Business logic for each endpoint
├── middleware.go     # Middleware: Logging, panic recovery, composition helpers
├── helper.go         # Utility functions: JSON helpers, validation, etc.
├── model.go          # Data models: Struct definitions
├── go.mod            # Go module definition
└── README.md         # Documentation
```
