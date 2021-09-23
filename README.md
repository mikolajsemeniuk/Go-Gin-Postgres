# Go-Gin-Postgres

### Commands
```sh
go mod init Go-Gin-Postgres

go get github.com/gin-gonic/gin
go get github.com/lib/pq

export postgres_username=root
export postgres_password=P%40ssw0rd
export postgres_host=localhost
export postgres_port=15432
export postgres_database=db

# list all env variables
env

# to remove env variable
unset postgres_username
```
### Links
* [Query Multiple Rows](https://golang.org/doc/database/querying#multiple_rows)
* [Query Multiple Rows](https://www.calhoun.io/querying-for-multiple-records-with-gos-sql-package/)
* [Query Single Row](https://golang.org/doc/database/querying#single_row)
* [Query Single Row](https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/)