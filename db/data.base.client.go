package db

import (
	"fmt"
	"github.com/Wilddogmoto/Oauth/configuration"
	"github.com/Wilddogmoto/Oauth/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Account struct {
		Server   string
		User     string
		Password string
		DB       string
	}
)

var (
	DataBase *gorm.DB
)

func DataBaseConnected()  {

	var (
		log = logger.Logger
		err error
		/*config = Account{
			//c:= "wild:777@tcp(127.0.0.1:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
			Server:   "127.0.0.1:3306",
			User:     "wild",
			Password: "777",
			DB:       "mytestdb",
		}*/
		connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&amp",
			configuration.Parameters.DB.User,
			configuration.Parameters.DB.Password,
			configuration.Parameters.DB.Server,
			configuration.Parameters.DB.Dbname,
		)
	)

	log.Infof(" >>> connect DB server: %v >>>",configuration.Parameters.DB.Server)

	DataBase, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatalf("bad db connection, error from DataBaseConnect:%v",err)
		return
	}

	log.Info("Database connect: success")
}
