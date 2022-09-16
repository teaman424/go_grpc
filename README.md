###嘗試練習grpc實作
以protobuf 產生新的 .proto.go 檔
```
 protoc *.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
```



參考
https://grpc.io/docs/languages/go/quickstart/
https://pjchender.dev/golang/grpc-getting-started/