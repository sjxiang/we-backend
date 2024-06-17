
# 健康检查
curl --location --request GET 'localhost:5678/api/v1/user/health' 


# 用户注册
curl --location --request POST 'localhost:5678/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe",
    "password_confirm": "123456@qwe"
}'


# 用户登录
curl --location --request POST 'localhost:5678/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe"
}'


# 用户详情
curl --location --request GET 'localhost:5678/api/v1/user/me' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjMsImVtYWlsIjoiMTUzNTQ4NDk0MEBxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTE3VDIwOjE3OjQzLjE3OTMzNDE4OSswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA2LTIzVDIwOjE3OjQzLjE3OTMzNDMyOSswODowMCJ9.fWp2iyiFAuGUguY0zS9NWOC3xzWEjqZ0auWBJ6LoJ-c' 


# 编辑用户信息
curl --location --request POST 'localhost:5678/api/v1/user/edit' \
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjMsImVtYWlsIjoiMTUzNTQ4NDk0MEBxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTE1VDE2OjEwOjI1LjYxMTYxMDU5NiswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA2LTIxVDE2OjEwOjI1LjYxMTYxMDczNiswODowMCJ9.C5XLRCE99xGsbqdNEkSikU6ncBu1XOPlaAiQUP1jG6k' \
--header 'Content-Type: application/json' \
--data-raw '{
    "avatar": "jisoo.jpeg",
    "intro": "kpop idol",
    "birthday": "2002-08-29",
    "nickname": "blackpink_jisoo"
}'


