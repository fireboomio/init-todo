
本仓库是 fireboom 的模板仓库，用于快速初始化 fireboom 项目。

> 飞布是下一代 API 开发平台，灵活开放、多语言兼容、简单易学，对标 Firebase，但无供应商锁定。 它帮助你构建生产级 WEB API，但无需花时间重复 coding。

> 产品愿景：**极致开发体验，`飞速布署`应用！**

- fireboom 主仓库：[前往->](https://github.com/fireboomio/fireboom)
- fireboom 官网：[前往->](https://www.fireboom.io/)

# 模板简介

`init-todo` 模板是 fireboom 的基础模板，只涉及 [OPERATION](https://docs.fireboom.io/he-xin-gai-nian/graphql#graphql-operation) ，不涉及钩子的编写。

> [!NOTE]  
> 后续分别制作针对 go 和 nodejs 开发者的模板～

点击下述按钮，立即线上体验该模板：

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/new/#https://github.com/fireboomio/init-todo)

本模板的核心是在不引入钩子的情况下，尽可能用 OPERATION 表达各种**增删改查**请求。

![operation 概念图](https://2723694181-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FNx22Cp3wzkuW1siRbMwW%2Fuploads%2Fgit-blob-9629f35549ab8028ae8b12686010fab95b2f1f48%2Fimage%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1).png?alt=media)

# 使用模板

## 初始化项目

使用如下命令，一键安装 fireboom ，并以当前仓库作为模板初始化项目：

```sh
curl -fsSL fireboom.io/install | bash -s fb-project -t init-todo
```

其，核心格式为：

```sh
curl -fsSL fireboom.io/install | bash -s fb-project -t [template_name]
```

其中 `[template_name]` 为对应模板的仓库名称，例如本仓库的 `init-todo`。


## 运行 fireboom

直接执行下面的命令行，以开发模式启动：

```shell
./fireboom dev
```

启动成功日志：

```sh
Web server started on http://localhost:9123
```

打开控制面板

[http://localhost:9123](http://localhost:9123)

## 更新 fireboom

此外，如果要更新 `fireboom` 版本，使用如下命令行：

```shell
curl -fsSL https://www.fireboom.io/update | bash
```

如果要使用预览版本，用如下命令行：
```shell
curl -fsSL https://www.fireboom.io/update-test | bash
```

# 模板详情

## 数据源

### sqlite 数据源

为了降低启动成本，本模板引入了最少的外部依赖，因此用 `sqlite` 数据库作为演示。

> [!NOTE]  
> 但额外的代价是，sqlite 不支持一些高级特性，例如数组类型、批量操作（见报错部分，如 n01curd/create/createMany、n02relation/write/createWithPost2）等。

**todo**数据库

一个待做事项数据库，只包含 `Todo` 表，其 schema 如下：

```prisma
datasource db {
  provider = "sqlite"
  # url中的绝对路径，因系统不同而不同
  url      = "file:/workspaces/init-todo/upload/sqlite/todo.db"
}

model Todo {
  id        Int      @id @default(autoincrement())
  title     String
  completed Boolean  @default(false)
  createdAt DateTime @default(now())
  authorId  Int?
}
```

**blog**数据库

一个简化版博客系统数据库，其 schema 如下：

```prisma
datasource db {
  provider = "sqlite"
  # url中的绝对路径，因系统不同而不同
  url      = "file:/workspaces/init-todo/upload/sqlite/blog.db"
}
/// 文章分类
model Category {
  id   Int    @id @default(autoincrement())
  name String @unique
  Post Post[]
}
/// 文章表
model Post {
  id        Int        @id @default(autoincrement())
  title     String
  published Boolean
  view      Int?
  like      Int?
  authorId  Int
  content   String?
  User      User       @relation(fields: [authorId], references: [id])
  Category  Category[]
}
/// 用户详细信息
model Profile {
  id      Int    @id @default(autoincrement())
  address String
  userId  Int    @unique
  User    User   @relation(fields: [userId], references: [id])
}
/// 用户表
model User {
  id      Int      @id @default(autoincrement())
  email   String   @unique
  age     Int?
  country String?
  Post    Post[]
  Profile Profile?
}
```

###  REST api 数据源

系统内置数据源，**system** REST

在fireboom 启动时默认生成，包含了 fireboom 的所有 REST 接口。

其 swagger 定义，详情见文件：`upload/oas/system.json`

## 数据库接口

fireboom 的数据库操作，基于 [prisma](https://docs.fireboom.io/he-xin-gai-nian/chao-tu#prisma) 实现，因此支持 prisma 引擎的所有特性。

参考 prisma 官方文档，梳理了 5 个目录的 operation。

- n01curd：[简单的增删改查](https://www.prisma.io/docs/orm/prisma-client/queries/crud#filter-by-multiple-field-values)
- n02relation：[关联表的增删改查](https://www.prisma.io/docs/orm/prisma-client/queries/relation-queries)
- n03filterAndSorting：[过滤和排序](https://www.prisma.io/docs/orm/prisma-client/queries/filtering-and-sorting)
- n04pagination：[分页](https://www.prisma.io/docs/orm/prisma-client/queries/pagination)
- n05aggregation：[分组、聚合](https://www.prisma.io/docs/orm/prisma-client/queries/aggregation-grouping-summarizing)

> [!NOTE]  
> 上述示例，未提及OPERATION 编译为 REST API后，如何传递入参。标量入参比较简单，无需额外介绍。若入参为对象或数组时，则需要掌握对应技巧，详情参见 [文档](https://docs.fireboom.io/ji-chu-ke-shi-hua-kai-fa/api-gou-jian/ke-shi-hua-gou-jian#query-operation)。

## 高级特性

`n06advanced` 目录包含了很多高级特性，可分为：

### 原生sql：

示例位于目录：`store/operation/n06advanced/raw`

- queryRaw：查询类 sql，返回查询的数组对象
- executeRaw：变更类 sql，如创建、更新、删除，返回影响行数

### 指令相关：

用指令修饰的 OPERATION，可实现复杂逻辑：

> [!NOTE]  
> 指令修饰的 OPERATION ，只有编译为 API 后才生效，因此无法在 OPERATION 编辑页测试，请在 API 预览页验证其最终效果。（跨域关联除外）

**跨域关联**

示例位于目录：`store/operation/n06advanced/join`

默认情况下，OPRATION 都是**并行**执行，但很多场景需要先查询数据，然后将返回值作为条件，执行下一个查询/变更。

因此，通过 字段： `_join`、`_join_mutation`和指令：`@internal`、`@export` 实现了该特性。

- 查询关联：将上一个返回值作为参数传给下一个**查询**
- 变更关联：将上一个返回值作为参数传给下一个**变更**

**事务控制**

示例位于目录：`store/operation/n06advanced/transaction`

默认情况下，MUTATION OPRATION 都是**并行**执行，但很多场景需要保证原子性。

-  @transaction 指令

> [!NOTE]  
> 除了用  @transaction 指令，还可以在钩子中执行事务，待完善文档。

**参数注入**

示例位于目录：`store/operation/n06advanced/variable_pro`

默认情况下，OPERATION 参数只能从请求 url path 或 body 中读取，有些场景需要在服务端设置。

- **@fromHeader**：读取客户端的请求头注入
- **@injectCurrentDateTime**：服务器生成时间注入
- **@injectEnvironmentVariable**：读取环境变量注入
- **@injectGeneratedUUID**：服务器生成 uuid 注入

**入参校验**

示例位于目录：`store/operation/n06advanced/variable_pro/jsonschema`

为保证安全性，OPERATION 参数需要校验，指令 `@jsonSchema` 用来实现该功能。

- **正则校验**：正则表达式
- **通用校验**：如内置邮箱校验
- **数字校验**：数字范围校验
- **文本校验**：文本数量校验

**响应控制**

示例位于目录：`store/operation/n06advanced/selection_pro`

默认情况下，OPERATION 编译的 REST API，其响应值为对应数据源中的字段命名。某些场景下，要根据前端需求返回对应的别名或结构。

- **@formatDateTime**：格式化 `DateTime` 类型的字段
- **@transform**：“拍平”深层嵌套，支持数组和对象

### 设置相关：

示例位于目录：`store/operation/n06advanced/setting`

除了 OPERATION 描述 API 外，还可以额外设置特性。

- **实时查询**：服务端定时轮询数据源，数据变化后，推送给前端，实现数据的准实时更新
- **速率限制**：限制任意接口的 请求频次，实现流控

## OIDC 相关

fireboom 还支持身份认证，基于 [OIDC 协议](https://docs.fireboom.io/ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/shou-quan-ma-mo-shi)实现了灵活的身份验证机制。

### 设置接口授权

示例位于文件：`store/operation/n07oidc/auth`

**API 设置步骤**

1. 选中接口，查看右侧面板，切换到 “设置” 
2. 开启“使用独立配置”，并开启“接口授权”
3. 再次访问该接口需要登录

### 用户参数注入

示例位于目录：`store/operation/n07oidc/fromclaim`

- **查询权限控制**：只查询登录用户拥有的数据
- **更新权限控制**：只更新登录用户拥有的数据
- **创建权限控制**：新建数据的所有人设置为登录用户

> [!NOTE]  
> OIDC 规范中只声明里用户相关的数据，如 UID、EMAIL。若想注入业务数据，可编写钩子，详情见[文档](https://docs.fireboom.io/jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi#mutatingpostauth)或[代码片段](https://github.com/fireboomio/amis-admin/blob/c56897a2b9def58a00416b9b5df135f0259726e3/backend/custom-go/authentication/mutatingPostAuthentication.go#L21)。

### 测试登录

系统内置了`身份验证`供应商 `auth0`（https://auth0.com/）

**如果用 `localhost:9123` 访问控制台**

1. 打开 api 预览页：http://`localhost:9123`/#/workbench/rapi
2. 在右上角 <未登录> 中点击 <auth0> 进行登录
3. 输入账户/密码：test@example.com/1QAZ2wsx
4. 登录后可查看当前登录用户的信息
5. 最后，在预览页测试 `@fromclaim` 修饰的 OPERATION

**如果用 自定义域名，如`http://test.fireboom.io:9123` 访问控制台，需要额外配置**

1. 前往 `auth0` 后台，设置 <Application URIs>-><Allowed Callback URLs> 为 `身份验证`供应商的「登录回调地址」
2. 前往 firebom 控制台 [设置]->[安全]-><重定向URL>，新增 url：`http://test.fireboom.io:9123`/#/workbench/rapi/loginBack

详情参考 [文档](https://docs.fireboom.io/ji-chu-ke-shi-hua-kai-fa/shen-fen-yan-zheng/shou-quan-ma-mo-shi#oidc-pei-zhi)。


## 钩子相关

至此，该模板的核心功能已经介绍完毕。后面内容可跳过。

该模板的目标是尽可能多的描述只用 OPERATION 就能完成的场景，但仍有些场景需要编写代码才能实现。这就需要用到 fireboom 的钩子机制。

![Fireboom 钩子架构图](https://2723694181-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FNx22Cp3wzkuW1siRbMwW%2Fuploads%2Fgit-blob-24c89a58be58a1feadda5631d0781b74ef2b6dc7%2Fimage%20(2)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1)%20(1).png?alt=media)

> [!NOTE]  
> 目前已支持 **NodeJS**、**Golang**、**Java** 的SDK，其他未提供SDK 的语言，可基于 [HTTP 规范](https://docs.fireboom.io/jin-jie-gou-zi-ji-zhi/operation-gou-zi)自行开发。

钩子使用，详情参考 [文档](https://docs.fireboom.io/jin-jie-gou-zi-ji-zhi/gou-zi-ji-zhi)。

### 内部调用  @internalOperation

示例位于文件：`store/operation/n10hook/internal`

默认情况下，OPERATION 将被编译为 REST API，某些场景下我们希望某 OPERATION 仅用于内部（类似内部函数，不对外导出）。

指令 `@internalOperation` 用于实现该功能，其修饰的 OPERATION 不会编译为 REST API，只能在钩子中调用，详情见[文档](https://docs.fireboom.io/jin-jie-gou-zi-ji-zhi/nei-bu-tiao-yong)。

### 接口权限控制  @rbac

示例位于文件：`store/operation/n10hook/rbac`

对于中后台应用，不仅要控制接口的数据权限，还要控制接口只能被拥有某些角色的用户（管理员）访问。

**角色注入**

在身份验证钩子 `mutatingPostAuth` 中，可为用户注入角色，详情参考[文档](https://docs.fireboom.io/jin-jie-gou-zi-ji-zhi/shen-fen-yan-zheng-gou-zi#mutatingpostauth)

为用户注入角色后，firebom 会判断用户的角色与  [`@rbac` 指令](https://docs.fireboom.io/ji-chu-ke-shi-hua-kai-fa/yan-zheng-he-shou-quan/rbac)要求的角色是否匹配，确定是否可以访问该接口。

**动态权限**

`@rbac` 指令不够灵活，只适用于静态场景，在中后台应用中一般用钩子实现。

在 全局钩子 `beforeRequest` 中，可为接口动态设置需要的角色，详情参考 [示例代码](https://github.com/fireboomio/amis-admin/blob/c56897a2b9def58a00416b9b5df135f0259726e3/backend/custom-go/global/beforeOriginRequest.go#L32)。


# 参考

- [fireboom 文档中心](https://docs.fireboom.io/)