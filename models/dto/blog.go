package dto

import (
	"database/sql"
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/utils/parse"
	"time"
)

type PostDTO struct {
	Id         uint     `json:"id"`
	Title      string   `json:"title"`
	Introduce  string   `json:"introduce"`
	ImgUrl     string   `json:"imgUrl"`
	Tags       []TagDTO `json:"tags"`
	CreateDate string   `json:"createDate"`
	UpdateDate string   `json:"updateDate"`
}

type TagDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func PostDTO2Model(dto PostDTO) (model *models.Post, err error) {
	model = &models.Post{}
	model.ID = dto.Id
	model.Title = dto.Title
	model.Introduce = dto.Introduce
	model.ImgUrl = sql.NullString{String: dto.ImgUrl}
	model.CreatedAt = parse.String2Time(dto.CreateDate)
	model.UpdatedAt = parse.String2Time(dto.UpdateDate)

}

func PostModel2DTO(model models.Post) (dto PostDTO, err error) {

}
