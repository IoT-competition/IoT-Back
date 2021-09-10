## 接口文档

### 基本说明

#### 接口通用格式

```json
{
    "code": "code",
    "data": null,
    "msg": "information"
}
```

#### URL 前缀

```shell
/api
```

### 信息相关

#### 获取

##### Route

```shell
GET /info
```

##### Request Data

```json
null
```

##### Response Data

```go
// 返回最新创建的10条信息(可以选择增加参数)
{
    "code": "0000",
    "data": [
        {
            "id": 10,
            "time": "2021-09-11T03:06:32.982+08:00",
            "illumination": "0",
            "temperature": "15",
            "humidity": "10"
        },
        ......
        {
            "id": 1,
            "time": "0001-01-01T00:00:00Z",
            "illumination": "0",
            "temperature": "0",
            "humidity": "0"
        }
    ],
    "msg": "处理成功"
}
```

#### 设置

##### Route

```shell
POST /info
```

##### Request Data

```json
{
    "illumination": "0",
    "temperature": "15",
    "humidity": "10"
}
```

##### Response Data

```go
{
    "code": "0000",
    "data": null,
    "msg": "处理成功"
}
```

### 操作相关

#### 获取

##### Route

```shell
GET /option
```

##### Request Data

```json
null
```

##### Response Data

```go
{
    "code": "0000",
    "data": {
        "led": 10,
        "fanN": 10,
        "fanM": 10,
        "engine": 10,
        "smoke": 10,
        "infrared": 10,
        "card": 10
    },
    "msg": "处理成功"
}
```

#### 设置

##### Route

```shell
POST /option
```

##### Request Data

```json
//值为0则表示不更新值  比默认值大 表示元素增加（正向）
{
    "led": 1,
    "fanN": 0,
    "fanM": 0,
    "engine": 0,
    "smoke": 0,
    "infrared": 5,
    "card": 4
}
```

##### Response Data

```go
{
    "code": "0000",
    "data": null,
    "msg": "处理成功"
}
```

