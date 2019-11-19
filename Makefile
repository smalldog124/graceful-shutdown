build :
	docker build -t sales-api:0.0.1 .

run :
	docker run --name sales-api -p 3000:3000 -d sales-api:0.0.1

stop :
	docker stop sales-api
	docker rm sales-api

logservice :
	docker logs sales-api -f

load-request :
	go run test/main.go