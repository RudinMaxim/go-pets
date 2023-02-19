package Cat

import "cats-server"

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsStrip   bool     `json:"is_strip" pg:"is_strip"`
	Color     string   `json:"color" pg:"color"`
}

// FindAllCats Получение котиков из таблицы

func FindAllCats() []Cat {
	var cats []Cat
	pgConnect := main.PostgresConnect()

	err := pgConnect.Model(&cats).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cats
}

// CreateCat Добовление котика в таблицу
func CreateCat(cat Cat) Cat {
	pgConnect := main.PostgresConnect()

	_, err := pgConnect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()

	return cat
}

func FindCatById(catid string) Cat {
	var cat Cat
	pgConnect := main.PostgresConnect()

	err := pgConnect.Model(&cat).Where("id = ?", catid).First()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}
func DropCatById(catId string) Cat {
	var cat Cat
	pgConnect := main.PostgresConnect()

	cat = FindCatById(catId)

	err := pgConnect.Model(&cat).Where("id = ?", catId).Deleted()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}
func UpdateCat(cat Cat) Cat {
	pgConnect := main.PostgresConnect()

	oldCat := FindCatById(cat.id)

	oldCat.Name = cat.Name
	oldCat.IsStrip = cat.IsStrip
	oldCat.Color = cat.Color

	_, err := pgConnect.Model(&oldCat).
		Set("name = ?", oldCat.Name).
		Set("is_strip = ?", oldCat.IsStrip).
		Set("color = ?", oldCat.Color).
		Where("id = ?", oldCat.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldCat
}
