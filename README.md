# ant-frame
## 平台简介
Ant frame 一个基于go语言的前后端分离的脚手架,相比于spring boot和laravel个人觉得go开发后端接口更简洁、明了.
项目使用于中小型项目,前期版本主要用于功能的完善,并未考虑高并发场景.项目启动于2023年10月,希望大家多提宝贵的意见和建议.
## 目录结构
```
filetree 
|--bootstrap       启动目录   
    |--config      配置目录   
    |--global      全局启动器   
    |--initialize  初始化模块   
    |-- cmd        命令行目录   
|-- internal       内部目录   
    |-- handler    控制器目录   
    |-- middleware 中间件目录   
    |-- router     路由目录   
    |-- model      模型目录  
        |-- entity 数据模型   
        |-- res    返回模型   
        |-- rep    请求模型   
    |-- service    服务目录   
    |-- dao        数据操作目录  
|-- pkg            开放目录  
    |-- tools      公用工具类  
    |-- encode     统一编码   
    |-- sdk        第三方sdk   
|-- api            api接口目录   
    |-- v1         version版本   
|-- storage        存储目录
    |-- logs       日志目录
    |-- assets     静态文件目录
main.go
```
## 项目的开发包
* gin 路由中间件框架
* viper 配置项管理
* zap&lumberjack 日志记录
* gorm 数据库操作
* casbin 权限认证
* cobra 命令行工具
* gin-jwt jwt工具