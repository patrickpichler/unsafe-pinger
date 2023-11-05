# Unsafe-Pinger

This application is a demonstration of command injection susceptibility designed for presentation
purposes at the `Vienna Devops Meetup`.

## Instructions

* Initiate the web server by executing `go run main.go`.
* Ping any website using `curl localhost:8080/ping/127.0.0.1`.
* Conduct a command injection by using `curl localhost:8080/ping/127.0.0.1';ls'`.
