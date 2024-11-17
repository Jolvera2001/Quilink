# Diagram

```mermaid
erDiagram
    x[User]
    p[Profile]
    b[Blog]
    l[Links]

    x||--o{b : writes
    x||--o{l : "has many"
    x||--||p : "has one"

```