# golang
This contains two projects, Web and Worker.

## Environment Variables
```shell
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export $AWS_REGION=
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

## Demo
```shell
$ docker compose up -d
$ curl -X POST http://localhost:8001/ -H 'Content-Type: application/json' -d '{"domains": ["www.youtube.com", "www.example.com"]}'
{"jobId":"156f4300-e122-4776-aea1-bbbc27c0bef5"}
$ curl http://localhost:8001/156f4300-e122-4776-aea1-bbbc27c0bef5
{"jobId":"156f4300-e122-4776-aea1-bbbc27c0bef5","completed":true,"progress":100,"data":[{"domain":"www.youtube.com","a":"142.251.220.206\n142.251.220.238\n142.251.221.14\n142.251.221.46\n142.251.220.142\n142.251.220.174\n2404:6800:4017:801::200e\n2404:6800:4017:802::200e\n2404:6800:4017:803::200e\n2404:6800:4017:804::200e","cname":"youtube-ui.l.google.com"},{"domain":"www.example.com","a":"93.184.216.34\n2606:2800:220:1:248:1893:25c8:1946","cname":""}]}
```
