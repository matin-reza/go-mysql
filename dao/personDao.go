package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-mysql/entity"
	"strconv"
)

func connection() *sql.DB {
	db, err := sql.Open("mysql", "dbtest:DBtest@123@tcp(192.168.111.111:3306)/tempp")
	if err != nil {
		panic(err.Error())
	}

	return db
}
func Save(person entity.Person) {
	db := connection()
	if person.Id == 0 {
		query := "insert into person(name, email)values(?, ?)"
		_, err := db.Exec(query, person.Name, person.Email)
		if err != nil {
			panic("Could not insert into table person. please try again")
		}
	} else {
		query := "update person set name = ?, email = ? where id = ?"
		_, err := db.Exec(query, person.Name, person.Email, person.Id)
		if err != nil {
			panic("Could not update the entity  " + strconv.Itoa(person.Id))
		}
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
}

func Delete(id int) {
	db := connection()
	query := "delete from person where id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		panic("Could not delete the row with the id = " + strconv.Itoa(id))
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

}

func ReadAll() []entity.Person {
	var persons []entity.Person
	db := connection()
	query := "select * from person"
	results, err := db.Query(query)
	if err != nil {
		panic("Could not retrieve data from database")
	}
	for results.Next() {
		var person entity.Person
		err := results.Scan(&person.Id, &person.Name, &person.Email)
		if err != nil {
			panic("Could not retrieve data from database")
		}
		persons = append(persons, person)
	}
	return persons
}

func LoadById(id int) entity.Person {
	var person entity.Person
	db := connection()
	query := "select * from person where id = ?"
	row := db.QueryRow(query, id)
	err := row.Scan(&person.Id, &person.Name, &person.Email)
	if err != nil {
		panic("Could not retrieve data from database")
	}
	return person
}
