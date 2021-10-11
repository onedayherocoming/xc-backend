package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onedayherocoming/xc-backend/domian/model"
	"github.com/onedayherocoming/xc-backend/domian/service"
	"strconv"
)


func NewPostHandler(dataService service.IPostDataService)IPostHandler{
	return &PostHandler{
		PostService:dataService,
	}

}

type PostHandler struct {
	PostService service.IPostDataService
}

type IPostHandler interface {
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	QueryByTitle(c *gin.Context)
	QueryByID(c *gin.Context)
	QueryAll(c *gin.Context)
}

func( p *PostHandler) Add(c *gin.Context){
	content,err1 := c.GetPostForm("content")
	title,err2 := c.GetPostForm("title")
	userIDs,err3 := c.GetPostForm("user_id")
	if err1 == false{
		c.JSON(200,gin.H{
			"sucess":false,
			"msg":"content error",
		})
		return
	}
	if err2 == false{
		c.JSON(200,gin.H{
			"sucess":false,
			"msg":"title error",
		})
		return
	}
	if err3 == false{
		c.JSON(200,gin.H{
			"sucess":false,
			"msg":"user_id error",
		})
		return
	}
	userID,_ :=strconv.Atoi(userIDs)
	post := model.Post{
		Content: content,
		Title: title,
		UserID: int64(userID),
	}
	ID,err:=p.PostService.AddPost(&post)
	if err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
	}
	c.JSON(200,gin.H{
		"sucess":true,
		"id":ID,
	})
}

func( p *PostHandler) Delete(c *gin.Context){
	keys := c.Param("key")
	key,err := strconv.Atoi(keys)
	fmt.Println(keys,key)
	if  err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"msg":"key error",
		})
		return
	}
	err=p.PostService.DeletePost(int64(key))
	if err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"success":true,
	})
}

func( p *PostHandler) Update(c *gin.Context){
	content,err1 := c.GetPostForm("content")
	keys,err2:= c.GetPostForm("key")
	key,err3 := strconv.Atoi(keys)
	fmt.Println(key)
	if err1==false{
		c.JSON(200,gin.H{
			"success":false,
			"msg":"content error",
		})
		return
	}
	if err2==false || err3!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"msg":"key error",
		})
		return
	}

	err:=p.PostService.UpdatePostContent(int64(key),content)
	if err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"success":true,
	})
}

func( p *PostHandler) QueryByTitle(c *gin.Context){
	title,flag := c.GetQuery("title")
	if flag==false{
		c.JSON(200,gin.H{
			"success":false,
			"message":"title error",
		})
		return
	}
	models,err:=p.PostService.FindByTitle(title)
	if err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
		return
	}
	data:=gin.H{
		"result":models,
	}
	c.JSON(200,data)
}

func( p *PostHandler) QueryByID(c *gin.Context){
	keys,_ := c.GetQuery("key")
	key,err := strconv.Atoi(keys)
	fmt.Println(keys,key)
	if  err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"msg":"key error",
		})
		return
	}
	model,err:= p.PostService.FindByID(int64(key))
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
		return
	}
	data:=gin.H{
		"result":model,
	}
	c.JSON(200,data)
}

func( p *PostHandler) QueryAll(c *gin.Context){

	models,err:= p.PostService.FindAllPost()
	if err!=nil{
		c.JSON(200,gin.H{
			"success":false,
			"message":err,
		})
		return
	}
	data:=gin.H{
		"result":models,
	}
	c.JSON(200,data)
}
