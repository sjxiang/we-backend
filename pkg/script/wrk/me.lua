wrk.method="GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["User-Agent"] = "Apifox/1.0.0 (https://apifox.com)"
-- 记得修改这个，你在登录页面登录一下，然后复制一个过来这里
wrk.headers["Authorization"]="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NTksImVtYWlsIjoiMTUzNTQ4NDk0MEBxcS5jb20iLCJpc3N1ZWRfYXQiOiIyMDI0LTA2LTE0VDIzOjU4OjEwLjA5NzA2MjY4OSswODowMCIsImV4cGlyZWRfYXQiOiIyMDI0LTA2LTIwVDIzOjU4OjEwLjA5NzA2MjgxOSswODowMCJ9.DQwD7n7TYPcANaRftvn7UJUsRE_5QnsjFyKgGwd3vVo"