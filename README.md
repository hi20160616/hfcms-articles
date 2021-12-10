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

docker exec -it hfcms-mariadb mysql -u root -pmy-secret-pw
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
Database
```
CREATE database articles;
CREATE USER 'articles_user'@localhost IDENTIFIED BY 'articles_user_passwd';
GRANT ALL PRIVILEGES ON articles.* TO 'articles_user'@localhost;
FLUSH PRIVILEGES;
USE articles;
```
Articles
```
DROP TABLE articles;
CREATE TABLE articles (id VARCHAR(24) NOT NULL, title VARCHAR(255), content TEXT(65535), category_id VARCHAR(10), user_id INT(10), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
describe articles;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | varchar(24)  | NO   | PRI | NULL                |                               |
| title       | varchar(255) | YES  |     | NULL                |                               |
| content     | mediumtext   | YES  |     | NULL                |                               |
| category_id | varchar(10)  | YES  |     | NULL                |                               |
| user_id     | int(10)      | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
``` 
Categories
```
DROP TABLE categories;
CREATE TABLE categories (id VARCHAR(10) NOT NULL , name VARCHAR(255), UNIQUE KEY (id));
describe categories;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| id    | varchar(10)  | NO   | PRI | NULL    |       |
| name  | varchar(255) | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
```
Tags
```
DROP TABLE tags;
CREATE TABLE tags (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255), UNIQUE KEY (id));
describe tags;
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int(10)      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(255) | YES  |     | NULL    |                |
+-------+--------------+------+-----+---------+----------------+
```
Attributes
```
DROP TABLE attributes;
CREATE TABLE attributes (id BIGINT(16) NOT NULL AUTO_INCREMENT, path VARCHAR(255), description VARCHAR(255), user_id INT(10), article_id VARCHAR(24), UNIQUE KEY (id));
describe attributes;
+-------------+--------------+------+-----+---------+----------------+
| Field       | Type         | Null | Key | Default | Extra          |
+-------------+--------------+------+-----+---------+----------------+
| id          | bigint(16)   | NO   | PRI | NULL    | auto_increment |
| path        | varchar(255) | YES  |     | NULL    |                |
| description | varchar(255) | YES  |     | NULL    |                |
| user_id     | int(10)      | YES  |     | NULL    |                |
| article_id  | varchar(24)  | YES  |     | NULL    |                |
+-------------+--------------+------+-----+---------+----------------+
```
ArticleTags
```
DROP TABLE article_tags;
CREATE TABLE article_tags (id BIGINT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), tag_id INT(10), UNIQUE KEY (id));
describe article_tags;
+------------+-------------+------+-----+---------+----------------+
| Field      | Type        | Null | Key | Default | Extra          |
+------------+-------------+------+-----+---------+----------------+
| id         | bigint(16)  | NO   | PRI | NULL    | auto_increment |
| article_id | varchar(24) | YES  |     | NULL    |                |
| tag_id     | int(10)     | YES  |     | NULL    |                |
+------------+-------------+------+-----+---------+----------------+
```
ArticleAttributes
```
DROP TABLE article_attributes;
CREATE TABLE article_attributes (id BIGINT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), attribute_id BIGINT(16), UNIQUE KEY (id));
describe article_attributes;
+--------------+-------------+------+-----+---------+----------------+
| Field        | Type        | Null | Key | Default | Extra          |
+--------------+-------------+------+-----+---------+----------------+
| id           | bigint(16)  | NO   | PRI | NULL    | auto_increment |
| article_id   | varchar(24) | YES  |     | NULL    |                |
| attribute_id | bigint(16)  | YES  |     | NULL    |                |
+--------------+-------------+------+-----+---------+----------------+
```
