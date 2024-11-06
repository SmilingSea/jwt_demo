---
title: jwt_demo
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# jwt_demo

Base URLs:

* <a href="http://localhost:8888">开发环境: http://localhost:8888</a>

# Authentication

# Default

## GET 测试接口

GET /auth/ping

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
    "code": 200,
    "expire": "2024-11-07T00:15:04+08:00",
    "message": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA5MDk3MDQsIm9yaWdfaWF0IjoxNzMwOTA2MTA0fQ.Z2l-cxcox8urysXGJzKWrits8eaya6raBubONPLFAP0"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|



## POST 登录-获取token

POST /login

> Body 请求参数

```json
{
  "Account": "jhh",
  "Password": "jhh72811"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» Username|body|string| 是 |none|
|» Password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
    "message": "pong"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|



