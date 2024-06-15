

# 拉取镜像
# docker pull bitnami/mysql:latest
# docker pull bitnami/redis:latest

# 登录控制台
# docker exec -it db sh
# mysql --host=127.0.0.1 --port=3306 --user=root --password=my-secret-pw
# docker exec -it cache sh
# redis-cli
# keys *
# get xxx


# 拉起容器
mysql:
	docker run -itd \
		--name db \
		-p 13306:3306 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		-e MYSQL_ROOT_PASSWORD=my-secret-pw \
		bitnami/mysql:latest

redis:
	docker run -itd \
		--name cache \
		-p 16379:6379 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		bitnami/redis:latest
	
# 测试
test:
	go test -count=1 -v ./...

# 启动
run:
	go run ./cmd/we/main.go 

# 压测
stress-register:
	wrk -t1 -d1s -c2 -s ./pkg/script/wrk/register.lua http://localhost:5678/api/v1/user/register

stress-login:
	wrk -t1 -d1s -c2 -s ./pkg/script/wrk/login.lua http://127.0.0.1:5678/api/v1/user/login

stress-me:
	wrk -t1 -d1s -c2 -s ./pkg/script/wrk/me.lua http://localhost:5678/api/v1/user/me


.PHONY: mysql redis test run stress-register stress-login stress-me