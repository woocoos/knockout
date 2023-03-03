# Knockout

knockout是提供了woocoos服务平台的基础定义及管理功能的核心库.

目录结构

```
├── api             // API相关
│   ├── graphql     // graphql文档
│   ├── knockoutpb  // protobuf文档
│   └── openapi     // openapi3文档
├── graph           // graphql代码目录
│   ├── entgen      // ent生成工具
│   │   ├── schema  // ent模型
│   │   ├── types   // ent模型中使用的自定义类型
│   ├── generated   // 生成的代码主要放置目录
│   ├── gqlgen      // gqlgen生成工具
│   └── model       // graphql生成模型
├── service         // 服务代码
└── test            // 测试辅助类及数据
```