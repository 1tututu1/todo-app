REST API for TODO list
===

Project description
---
 Golang application built using the Gin framework, logrus logging, PostgreSQL for data storage, Docker. This project provides CRUD (Create, Read, Update, Delete) operations for managing todo items, along with user authentication with JWT tokens to secure the data.

Features
---
* User authentication for securing todo items.
* Create, Read, Update, and Delete todo items.

About
---
Server \
Host: localhost \
Port: 3231 \
Db:   postgres

Routes
===
``` JAVA
    POST /auth/sign-in - sign-in user
```

``` JAVA
    POST /auth/register - register user
```

``` JAVA
    POST /api/lists - create list
```

``` JAVA
    GET /api/lists/1 - get list with id 1
```

``` JAVA
    PUT /api/lists/1 - update list with id 1
```

``` JAVA
    DELETE /api/lists/1 - delete list with id 1
```

``` JAVA
    GET /api/lists/1/items - get all tasks from list with id 1
```

``` JAVA
    GET /api/items/1 - get task with id 1
```

``` JAVA
    PUT /api/items/1 - update task with id 1
```

``` JAVA
    DELETE /api/items/1 - delete task with id 1
```