CONTAINER_NAME=mysql
DB_NAME=launchPad

mysql:
	docker run --platform linux/amd64 -p 3306:3306 --env MYSQL_DATABASE=${DB_NAME} --env MYSQL_ROOT_PASSWORD=root --name ${CONTAINER_NAME} -d mysql

.PHONY: mysql