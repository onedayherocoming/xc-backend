package service

import (
	"github.com/onedayherocoming/xc-backend/domian/model"
	"github.com/onedayherocoming/xc-backend/domian/repository"
)

type IPostDataService interface {
	AddPost(post *model.Post)(int64,error)
	DeletePost(ID int64)error
	UpdatePost(post *model.Post)(error)
	FindAllPost()([]model.Post,error)
	FindByTitle(string2 string)([]model.Post,error)
	FindByID(id int64)(*model.Post,error)
	UpdatePostContent(ID int64,content string)(error)
}

func NewPortDataService(postRepository repository.IPostRepository)(IPostDataService){
	return &PostDataService{portRepository: postRepository}
}

type PostDataService struct {
	portRepository repository.IPostRepository
}


func (p *PostDataService)AddPost(post *model.Post)(int64,error){
	return p.portRepository.CreatePost(post)
}
func (p *PostDataService)DeletePost(ID int64)error{
	return p.portRepository.DeletePostByID(ID)
}
func (p *PostDataService)UpdatePost(post *model.Post)(error){
	return p.portRepository.UpdatePost(post)
}

func (p *PostDataService)FindAllPost()([]model.Post,error){
	return p.portRepository.FindAll()
}
func (p *PostDataService)FindByTitle(title string)([]model.Post,error){
	return p.portRepository.FindPostByTitle(title)
}
func (p *PostDataService)FindByID(id int64)(*model.Post,error){
	return p.portRepository.FindPostByID(id)
}

func (p *PostDataService) UpdatePostContent(ID int64,content string)(error){
	model,err := p.FindByID(ID)
	if err!=nil{
		return err
	}
	model.Content=content
	err=p.UpdatePost(model)
	if err!=nil{
		return err
	}
	return nil
}
