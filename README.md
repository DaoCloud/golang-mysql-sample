Golang Web Application with Mysql connection
# Build Image
docker build -t daocloud/go-mysql .


## Below Mysql Connection Env and their default value
* MYSQL_PORT_27017_TCP_ADDR=localhost
* MYSQL_PORT_27017_TCP_PORT=27017
* MYSQL_USERNAME /* leave empty by default */
* MYSQL_PASSWORD /* leave empty by default */
* MYSQL_INSTANCE_NAME=test

# Run Container
docker run --link your_mysql:mysql -d -p 80:80 daocloud/go-mysql

# That's it
