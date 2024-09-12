PostgreSQl:

docker run --name=ttbd -e POSTGRES_PASSWORD=testtaskbd -p 8090:5432 -d --rm postgres

migrate -path ./schema -database postgres://postgres:testtaskbd@localhost:8090/postgres?sslmode=disable up
