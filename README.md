# Go Functional Layered Architecture

A sample Go project implementing layered architecture with functional programming paradigms. This project demonstrates clean architecture using dependency injection and functional approaches through a reservation system example.

This project follows the [Go Module Layout](https://go.dev/doc/modules/layout#server-project) guidelines for server projects, keeping all Go packages implementing the server's logic in the `internal` directory.

## Project Overview

This project provides an implementation example of functional layered architecture in Go. Rather than traditional object-oriented approaches, the design emphasizes treating functions as first-class objects, focusing on pure functions and function composition.

- **Functional Approach**: Uses function types instead of interfaces across layers
- **Pure Dependency Injection**: Dependencies passed as function arguments
- **Layer Separation**: Clear separation of responsibilities between Controller, Service, and Infrastructure layers
- **Comprehensive Testing**: Integration tests for behavior verification

## Architecture

```
┌─────────────────┐
│   Controller    │ ← HTTP request/response handling
├─────────────────┤
│    Service      │ ← Business logic implementation
├─────────────────┤
│     Infra       │ ← Data access layer (in-memory implementation)
└─────────────────┘
```
