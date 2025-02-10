
### 目录结构
``` shell
gomall_demo
├─ app # 业务开发代码目录
    ├─ auth # 认证服务
    │   ├─ biz # 业务部分代码
    │   ├─ conf # 配置文件相关
    │   └─ infra # rpc初始化相关
    ├─ cart # 购物车服务
    ├─ checkout # 结算服务
    ├─ email # 邮件服务
    ├─ frontend # 前后端不分离项目；前端代码，api请求，hanler处理，service业务
    ├─ order # 订单服务
    ├─ payment # 订单服务
    ├─ product # 产品服务    
    └─ user # 用户服务
├─ common # 自定义工具包
├─ db # 初始化数据库
├─ deploy #
├─ rpc_gen # kitex的代码生成目录
└─ idl # idl目录

```

### 开发流程
1. 编写 idl
2. 生成 idl
   cwgo client --type RPC --service user --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/user.proto	
3. 完善各类配置（如 pkg/global/rpc.go ）
   每一个服务下对应的 /conf 包含对应需要修改的配置文件信息 例如 /conf/test/conf.yaml中的mysql地址等
   注意需要完善对应配置地址和 .env文件才能正确启动consul和进行服务注册和发现
4. 在 app 目录下生成对应服务
5. 在 fontend 目录下生成服务对应的前端处理逻辑 
   cwgo server -I ../../idl --type HTTP --service frontend --module  gomall_demo/app/frontend --idl ../../idl/frontend/home_page.proto
6. 利用 docker-compose 启动 docker 中的 mysql consul 等容器: `sudo docker-compose up`
7. 服务完善后，依次启动微服务


### 环境准备

#### docker-compose
[dokcer-compose 安装](https://cloud.tencent.com/developer/article/2204414)

使用 docker 配置 consul  :
```shell
sudo docker-compose up # 项目根目录终端运行
```

### 微服务交互步骤
1. handleer.go 实现我们的方法
2. service   文件夹中对应的函数处理业务
3. main.go 启动微服务
4. consul 生成服务发现客户端
5. 另一个服务在自己的 main.go init函数中完成服务发现，并且调用客户端或者方法。