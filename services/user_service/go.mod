module github.com/SmartDuck9000/travelly-api/services/user_service

go 1.15

require (
	github.com/SmartDuck9000/travelly-api/config_reader v0.0.0-20210413133644-16ee7d6a243a
	github.com/SmartDuck9000/travelly-api/token_manager v0.0.0-20210413133644-16ee7d6a243a
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.7.1
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.7
)
