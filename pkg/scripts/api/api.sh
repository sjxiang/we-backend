
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
--header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjgsImVtYWlsIjoiMTUzNTQ4NDk0MEBxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTExVDAxOjQwOjI3Ljg2MDU4NjY2OSswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA2LTE3VDAxOjQwOjI3Ljg2MDU4NjgrMDg6MDAifQ.mmF8PccM4_mk-zA49K7o499OzttKkISsI7Cm-Phz4IY' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'


# 编辑用户信息
curl --location --request POST 'localhost:5678/api/v1/user/edit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "1535484940@qq.com",
    "password": "123456@qwe"
}'




=============================================================================


# *** 岗位 ***

# 新增岗位接口
curl --location --request POST 'localhost:2002/api/v1/post/add' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "post_name":   "pm",
    "post_status": 1,
    "post_code":   "63",
    "remark":      "m3"
}'

curl --location --request POST 'localhost:2002/api/v1/post/add' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "post_name": "engineer",
    "post_status": 1,
    "post_code": "99",
    "remark": "p8"
}'


# 查询岗位列表
curl --location --request GET 'localhost:2002/api/v1/post/list?page_size=10&&page_num=1&&post_status=2&&post_num=pm&&begin_time=&&end_time=' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' 


# 根据 id 查询岗位接口
curl --location --request GET 'localhost:2002/api/v1/post/info?id=10' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' 


# 修改岗位
curl --location --request PUT 'localhost:2002/api/v1/post/update' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 12,
    "post_name": "engineer",
    "post_status": 1,
    "post_code": "199",
    "remark": "manage"
}'


# 根据 id 删除岗位
curl --location --request DELETE 'localhost:2002/api/v1/post/delete/1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'


# 批量删除
curl --location --request DELETE 'localhost:2002/api/v1/post/batch/delete' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id_member": [
        10, 13
    ]
}'


# 修改状态
curl --location --request PUT 'localhost:2002/api/v1/post/updateStatus' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 10,
    "post_status": 2
}'


# 岗位下拉列表
curl --location --request GET 'localhost:2002/api/v1/post/vo/list' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'



# *** 部门 ***

# 新增部门
curl --location --request POST 'localhost:2002/api/v1/dept/add' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "parent_id":   0,
    "dept_type":   1,
    "dept_name":   "shenzhen dc",
    "dept_status": 1
}'

# 根据 id 查询部门
curl --location --request GET 'localhost:2002/api/v1/dept/info/11' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'

# 查询部门列表
curl --location --request POST 'localhost:2002/api/v1/dept/list' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dept_name":   "ui",
    "dept_status": 1
}'


# 修改部门
curl --location --request PUT 'localhost:2002/api/v1/dept/update' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":          11,
    "parent_id":   0,
    "dept_type":   1,
    "dept_name":   "shenzhen r&d",
    "dept_status": 1
}'

# 根据 id 删除部门
curl --location --request DELETE 'localhost:2002/api/v1/dept/delete/7' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'


# 部门下拉列表
curl --location --request GET 'localhost:2002/api/v1/dept/vo/list' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6ODksIlVzZXJuYW1lIjoiYWRtaW4iLCJFbWFpbCI6ImFkbWluQHFxLmNvbSIsIlBob25lIjoiMTMxMjIyMjMzMzMiLCJleHAiOjE3MTg5NTAxNDUsImlzcyI6ImFkbWluIn0.D6bJzCMkkvrD19BSexjhNP2coPXuWg1r2r5W9aKjqvo' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)'