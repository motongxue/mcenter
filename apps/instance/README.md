# 服务实例

所有的服务都需要通过接口注册到注册中心

+ 第三方服务推荐使用HTTP接口注册
+ 内部服务使用GRPC接口注册



## 服务注册流程
```
instance --注册请求--> 注册中心
         <--实例配置--              
```


+ heartbeat
    + interval: 心跳上报的间隔, 单位秒


## 注册中心高可用

注册中心启动时 会把自己注册给网关, 客户端通过对外暴露出来的网关访问注册中心

```
                       |--实例注册服务
instance --> treafik---|--实例注册服务
                |      |--实例注册服务
                |            ｜
              etcd <--注册----|
```