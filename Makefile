build:
	docker build  --tag wow .
run:
	docker run wow -p 8080:8080
