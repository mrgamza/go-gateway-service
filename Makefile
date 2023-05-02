NAME=gateway

build:
	docker build --no-cache -t ${NAME}-web -f Dockerfile-nginx .
	docker build --no-cache -t ${NAME}-service .

build-alpha:
	docker build --no-cache -t ${NAME}-web -f Dockerfile-nginx-alpha .
	docker build --no-cache -t ${NAME}-service .

deploy: build
	docker-compose up -d

up: down
	docker-compose up -d

run: build
	docker-compose up -d

debug: build-alpha
	docker-compose -f docker-compose.alpha.yml up

down:
	docker-compose down