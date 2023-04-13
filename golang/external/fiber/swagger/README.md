```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
export PATH=$PATH:$HOME/go/bin
swag init -g app.go (default: main.go)
go get -u github.com/gofiber/swagger
```

 - 먼저 cmd를 통해 docs를 생성해야 한다.