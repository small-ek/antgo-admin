#### 1. 使用说明

```
- golang版本 >= v1.16
- 本项目是一个基于golang的web项目，使用了gin框架、gorm框架、jwt框架、swaggo框架、antgo框架等，实现了一个简单的web项目,如果有问题请联系作者邮箱：56494565@qq.com
- IDE推荐：Goland
```

#### 1.1 接口文档生成

```
- 接口文档生成工具：swaggo/swag
- 安装：go install github.com/swaggo/swag/cmd/swag@latest
- 生成接口文档：swag init --parseDependency --parseInternal
- 访问接口文档：http://localhost:9001/swagger/index.html
```

#### 1.2 ant-cli 使用说明

```
- 框架Cli工具：ant-cli
- 安装：go install github.com/small-ek/ant-cli@latest
- 帮助文档：ant-cli h
- RSA生成：ant-cli rsa
- 创建项目:ant-cli create demo
- 生成接口相关内容,请先设置配置文件,生成后需要执行:ant-cli install
- 命令根据表生成接口：ant-cli gen api database.table
- 命令根据表生成模型：ant-cli gen model database.table
- 命令根据表生成Dao：ant-cli gen dao database.table
- 前端接口生成,可以生成关联模型相关：ant-cli ui
- CLI编译：ant-cli build main.exe
- CLI运行：ant-cli run main.go
- 包项目安装：ant-cli install
- 访问地址：http://127.0.0.1:49000
```

#### 1.3 go mod 使用说明

```
- 依赖管理工具：go mod
- 初始化：go mod init
- 整理和更新项目的依赖关系：go mod tidy
- 依赖的模块复制到项目：go mod vendor
```

#### 1.3 go run 使用说明

```
- 命令用于编译并运行Go程序工具：go run
- 初始化：go run main.go
- 运行项目：go run main.go
```

#### 1.3 go build 使用说明

```
- 命令用于编译Go程序工具：go build
- 编译：go build -o main.exe -ldflags "-s -w"
- 运行：./main.exe
- linux交叉编译：
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o main-linux -ldflags "-s -w"
- mac交叉编译：
go env -w GOOS=darwin
go env -w GOARCH=amd64
go build -o main-macos -ldflags "-s -w"
- windows交叉编译：
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o main.exe -ldflags "-s -w"
```