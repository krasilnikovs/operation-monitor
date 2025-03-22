up:
	docker compose up --build -d

down:
	docker compose down

ssh:
	docker exec -it operation-monitor-app-1 sh