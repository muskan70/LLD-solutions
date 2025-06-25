## Batch Processing

### Requirements:
1. Create 3 go routines where first one is pushing data to a channel
2. Second one consuming it from channel, such that given a timeout and fileSize, if any of timeout or fileSize has reached , flush that batch data to third goroutine to print it.
3. Implement graceful shutdown.

- “Shutdown works by first closing all open listeners, then closing all idle connections, and then waiting indefinitely for connections to return to idle and then shut down.”