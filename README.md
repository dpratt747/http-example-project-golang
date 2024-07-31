### To get the go version installed on your system:

```bash
go version
```

### To init a GoLang project:

-   This creates a `go.mod` file

```bash
go mod init http-example-project
```

---

### Example main method in main.go

To run:

```bash
go run .
```

### To list possible go commands:

```bash
go help
```

---

### HTTP Server:

```bash
cd http-server
go run .
curl http://localhost:8080
```

Curl with body:

```bash
curl --location --request GET 'http://localhost:8080/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "David",
    "age": 29
}'
```

---

### HTTP Client:

```bash
cd http-client
go run .
```

---

### Related links:

-   [Digital Ocean HTTP server](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go#prerequisites)
