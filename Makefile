docker run --name=todo -e POSTGRES_PASSWORD='helloWorld5' -p 5436:5432 -d postgres
migrate -path ./migrations -database 'postgres://postgres:helloWorld5@localhost:5436/todo?sslmode=disable' up