

# 正常请求，示范

```bash
curl --location --request POST 'localhost:5678/api/v1/user/register' --header 'Content-Type: application/json' --data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe",
    "password_confirm": "123456@qwe"
}'
```

# 异常请求，示范

> 1. 不满足字段 tag（require 等等），参数校验失败

```bash
curl --location --request POST 'localhost:5678/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe",
    "password_confirm": "123456@qw"
}'
```

```json
{
    "code":100101,
    "message":"[{\"field\":\"password_confirm\",\"error\":\"password_confirm must be equal to Password\"}]"
}
```


> 2. invalid json 格式错误（标点符号遗漏），参数绑定，序列化失败

```bash
curl --location --request POST 'localhost:5678/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe",
}'
```

```json
{
    "code":100102,
    "message":"invalid character '}' looking for beginning of object key string"
}
```


> 3. invalid json 类型不匹配（string => int64），参数绑定，序列化失败

```bash
curl --location --request POST 'localhost:5678/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": 9
}'
```

```json
{
    "code":100102,
    "message":"json: cannot unmarshal number into Go struct field RegisterRequest.password of type string"
}
```