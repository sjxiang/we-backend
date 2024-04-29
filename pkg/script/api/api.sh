
# 健康检查
curl --location --request GET 'localhost:8000/api/v1/health' 

# 用户注册
curl --location --request POST 'localhost:8000/api/v1/user/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484943@qq.com",
    "password": "123456@qwe",
    "confirm_password": "123456@qwe"
}'

# 用户登录
curl --location --request POST 'localhost:8000/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484943@qq.com",
    "password": "123456@qwe"
}'

# 用户详情


# 编辑用户信息
