
TAG=${shell git describe --tag --long}
PREFIX=

db:
	docker rm -f wechatshop-db-test || true
	docker run -d --name wechatshop-db-test -v `pwd`/db_data:/var/lib/mysql -p 3306:3306 -e MYSQL_DATABASE=wechatshop -e MYSQL_ROOT_PASSWORD=123456 mysql
