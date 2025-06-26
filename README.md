# 🧱 Blog API – Modular Architecture in Go

This project demonstrates a modular architecture structure in Go, focused on decoupling business logic from entry interfaces (REST and GraphQL).

The main goal is to show how we can structure scalable and maintainable APIs, enabling framework switching or adding new interfaces without impacting the core logic.

---

## ✨ Technologies used

- [Go](https://golang.org)
- [Fiber](https://gofiber.io/) — REST
- [gqlgen](https://gqlgen.com/) — GraphQL
- Project structure using `internal/`, `cmd/`, `pkg/`

---

## 📁 Project structure

```

blog-api/
├── cmd/                # Entrypoints (REST and GraphQL)
│   ├── rest/
│   └── graphql/
├── internal/
│   ├── domain/         # Entities and interfaces
│   ├── service/        # Business logic
│   └── database/       # Repositories (in-memory)
├── graph/              # GraphQL schema and resolvers
├── pkg/                # Utilities (logger)

````

---

## 🚀 Running the REST API

```bash
cd cmd/rest
go run main.go
````

Access: `http://localhost:3000`

### Available endpoints:

* `GET /articles`
* `GET /articles/:id`
* `POST /articles`
* `PUT /articles/:id`
* `DELETE /articles/:id`

---

## 🚀 Running the GraphQL API

```bash
cd cmd/graphql
go run main.go
```

Access the Playground at: `http://localhost:8080`

### Example queries:

```graphql
query {
  posts {
    id
    title
    content
    author
  }
}
```

```graphql
mutation {
  createArticle(input: {
    title: "Example",
    content: "Post content"
  }) {
    id
  }
}
```

---

## 🧪 Architecture goals

* Separate business logic from external interfaces
* Reuse core logic across REST, GraphQL, gRPC, etc.
* Simplify maintenance and testing
* Enable flexible growth and framework swapping
