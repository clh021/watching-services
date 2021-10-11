# cloud-server

这里是一个极简 cloud-server 模型，通过 wordpress 灵活的扩展功能，延展接口来达到快速测试的目的。

## 使用

```
make init # 初始化项目
make run  # 启动项目
```

## 接口工作流程

```
* 表示 已经完成实现
d 表示 不考虑 cloud 处实现，而是通过请求转发的方式由 deviceBox 响应请求，更贴近最终效果。

步骤  计划   接口描述        路由               参数
1.1. [*]客户端注册账户       users/register   username,password,email
2.1. [ ]设备登录            device/connect   username,password,deviceId
2.2. [d]设备传递在线数据     device/status     uptime,disk,cpu,mem,tasklist,applist
3.1. [*]客户端登陆          users/login       username,password
3.2. [ ]客户端获取设备列表    device/list
3.3. [d]客户端查看设备状态    device/info       deviceId
4.1. [ ]客户端获取应用列表    app/list
4.2. [d]客户端选择安装应用    app/install       deviceId,appId
4.3. [ ]设备应用配置下载     app/download      appId 设备请求app配置文件用
5.1. [d]客户端获取应用状态    app/info          deviceId,appId
5.2. [d]客户端选择卸载应用    app/uninstall     deviceId,appId
```