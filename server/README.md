# G站: https://www.gourouting.com

欢迎来到[w站](www.wjkj.com)，本站是[Singo](https://github.com/bydmm/singo)框架的学习项目。

## 项目地址



## 项目目的

测试 lotus 的客户端实现 IPFS

## 重要: 如何运行

#### 1.学习Go Module管理依赖

本项目已经迁移到使用Go Module来管理依赖,和视频的开始有所不同! 所以按照视频的方法是跑不起来的。

请参考本视频了解什么是Go Module：https://www.bilibili.com/video/av63052644/

Go Module会让你未来面对各种依赖问题迎刃而解，所以学习和使用对你是非常有价值的

#### 2.配置数据库

本项目依赖于任何网站项目都会使用的Mysql和Redis，所以你需要提前安装和启动这两个服务。

如果你是windows用户，可以快速的解决mysql和redis安装的问题,通过: PHPStudy。

本视频用几分钟教会你如何使用PHPStudy，https://www.bilibili.com/video/av64485001/

如果你是OSX或者linux的硬核用户，相必启动Mysql和Redis对你不是问题😁

#### 3.配置环境变量

> 设置环境变量，你可以参考singo框架的文档: https://singo.gourouting.com/quick-guide/set-env.html

由于每个用户的电脑环境不同，所以我们通过环境变量来改变着些容易变化的属性。

你需要复制项目根目录下的.env.example文件，然后建立.env文件，然后把内容帖进去

```ini
MYSQL_DSN="user:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local" # mysql连接串
REDIS_ADDR="127.0.0.1:6379" # redis地址
REDIS_PW="" # redis密码(可以不填)
REDIS_DB="" # redis数据库(可以不填)
SESSION_SECRET="youneedtoset" # session密钥，开发环境可以不用改
GIN_MODE="debug" # 服务状态，开发环境不用改
# 下面是OSS对象存储的参数
# 参考本视频来管理上传文件：https://www.bilibili.com/video/av60189734/
OSS_END_POINT="oss-cn-hongkong.aliyuncs.com" # OSS端点
OSS_ACCESS_KEY_ID="xxx"
OSS_ACCESS_KEY_SECRET="qqqq"
OSS_BUCKET="lalalal"

```

#### linux / OSX 系统启动

```bash
export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GO111MODULE=on

go run main.go
```

## 神奇的接口文档

服务启动后: http://localhost:3000/swagger/index.html

接口文档位于项目swagger目录下。请阅读目录内的文档