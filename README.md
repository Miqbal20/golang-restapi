# golang-restapi

### Step by Step 
```
go mod init github.com/Miqbal20/golang-api-template
```

```
go get github.com/gin-gonic/gin@latest 
go get gorm.io/gorm@latest 
go get grom.io/driver/mysql@latest
```

### Models dan Controllers
1. Buat file models/setup.go => buat setup
2. Buat file models/___.go => buat model ___
3. Edit main.go => Buat Route untuk API
4. buat file controllers/___Controller/___Controller.go
5. Install Swagger
```
go get -u github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger

```

```
go run main.go
```

```
vercel .
```
