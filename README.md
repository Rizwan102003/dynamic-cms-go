# Dynamic CMS in Golang

A **Dynamic Content Management System (CMS)** built with **Golang** and the **Gin framework**, designed to offer a flexible, schema-less approach to content management. Unlike traditional CMS solutions that rely on rigid database schemas, this system leverages **reflection** to handle dynamic content types and **metaprogramming** to generate API routes dynamically.

## 🚀 Features

- **Dynamic Content Types**: Define and manage content types (e.g., articles, blog posts, events) without a predefined database schema.
- **Metaprogramming for Routing**: Automatically generate RESTful API routes for CRUD operations based on the registered content types.
- **Flexible Storage Options**: Easily switch between SQLite, PostgreSQL, MongoDB, or an in-memory store.
- **Lightweight and Fast**: Built using Golang and the Gin framework for high performance.

## 📦 Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/dynamic-cms-go.git
   cd dynamic-cms-go
   
2. Install dependencies:

```sh
go mod tidy
```

3. Run the server:

```sh
go run main.go
