package repositories

import "learngo/iris-demo/datamodels"

// 定义一个接口
type MovieRepository interface {
	GetMovieName() string
}

// 定义封装类
type MovieManger struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManger{}
}

func (m *MovieManger) GetMovieName() string {
	// 模拟已经查询了数据库赋值给模型
	movie := &datamodels.Movie{Name: "YouTube视频"}

	return movie.Name

}
