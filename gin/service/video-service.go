package service

import "github.com/jishnup2930/golang.git/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoServise struct {
	videos []entity.Video
}

func New() VideoService
