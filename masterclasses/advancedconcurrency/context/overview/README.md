# Package context Overview

## New process everywhere
- Inter service
  - Client to server
  - Server to server
  - Client to client (client to server to client)
- Within service
  - New goroutine/thread

## Package context
Concurrently safely communicating:

- Deadlines
- Cancellation signals
- Other request-scoped values

Across API boundaries and between processes

## Advantages
- Concurrently safe out of the box
  - Built-in concurrently safe components and functionalities
  - Use of other concurrently safe libraries & methodologies 
  - Limiting manipulations within itself & its children
- Standardization
  - Common interface
  - Common best practices