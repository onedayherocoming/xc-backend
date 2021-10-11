package repository

import (
	gorm "github.com/jinzhu/gorm"
	"github.com/onedayherocoming/xc-backend/domian/model"
)


type IPostRepository interface {
	InitTable() error
	FindPostByTitle(string)([]model.Post,error)
	FindPostByID(id int64)(*model.Post,error)
	CreatePost(post *model.Post)( int64,error)
	DeletePostByID(int64)error
	UpdatePost(post *model.Post)error
	FindAll()([]model.Post,error)
}

//创建PostRepository
func NewPostRepository(db *gorm.DB)IPostRepository{
	return &PostRepository{mysqlDb: db}
}

type PostRepository struct {
	mysqlDb *gorm.DB
}


func (p *PostRepository)InitTable() error{
	return p.mysqlDb.CreateTable(&model.Post{}).Error
}

func (p *PostRepository)FindPostByTitle(title string)(res []model.Post,err error){
	return res,p.mysqlDb.Where("title LIKE ?",title).Find(&res).Error
}

func (p *PostRepository)FindPostByID(id int64)(res *model.Post,err error){
	res = &model.Post{}
	return res,p.mysqlDb.First(res,id).Error
}

func (p *PostRepository)CreatePost(post *model.Post)( int64, error){
	return post.ID, p.mysqlDb.Create(post).Error
}
func (p *PostRepository)DeletePostByID(id int64)error{
	return p.mysqlDb.Where("id=?",id).Delete(&model.Post{}).Error
}
func (p *PostRepository)UpdatePost(post *model.Post)error{
	return p.mysqlDb.Model(post).Update(post).Error
}
func (p *PostRepository)FindAll()(res []model.Post,err error){
	return res,p.mysqlDb.Find(&res).Error
}