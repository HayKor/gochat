# GoChat 

Simple chat application built on go.

## Getting Started

### Dependencies

- `Go` (built on 1.23)

### Installing

```shell
git clone https://github.com/HayKor/gochat.git
```

### Executing program

Run the server locally

```shell 
go run cmd/server/main.go
```

After that, start the client from another terminal

```shell
go run cmd/client/main.go
```

Or alternatively just connect from any TCP client.

```shell
nc localhost 3000
ncat localhost 3000
telnet localhost 3000
# and so on
```

**Enjoy**!!

## Authors

Creator - [HayKor](https://github.com/HayKor)

