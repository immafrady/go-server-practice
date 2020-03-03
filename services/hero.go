package services

import (
	"encoding/json"
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/models/dao"
	"fradyspace.com/go-server-practice/utils/request"
	"fradyspace.com/go-server-practice/utils/response"
	"net/http"
)

func HeroPostService(w http.ResponseWriter, req *http.Request) (err error) {
	// 读取
	var hero models.HeroModel
	err = json.Unmarshal(request.ReadJsonBody(req), &hero)
	if err != nil {
		return
	}
	// 操作
	err = dao.AddNewHero(hero.Name)
	if err != nil {
		return
	}
	// 返回
	resp := response.SuccessResponse(nil, "成功添加")
	b := resp.ToJson()
	_, err = w.Write(b)
	return
}

func HeroGetService(w http.ResponseWriter, req *http.Request) (err error) {
	heroes, err := dao.GetHeroesList()
	if err != nil {
		w.Write(response.DbError(err).ToJson())
		return
	}

	w.Write(response.SuccessResponse(heroes, "").ToJson())
	return
}

func HeroPutService(w http.ResponseWriter, req *http.Request) (err error) {
	hero := new(models.HeroModel)
	err = json.Unmarshal(request.ReadJsonBody(req), &hero)
	if err != nil {
		return
	}

	err = dao.UpdateHero(hero.Id, hero.Name)
	if err != nil {
		w.Write(response.DbError(err).ToJson())
		return
	}
	w.Write(response.SuccessResponse(nil, "更新成功").ToJson())
	return
}
