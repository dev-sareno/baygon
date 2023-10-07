FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o ginamus .

FROM golang:1.21-alpine
WORKDIR /app
COPY --from=build /app/ginamus .
RUN chmod +x /app/ginamus
ENTRYPOINT ["/app/ginamus"]
