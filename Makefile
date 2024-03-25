run:
	go run .

test:
	go test -cover -v ./...

start-postgres:
	docker run -d --name postgres -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 postgres

curl-episodes:
	curl http://localhost:8080/episodes | jq .

query-postgres-test:
	docker exec -it postgres psql -U $(DB_USERNAME) -d $(DB_NAME)  -c "SELECT * FROM jeopardy_game_box_scores"