# 01-bbs

## prepare mysql server

```
docker run --name go-bbs -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:5.7
```

## Setup the value (run once)

```
mysql -h127.0.0.1 -uroot -proot -e "create database bbs"
mysql -h127.0.0.1 -uroot -proot bbs < ./sql/setup.sql
```

## prepare library

```
go get -u github.com/golang/dep/cmd/dep
dep ensure
```
