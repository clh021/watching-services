# home-cloud-server

### 介绍
HomeCloud项目中，提供应用配置文件即可启动应用服务提供统一访问接口的服务端

### 项目结构
项目结构说明
```
├── bin                         # 测试编译的临时目录
│   └── server
├── conf
│   ├── conf.go                 # 程序启动自身所依赖的配置
│   └── service.go              # 被扫描服务 的配置
├── controller                  # 控制中心
│   ├── controller.go           # 控制器 发现服务后的调度者
│   ├── generator               # 生成器 主要用于生成 nginx 配置
│   │   └── generator.go
│   └── operator                # 操作者 主要用于操作 docker 容器
│       └── opera.go
├── example.conf.yaml           # 程序使用所依赖的配置 示例
├── example.service.yaml        # 扫描服务时 服务自身的配置 service 示例
├── go.mod
├── go.sum
├── main.go                     # 程序入口
├── Makefile
├── models                      # 保存一些公共的结构体
│   ├── server.go
│   └── trigger.go
├── README.md
├── scaner.go                   # 扫描器，用于发现服务
├── tests                       # 测试目录， make serve 会直接使用此处测试环境
│   ├── conf.yaml
│   ├── nginx.conf
│   ├── services
│   ├── test2.yaml
│   ├── test.sh                 # 测试脚本，针对启动好的服务，操作目录，以验证服务是否正确进行
│   └── test.yaml
└── util.go
```

### 使用说明


```bash
make build # 编译
make buildcross # 交叉编译 mips64le, arm64, amd64 版本
make run # 运行编译好的程序
make serve # 编译并运行

# 项目中会监控指定目录的配置文件增改删的操作，可以在运行服务时调用脚本进行测试
# ./tests/test.sh
```
