package Dog

type Dog struct {
	tableName  struct{} `pg:"dogs"`
	ID         string   `json:"id" pg:"id"`
	Name       string   `json:"name" pg:"name"`
	Breed      string   `json:"breed" pg:"breed"`
	Color      string   `json:"color" pg:"color"`
	ear_length int64    `json:"ear_length" pg:"ear_length"`
}

func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := main.PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}
func CreateDog(dogs Dog) Dog {
	pgConnect := main.PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()

	return dog
}

func FindDogById(dogid string) Dog {
	var dog Dog
	pgConnect := main.PostgresConnect()

	err := pgConnect.Model(&dog).Where("id = ?", dogid).First()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}
func DropDogById(dogId string) Dog {
	var dog Dog
	pgConnect := main.PostgresConnect()

	dog = FindDogById(dogId)

	err := pgConnect.Model(&dog).Where("id = ?", dogId).Deleted()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}
func UpdateDog(dog Dog) Dog {
	pgConnect := main.PostgresConnect()

	oldDog := FindDogById(dog.id)

	oldDog.Name = dog.Name
	oldDog.Breed = dog.Breed
	oldDog.Color = dog.Color
	oldDog.ear_length = dog.ear_length

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("breed = ?", oldDog.Breed).
		Set("color = ?", oldDog.Color).
		Set("ear_length = ?", oldDog.ear_length).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldDog
}
