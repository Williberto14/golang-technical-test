# golang-technical-test

## Create database in docker with mariadb

```bash
docker run --name golang-technical-test -p 127.0.0.1:3306:3306 -e MYSQL_ROOT_PASSWORD=qwerty -e MYSQL_DATABASE=golang_technical_test -d mariadb:latest
```
