MAIN_FILE = ./cmd/MyDocu/main.go
OUT_FILE = MyDocu

run:
	go run ${MAIN_FILE}

build:
	go build -o ${OUT_FILE} ${MAIN_FILE}

postgres:
	docker run --rm --name MyDocu-postgres \
		-p 5432:5432 \ 
		-e POSTGRES_USER=user \
		-e POSTGRES_PASSWORD=1234 \
		-v /postgres-volume:/var/lib/postgresql/data
		-d postgres

createdb:
	docker exec -it MyDocu-postgres createdb --username=user --owner=user MyDocuDB

docker:
	docker build -t yosaa5782/mydocu .
	docker run --rm --name mydocu-api -p 5000:5000 -d yosaa5782/mydocu