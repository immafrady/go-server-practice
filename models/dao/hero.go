package dao

import (
	"errors"
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/utils/db"
)

func GetHeroesList() (heroes []*models.HeroModel, err error) {
	DB := db.GetMysqlDB()
	query := `SELECT id, name, create_date, update_date FROM heroes;`
	heroes = make([]*models.HeroModel, 0)
	rows, err := DB.Query(query)
	if err != nil {
		return
	}
	for rows.Next() {
		hero := new(models.HeroModel)
		if err := rows.Scan(&hero.Id, &hero.Name, &hero.CreateDate, &hero.UpdateDate); err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}
	return
}

func AddNewHero(name string) (err error) {
	_, err = db.ExecDb(`INSERT INTO heroes (name) VALUE (?);`, name)
	if err != nil {
		return
	}
	return
}

func UpdateHero(id int, name string) (err error) {
	_, err = db.ExecDb(`UPDATE heroes SET name = ?, update_date = CURRENT_TIMESTAMP WHERE id = ?;`, name, id)
	if err != nil {
		return
	}
	return
}

func DeleteHero(id int) (err error) {
	result, err := db.ExecDb(`DELETE FROM heroes where id = ?;`, id)
	if err != nil {
		return
	}
	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("没有删除项目")
	}
	return
}
