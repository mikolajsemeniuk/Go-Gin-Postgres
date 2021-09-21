# Go-Gin-Postgres

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