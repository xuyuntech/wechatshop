pkg 包说明
====
* `api` 业务相关的 api 接口
* `config` 项目启动时的 yaml 配置文件类
* `manager` 操作类，比如 数据库、MQ 等等
* `middleware` 中间件，比如 auth 等等
* `models` 数据库表模型，以及 engine 创建实例，为了便于开发使用 `xorm` ORM 框架
* `modules` 一些工具类
  * `httplib` Request 工具