package dao

import (
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/utils/db"
)

func GetHeroesList() (heroes []*models.HeroModel, err error) {
	DB := db.GetMysqlDB()
	query := `SELECT id, name, create_date, update_date FROM heroes`
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
	DB := db.GetMysqlDB()
	query := `INSERT INTO heroes(name) VALUE (?);`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(name)
	return
}

func UpdateHero(id int, name string) (err error) {
	DB := db.GetMysqlDB()
	query := `UPDATE heroes SET name = ?, update_date = CURRENT_TIMESTAMP WHERE id = ?`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(name, id)
	return
}
