# Go Todo REST API Example with MySQL
A RESTful API example for simple todo application with Go


## Installation & Run


```bash
cd golang_todo_list_mysql
docker-compose up -d
make install
make migrate
make build
./apiserver

# API Endpoint : http://127.0.0.1:3000
```


## API

#### /projects
* `GET` : Get all projects

 ```curl --location --request GET 'http://127.0.0.1:3000/projects'```
* `POST` : Create a new project

```curl --location --request POST 'http://127.0.0.1:3000/projects' -H 'Content-Type: application/json' --data-raw '{"title": "first"}'```

#### /projects/:title
* `GET` : Get a project

```curl --location --request GET 'http://127.0.0.1:3000/projects/[title]'```
* `PUT` : Update a project

```curl --location --request PUT 'http://127.0.0.1:3000/projects/[title]' -H 'Content-Type: application/json' --data-raw '{"title": "first", "archived": true }'```
* `DELETE` : Delete a project

```curl --location --request DELETE 'http://127.0.0.1:3000/projects/[title]'```


#### /tasks
* `GET` : Get all tasks

```curl --location --request GET 'http://127.0.0.1:3000/tasks'```


#### /task/:id
* `PUT` : Update a task 

```curl --location --request PUT 'http://127.0.0.1:3000/task/[id]' -H 'Content-Type: application/json' --data-raw '{"title": "testing 3 ...", "priority": 5, "deadline": "2022-09-23T10:49:47Z", "done": false}'```

* `DELETE` : Delete a task 

```curl --location --request DELETE 'http://127.0.0.1:3000/task/2'```


#### /projects/:title/task
* `POST` : Create a new task in a project

```curl --location --request POST 'http://127.0.0.1:3000/projects/[title]/task' -H 'Content-Type: application/json' --data-raw '{"title": "testing 3 ...", "priority": 5, "deadline": "2022-09-22T10:49:47Z", "done": false}'```

* `GET` : Get a task of a project

```curl --location --request GET 'http://127.0.0.1:3000/projects/[title]/task'```


