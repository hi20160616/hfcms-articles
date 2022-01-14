# 1. MariaDB
Refer: https://hub.docker.com/\_/mariadb

## 1.1. docker
```
docker pull mariadb
docker network create hfcms-mariadb
docker run --detach \
--publish 3306:3306 \
--env MARIADB_ROOT_PASSWORD=my-secret-pw \
--network hfcms-mariadb \
--name hfcms-mariadb \
mariadb:latest

docker exec -it hfcms-mariadb mysql -u root -pmy-secret-pw
```
MARIADB\_ROOT\_PASSWORD=`my-secret-pw`

## 1.2 mariadb

1. Change root password:
```
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '[newpassword]';
```
2. Creating database dumps
```
$ docker exec hfcms-mariadb sh -c 'exec mysqldump -uroot -p"$MARIADB_ROOT_PASSWORD" hfcms_articles' > ./hfcms_articles.sql
```
3. Restoring data from dump files
```
$ docker exec -i hfcms-mariadb sh -c 'exec mysql -uroot -p"$MARIADB_ROOT_PASSWORD" hfcms_articles' < ./hfcms_articles.sql
```
4. Create tables:  
Database
```
CREATE database hfcms_articles;
CREATE USER 'hfcms_articles_user'@localhost IDENTIFIED BY 'hfcms_articles_user_passwd';
GRANT ALL PRIVILEGES ON hfcms_articles.* TO 'hfcms_articles_user'@localhost;
FLUSH PRIVILEGES;
USE hfcms_articles;
```
Articles
```
DROP TABLE articles;
CREATE TABLE articles (id VARCHAR(24) NOT NULL, title VARCHAR(255), content TEXT(65535), category_id INT(10), user_id INT(10), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
describe articles;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | varchar(24)  | NO   | PRI | NULL                |                               |
| title       | varchar(255) | YES  |     | NULL                |                               |
| content     | mediumtext   | YES  |     | NULL                |                               |
| category_id | int(10)      | YES  |     | NULL                |                               |
| user_id     | int(10)      | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
``` 
Categories
```
DROP TABLE categories;
CREATE TABLE categories (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255), code VARCHAR(10), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
describe categories;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| name        | varchar(255) | YES  |     | NULL                |                               |
| code        | varchar(10)  | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
```
Tags
```
DROP TABLE tags;
CREATE TABLE tags (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
describe tags;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(10)      | NO   | PRI | NULL                | auto_increment                |
| name        | varchar(255) | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
```
Attributes
```
DROP TABLE attributes;
CREATE TABLE attributes (id INT(16) NOT NULL AUTO_INCREMENT, path VARCHAR(255), description VARCHAR(255), user_id INT(10), article_id VARCHAR(24), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
describe attributes;
+-------------+--------------+------+-----+---------------------+-------------------------------+
| Field       | Type         | Null | Key | Default             | Extra                         |
+-------------+--------------+------+-----+---------------------+-------------------------------+
| id          | int(16)      | NO   | PRI | NULL                | auto_increment                |
| path        | varchar(255) | YES  |     | NULL                |                               |
| description | varchar(255) | YES  |     | NULL                |                               |
| user_id     | int(10)      | YES  |     | NULL                |                               |
| article_id  | varchar(24)  | YES  |     | NULL                |                               |
| update_time | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+--------------+------+-----+---------------------+-------------------------------+
```
ArticleTags
```
DROP TABLE article_tags;
CREATE TABLE article_tags (id INT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), tag_id INT(10), UNIQUE KEY (id));
describe article_tags;
+------------+-------------+------+-----+---------+----------------+
| Field      | Type        | Null | Key | Default | Extra          |
+------------+-------------+------+-----+---------+----------------+
| id         | int(16)     | NO   | PRI | NULL    | auto_increment |
| article_id | varchar(24) | YES  |     | NULL    |                |
| tag_id     | int(10)     | YES  |     | NULL    |                |
+------------+-------------+------+-----+---------+----------------+
```
ArticleAttributes
```
DROP TABLE article_attributes;
CREATE TABLE article_attributes (id INT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), attribute_id INT(16), UNIQUE KEY (id));
describe article_attributes;
+--------------+-------------+------+-----+---------+----------------+
| Field        | Type        | Null | Key | Default | Extra          |
+--------------+-------------+------+-----+---------+----------------+
| id           | int(16)     | NO   | PRI | NULL    | auto_increment |
| article_id   | varchar(24) | YES  |     | NULL    |                |
| attribute_id | int(16)     | YES  |     | NULL    |                |
+--------------+-------------+------+-----+---------+----------------+
```

SQL All in one
```
CREATE database hfcms_articles;
CREATE USER 'hfcms_articles_user'@'%' IDENTIFIED BY 'hfcms_articles_user_passwd';
GRANT ALL PRIVILEGES ON hfcms_articles.* TO 'hfcms_articles_user'@'%';
FLUSH PRIVILEGES;
USE hfcms_articles;
CREATE TABLE articles (id VARCHAR(24) NOT NULL, title VARCHAR(255), content TEXT(65535), category_id INT(10), user_id INT(10), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
CREATE TABLE categories (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255), code VARCHAR(10), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
CREATE TABLE tags (id INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(255), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
CREATE TABLE attributes (id INT(16) NOT NULL AUTO_INCREMENT, path VARCHAR(255), description VARCHAR(255), user_id INT(10), article_id VARCHAR(24), update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, UNIQUE KEY (id));
CREATE TABLE article_tags (id INT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), tag_id INT(10), UNIQUE KEY (id));
CREATE TABLE article_attributes (id INT(16) NOT NULL AUTO_INCREMENT, article_id VARCHAR(24), attribute_id INT(16), UNIQUE KEY (id));
```
Validate
```
select host,
       user as username,
       password,
       password_expired
from mysql.user
order by user;
SHOW GRANTS FOR hfcms_articles_user;
describe articles;
describe categories;
describe tags;
describe attributes;
describe article_tags;
describe article_attributes;
```
