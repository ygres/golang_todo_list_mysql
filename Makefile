.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: install
install:
	go mod init app
	go get github.com/BurntSushi/toml
	go get github.com/go-sql-driver/mysql
	go get github.com/golang-migrate/migrate/v4
	go get github.com/gorilla/mux

.PHONE: migrate
migrate:
	docker exec -t mysql mysql -uroot --password="112233" -e "use todolist; CREATE TABLE projects (id int NOT NULL AUTO_INCREMENT,title VARCHAR(255) UNIQUE,archived TINYINT,created_at DATETIME,updated_at DATETIME,deleted_at DATETIME,PRIMARY KEY (ID));" > /dev/null
	docker exec -t mysql mysql -uroot --password="112233" -e "use todolist; CREATE TABLE tasks (id int NOT NULL AUTO_INCREMENT,title VARCHAR(255),priority SMALLINT,deadline DATETIME,done TINYINT,project_id INT, created_at DATETIME,updated_at DATETIME,deleted_at DATETIME,PRIMARY KEY(ID));" > /dev/null

.PHONE: drop
drop:
	docker exec -t mysql mysql -uroot --password="112233" -e "use todolist; DROP Table projects">/dev/null
	docker exec -t mysql mysql -uroot --password="112233" -e "use todolist; DROP Table tasks">/dev/null
.DEFAULT_GOAL := build