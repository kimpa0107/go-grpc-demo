# Golang gRPC demo

### install dependency module

```bash
go mod tidy
```

### start server

```bash
go run server/main.go
```

### start client

```bash
go run client/main.go
```

#### response

```bash
2022/04/09 13:56:03 Response: [id:"article-1" title:"Article 1" category:{id:1 name:"Go"} tags:"Go" tags:"GoLang" id:"article-2" title:"Article 2" category:{id:2 name:"gRPC"} tags:"grpc"]
```