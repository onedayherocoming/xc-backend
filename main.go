package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/onedayherocoming/xc-backend/domian/repository"
	"github.com/onedayherocoming/xc-backend/domian/service"
	handler2 "github.com/onedayherocoming/xc-backend/handler"
)

func main() {

	r := gin.Default()
	//数据库
	db,err := gorm.Open("mysql","root:123456@/xc?charset=utf8&parseTime=true")
	if err!=nil{
		fmt.Println()
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	rp:= repository.NewPostRepository(db)
	err=rp.InitTable()
	if err!=nil{
		fmt.Println(err)
	}
	myhandler := handler2.NewPostHandler(service.NewPortDataService(repository.NewPostRepository(db)))
	r.GET("/getByID", myhandler.QueryByID)
	r.GET("/getByTitle", myhandler.QueryByTitle)
	r.GET("/getAll", myhandler.QueryAll)
	r.POST("/add",myhandler.Add)
	r.DELETE("/delete/:key",myhandler.Delete)
	r.PUT("/put",myhandler.Update)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
