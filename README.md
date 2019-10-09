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

#####   编写Main方法

#####   编译&运行

#####    TODO

*[ ] 参数不是合法ID时，没有明确提示错误 比如：{"id":"x"}

#####   更多参考

https://github.com/metaverse/truss

https://github.com/go-kit/kit

https://juejin.im/post/5c6eaff8e51d4572c9581069