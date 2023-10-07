# golang
This contains two projects, Web and Worker.

## Environment Variables
```shell
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_DEFAULT_REGION=
export WORKER_DNS_LOOKUP_TYPE=A
export RMQ_URL=amqp://guest:guest@localhost:5672/
```

## Web
For serving API using Gin Gonic.
```shell
$ go build -o ginamus main.go
$ ./ginamus web
```

# Worker
For handling DNS resolution jobs.
```shell
$ go build -o ginamus main.go
$ ./ginamus worker
```
