
# 健康检查
curl --location --request GET 'localhost:5678/api/v1/user/health' 


# 用户注册
curl --location --request POST 'localhost:5678/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484943@qq.com",
    "password": "123456@qwe",
    "password_confirm": "123456@qwe"
}'


# 用户登录
curl --location --request POST 'localhost:5678/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484943@qq.com",
    "password": "123456@qwe"
}'


# 用户详情
curl --location --request GET 'localhost:5678/api/v1/user/me' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsImVtYWlsIjoiMTUzNTQ4NDk0M0BxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTMwVDE1OjQ5OjA1LjY4MzMwMDE5MiswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA3LTA2VDE1OjQ5OjA1LjY4MzMwMDI1MiswODowMCJ9.rVgQ3_dWi3PWzZfJrpEEoT3RW_l58C1g4TLY2fAqy9A' 


# 编辑用户信息
curl --location --request POST 'localhost:5678/api/v1/user/edit' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsImVtYWlsIjoiMTUzNTQ4NDk0M0BxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTMwVDE1OjQ5OjA1LjY4MzMwMDE5MiswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA3LTA2VDE1OjQ5OjA1LjY4MzMwMDI1MiswODowMCJ9.rVgQ3_dWi3PWzZfJrpEEoT3RW_l58C1g4TLY2fAqy9A' \
--header 'Content-Type: application/json' \
--data-raw '{
    "avatar": "jisoo.jpeg",
    "intro": "kpop idol",
    "birthday": "2002-08-29",
    "nickname": "blackpink_jisoo"
}'


# 验证码请求登录
curl --location --request POST 'localhost:5678/api/v1/user/login_sms/otp/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone_number": "sjxiang2020@outlook.com"
}'

# 验证码校验登录
curl --location --request POST 'localhost:5678/api/v1/user/login_sms/otp/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone_number": "sjxiang2020@outlook.com",
    "input_code": ""

}'


# 用户列表
curl --location --request GET 'localhost:5678/api/v1/user/admin' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjUsImVtYWlsIjoiMTUzNTQ4NDk0M0BxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTMwVDE1OjQ5OjA1LjY4MzMwMDE5MiswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA3LTA2VDE1OjQ5OjA1LjY4MzMwMDI1MiswODowMCJ9.rVgQ3_dWi3PWzZfJrpEEoT3RW_l58C1g4TLY2fAqy9A' 