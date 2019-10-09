#####   go-kit微服务：HTTP REST

gokit工具集

```
go get github.com/go-kit/kit；
```

http请求路由组件

```
go get github.com/gorilla/mux；
```

#####   设置GOLANG环境变量

* GO111MODULE 最好设置到环境变量中


Powershell下:

* Enable the go modules feature


    $env:GO111MODULE="on"
       
* Set the GOPROXY environment variable


    $env:GOPROXY="https://goproxy.io"


#####   创建Service


    
#####   创建请求、响应模型

Protobuf功能是代码生成，可以免去手写一些模板化的代码。

1. 下载 https://github.com/protocolbuffers/protobuf/releases 相应版本

2. 解压后，复制 protoc 至 /usr/local/bin/下

3. 尝试 protoc --version 是否成功


    安装golang的protobuf代码生成器 protoc-gen-go
    go get -u github.com/golang/protobuf/protoc-gen-go
    安装micro的protobuf插件 protoc-gen-micro
    go get -u github.com/micro/protoc-gen-micro


#####  创建Endpoint


编写 user.proto

```
syntax = "proto3";

service User {
	rpc info(UserRequest) returns (UserResponse) {}
}

message UserRequest {
	string name = 1;
}

message UserResponse {
	string code = 1;
}

```

进入该文件所在目录生成代码

protoc --proto_path=. --micro_out=. --go_out=. ./services/user/user.proto

#####   创建Transport

#####   编写Main方法

#####   编译&运行

#####    TODO

*[ ] 参数不是合法ID时，没有明确提示错误 比如：{"id":"x"}
#####   更多参考

https://github.com/micro-in-cn/tutorials


https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro

https://micro.mu/docs/cn/features.html