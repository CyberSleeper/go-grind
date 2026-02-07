## Problem 5: Dependency Injection (Testability)

### 📌 Concept

Global variables or hard-coded database connections inside functions make code impossible to unit test.
To fix this, we use **Interfaces**. Instead of depending on a _concrete_ struct (like `PostgresConnection`), our function should depend on an _interface_ (like `Datastore`).

This allows us to inject a **Real Database** in production, but a **Mock Database** (fake) during testing.

### 📝 Task

Refactor the `GetUser` function and the `Service` struct.

1.  Define an interface `Datastore` that has a method `Query(id string) string`.
2.  Refactor `Service` to hold this interface, not the concrete struct.
3.  **In Main:**
    - Create a struct `MockDB` that implements `Datastore` but just returns "Fake User" (no real DB logic).
    - Inject this `MockDB` into the `Service`.
    - Call `GetUser` and print the result.

### 🚫 Starter Code

```go
package main

import "fmt"

// Imagine this is a heavy SQL driver
type RealPostgresDB struct{}

func (db RealPostgresDB) Query(id string) string {
    return "Real User from DB: " + id
}

type Service struct {
    // BAD: Hard dependency on the concrete struct
    db RealPostgresDB
}

func (s Service) GetUser(id string) string {
    return s.db.Query(id)
}

func main() {
    db := RealPostgresDB{}
    svc := Service{db: db}

    // This works, but we can't test it without a running DB!
    fmt.Println(svc.GetUser("123"))
}
```
