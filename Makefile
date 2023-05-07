gen:
	docker-compose exec backend go generate ./... && swag init && swag fmt
