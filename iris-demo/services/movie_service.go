package services

import (
	"fmt"
	"learngo/iris-demo/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManger struct {
	// 引入之前定义好的接口
	repo repositories.MovieRepository
}

func NewMovieServiceManger(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManger{repo}
}

func (m *MovieServiceManger) ShowMovieName() string {
	fmt.Println("我们获取到的视频名称" + m.repo.GetMovieName())
	return "我们获取到的视频名称" + m.repo.GetMovieName()
}
