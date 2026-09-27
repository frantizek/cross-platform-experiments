# Cross-Platform Experiments

A single repository for experimenting with **Node.js**, **Go**, and **Rust**.
Each project demonstrates the **strengths** of its respective language.

---

## **Structure**

```
cross-platform-experiments/
│
├── nodejs/          # Node.js: Recursive file system traversal (async/await)
│   ├── src/
│   │   └── app.js   # Recursively reads files in a directory
│   └── package.json
│
├── go/              # Go: Concurrent recursive web scraper (goroutines)
│   ├── src/
│   │   └── main.go  # Recursively scrapes links from a webpage
│   └── go.mod
│
├── rust/            # Rust: Recursive file search (error handling, safety)
│   ├── src/
│   │   └── main.rs  # Recursively searches for a file by name
│   └── Cargo.toml
│
└── README.md        # Repository documentation
```

---

### **Node.js**

1. Navigate to the `nodejs` folder:
   ```bash
   cd nodejs
   ```
2. Install dependencies
   ```bash
   npm install
   ```
3. Run the script
   ```bash
   node src/app.js
   ```

### **Go**

1. Navigate to the `go` folder:
   ```bash
   cd go
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the program:
    ```bash
    go run src/main.go


### **Rust**

1. Navigate to the `rust` folder:
   ```bash
   cd rust
   ```

2. Run the program:
   ```bash
   cargo run
   ```





### **Key Features Demonstrated**

| Language | Example | Strengths Demonstrated |
|---------|---------|---------|
| Node.js | Recursive file system traversal | Async/await, non-blocking I/O, file system ops |
| Go | Concurrent web scraper | Goroutines, concurrency, simplicity |
| Rust | Recursive file search | Safety, error handling, performance |