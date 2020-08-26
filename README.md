# grpc-go

## 1. 编写proto文件
    
## 2. 根据proto编译pb
```shell script
    protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

## 3. 运行server
```shell script
    go run greeter_server/main.go 
```

## 4.运行client
```shell script
    go run greeter_client/main.go
```
启动server之后运行client会打印出
```
    grpc test for go!
```

