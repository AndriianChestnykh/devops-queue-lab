package main

import "example.com/queue-lab/internal/server"

// TODO: start the consumer loop alongside the health server.
func main() { server.Run("worker") }
