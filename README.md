A Shop for TongRenTang E-Business
================

前端部分
---

为了开发方便，项目构建使用 `create-react-app` 工具包，组件库使用 `antd-mobile` , 这两个工具文档都比较详细, `create-react-app` 可以了解一下，主要需要看一下 `antd-mobile` 组件的使用文档，文档链接如下：
* `antd-mobile` https://mobile.ant.design/docs/react/introduce-cn
* `create-react-app` https://github.com/facebookincubator/create-react-app/blob/master/packages/react-scripts/template/README.md

自行安装 `nodejs`

然后根目录下执行 `npm install` 安装依赖包

命令对应如下：

* 启动项目 , 启动之后就可以开发了
```
$ npm start
```
* 生产构建，会将资源文件生成到 `build` 目录下面
```
$ npm run build
```

后端部分
---
cmd/main.go 是入口

具体模块说明请看 pkg 目录的 readme
