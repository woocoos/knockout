# Knockout

knockout是提供了woocoos服务平台的资源管理功能的核心库。

目录结构

```
├── api                 // API相关
│   ├── graphql         // graphql文档
│   │   ├── generated   // 生成的代码主要放置目录
│   │   └── model       // graphql生成模型
│   ├── proto           // protobuf文档
│   └── osa             // openapi3文档
├── cmd                 // 可执行程序目录,目录下每个子目录为一个可执行程序
├── codegen             // graphql代码目录
│   ├── entgen          // ent生成工具
│   │   ├── schema      // ent模型
│   │   ├── types       // ent模型中使用的自定义类型
│   ├── gqlgen          // gqlgen生成工具
│   └── oasgen          // openapi3文档
├── ent                 // Ent数据层生成代码
├── security            // 鉴权及身份验证相关
├── service             // 业务逻辑代码
├── status              // 错误代码相关
├── test                // 测试辅助类及数据
│   ├── testdata        // 测试数据初始化
│   ├── test.go         // 测试环境初始化
│   └── initdb.go       // 测试数据库初始化
├── version             // 应用版本信息
│   └── info.go         
```

# 快速安装

快速安装以Mysql为例，其他的可参考Knockout文档。

## 准备数据库


```shell
docker run --name knockout-db --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=portal -d mysql
```

```shell
make db
```