## grpc-go目录

参考：[https://github.com/golang-standards/project-layout/blob/master/README_zh.md](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

 ```
 .
├── build # 构建相关
│   ├── build.sh # 构建指定服务的二进制程序
│   ├── ci # 流水线，可以多条
│   │   ├── xxx-pipeline1.yaml
│   │   └── xxx-pipeline2.yaml
│   ├── custom # 不同服务的自定义构建，可替换默认的方式
│   │   └── server3
│   │       └── Dockerfile
│   ├── Dockerfile # 构建镜像使用
│   └── assets # 打包到镜像里的文件
│       └── health_check.sh
├── cmd # 程序入口，不要在这个目录放置太多代码
│   ├── server1 # 每个服务对应的一个main.go
│   │   └── main.go
│   │   └── grpc.yaml # 如果grpc.yaml不分环境，将这个配置跟二进制程序一次打包
│   └── server2
│       └── main.go
├── deployments # （暂时不用）部署相关，每个环境都有一套完整的配置，原则上环境之间不共享配置
│   ├── development # 开发环境
│   ├── production # 生产环境
│   │   ├── cd.yaml # 部署编排模板
│   │   ├── config.yaml # 环境级别的业务配置
│   │   ├── dashboard.json # 环境级别的监控视图配置
│   │   ├── grpc.yaml # 环境级别的框架配置
│   │   ├── k8s.yaml # 环境级别的k8s yaml
│   │   └── servers
│   │       ├── server1 # 服务级配置
│   │       │   ├── clusters # 多集群部署的配置
│   │       │   │   └── guangzhou
│   │       │   │       └── k8s.yaml
│   │       │   ├── config.yaml # 服务级别的业务配置
│   │       │   └── grpc.yaml # 服务级别的框架配置
│   │       └── server2
│   ├── staging # 预发环境
│   └── testing # 测试环境
├── go.mod # 依赖管理
├── internal # 服务内部的业务逻辑
│   ├── pkg # 仓库内使用公共包
│   │   └── README.md
│   └── services # 业务接口逻辑入口
│       ├── server1 # 服务的接口逻辑
│       │   └── greet.go # service为greet的接口
│       └── server2
│           └── greet.go
├── pkg # 给仓库外部使用的公共包
│   └── README.md
├── README.md # 项目的说明文档
└── scripts # 公共脚本，如格式化等
    └── fmt.sh
 ```

 另外自动生成的`README.md`内容包含如何构建-部署-查看可观测数据等

