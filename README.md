# Skooldio Go Robust API with Todolist

## Setup
1. go mod init <github_path>
2. run
```
go get gorm.io
go get github.com/glebarez/sqlite
go get github.com/gin-gonic/gin
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
go get golang.org/x/time/rate
go get github.com/gin-contrib/cors
```
3. run `go mod tidy`

To run the project, run `go run main.go` and it will create `test.db` automatically