# README
## 1. 介绍
本库实现gin和kratos框架基于validator实现友好式的参数校验，优势在于接入方便、错误提示信息进行中文翻译，使其可读性更强。
## 2. 安装依赖
确保您已经安装了 Go，并且设置了 Go 模块（Go modules）
```shell
go get -u github.com/go-playground/validator/v10
go get -u github.com/go-playground/universal-translator
go get -u github.com/go-playground/validator/v10/translations/zh
```

## 3. 基于Gin WEB框架下的使用
### 3.1 直接在main.go中引入如下package
```go
import _ "github.com/yueja/go-validator"
```
### 3.2 在Controller层参数绑定中使用本工具ShouldBindGinValidator
```go
import  validator "github.com/yueja/go-validator"

   ...
   var register Register
   if err := validator.ShouldBindGinValidator(ctx, &register); err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }
   ...
```
### 3.3 更多test测试用例，见gin_test.go

## 4. 基于Kratos 微服务框架下的中间件使用
### 4.1 安装相关基础依赖
```shell
# kratos相关依赖
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/google/wire/cmd/wire@latest
# proto自定义tag依赖
go install github.com/favadi/protoc-go-inject-tag@latest
```
### 4.2 定义proto文件
```protobuf
message DemoRequest {
  // id必须大于 1
  int64 id = 1 ; // @gotags: json:"id" validate:"required" zhtrans:"用户ID"
  // age必须在 0 到 120 之间
  int32 age =2 ;// @gotags: json:"age" validate:"gte=0,lte=120" zhtrans:"年龄"
}
注：其中zhtrans是可选tag，主要用于参数校验错误提示友好性，进行对应字段中文翻译
```

### 4.3 编译pb.go文件
在Makefile文件api添加如下命令，需要注意此命令需要在protoc命令之后执行
```shell    
protoc-go-inject-tag -input="/you_pb.go_path/*.pb.go"
```

### 4.4 文件生成
```go
以下为proto文件通过插件生成的文件部分内容

type StudentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id必须大于 1
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"required" zhtrans:"用户ID"` // @gotags: json:"id" validate:"required" zhtrans:"用户ID"
	// age必须在 0 到 120 之间
	Age int32 `protobuf:"varint,2,opt,name=age,proto3" json:"age" validate:"gte=0,lte=120" zhtrans:"年龄"` // @gotags: json:"age" validate:"gte=0,lte=120" zhtrans:"年龄"
}
```

### 4.5 项目引用
#### 4.5.1 直接在main.go中引入如下package
```go
import _ "github.com/yueja/go-validator"
```
#### 4.5.2 接入中间件
```go
import  validator "github.com/yueja/go-validator"

    ...
    var opts = []http.ServerOption{
        http.Middleware(
             validator.ValidatorMiddlewareKratos(),
             ...
        ),
    }
    ...
```

### 4.6 测试用例
限于篇幅以及框架复杂原因，本工具尚未提供基于kratos中间件实现validator的测试用例

### 4.7 更多validator语法学习见以下网址
```shell
https://pkg.go.dev/github.com/go-playground/validator/v10
```
