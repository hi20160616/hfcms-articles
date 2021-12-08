# hfcms-api
The api for hfcms web crawler

# 1. gRPC

## 1.1. grpc-gateway
Refer:  
https://github.com/grpc-ecosystem/grpc-gateway  
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/  
https://www.cnblogs.com/FireworksEasyCool/p/12782137.html  

```
go mod init
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
go mod tidy
```

## 1.2. fix import "google/api/annotations.proto"  error
### Copy google/api/annotations.proto and google/api/http.proto
Refer:  
https://github.com/grpc-ecosystem/grpc-gateway/issues/1935  
```
mkdir -p google/api
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -O google/api/annotations.proto
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -O google/api/http.proto
```

## 1.3. Generate
### Article
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    api/articles/v1/hfcms-articles.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    api/articles/v1/hfcms-categories.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
    api/articles/v1/hfcms-tags.proto
```

# 2. MariaDB
Refer: https://hub.docker.com/\_/mariadb

## 2.1. docker
```
docker pull mariadb
docker run --detach --name hfcms-mariadb --env MARIADB_USER=example-user --env MARIADB_PASSWORD=my_cool_secret --env MARIADB_ROOT_PASSWORD=my-secret-pw  mariadb:latest

docker exec -it hfcms-mariadb mysql -u example-user -pmy_cool_secret
```

MARIADB\_USER=`example-user`
MARIADB\_PASSWORD=`my_cool_secret`
MARIADB\_ROOT\_PASSWORD=`my-secret-pw`

## 2.2 mariadb

1.Change root password:
```
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '[newpassword]';
```
2.Create tables:
Articles
```
mysql> DROP TABLE articles;
mysql> CREATE TABLE articles (id VARCHAR(24) NOT NULL, title VARCHAR(255), content TEXT(65535), category_id VARCHAR(10), user_id BIGINT(8), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
``` 
Categories
```
mysql> DROP TABLE categories;
mysql> CREATE TABLE categories (id VARCHAR(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255));
```
Tags
```
mysql> DROP TABLE tags;
mysql> CREATE TABLE tags (id BIGINT(8) NOT NULL AUTO_INCREMENT, name VARCHAR(255));
```
Attributes
```
mysql> DROP TABLE attributes;
mysql> CREATE TABLE attributes (id BIGINT(16) NOT NULL AUTO_INCREMENT, path VARCHAR(255), description VARCHAR(255), user_id BIGINT(8), article_id VARCHAR(24));
```

