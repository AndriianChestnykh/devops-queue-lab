package queue

// TODO: add github.com/redis/go-redis/v9 when implementing this adapter.
// Planned operations: XADD, XREADGROUP, XACK, XAUTOCLAIM.
// Store a result before acknowledging; replayed jobs must be idempotent.
