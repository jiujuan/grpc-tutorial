## 运行测试
运行服务端main.go
```shell script
cd ./simple/server
go run main.go
```

运行客户端main.go
```shell script
cd ./simple/client

go run main.go
2019/07/17 12:31:34 receive response: age:50 name:"Jimmy two"
```

修改客户端 main.go 中的代码，把 ID 修改为 1
```shell script
req := pb.UserRequest{ID: 1}
```
继续测试客户端main.go
```shell script
go run main.go
2019/07/17 12:32:16 receive response: age:30 name:"Tom one"
```
