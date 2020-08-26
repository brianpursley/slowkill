# slowterm

A simple docker container that just waits for SIGINT or SIGTERM, sleeps for a specified amount of time, and then exits with a specified exit code.

Default amount of time to sleep is 75 seconds.
Default exit code is 1.

# Usage Examples

Wait for SIGINT or SIGTERM, sleep for 75 seconds and exit with 1 (Default Settings)
```
docker run brianpursley/slowterm
```

Wait for SIGINT or SIGTERM, sleep for 5 seconds and exit with 2
```
docker run -e SLEEP=5 -e EXIT=2 brianpursley/slowterm
```