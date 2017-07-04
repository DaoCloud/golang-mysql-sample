package main

import (
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func Insert(person *Person) {

	if person.Name == "" || GetResult(person.Name) != "" {
		panic("insert conflict")
	}

	stmt, err := db.Prepare(peopleInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&person.Name,
		&person.Phone,
	)
	if err != nil {
		panic(err)
	}

}

func List() []*Person {
	rows, err := db.Query(peopleQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	list := []*Person{}
	for rows.Next() {
		m := new(Person)
		if err := rows.Scan(
			&m.Name,
			&m.Phone,
		); err != nil {
			log.Println(err)
			panic(err)
		}

		list = append(list, m)
	}

	return list
}

func GetResult(name string) string {
	rows, err := db.Query(peopleQuery+" WHERE username = ?", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	list := []*Person{}
	for rows.Next() {
		m := new(Person)
		if err := rows.Scan(
			&m.Name,
			&m.Phone,
		); err != nil {
			panic(err)
		}

		list = append(list, m)
	}

	if len(list) > 0 {
		return list[0].Name
	}
	return ""
}

var peopleInsert = `
INSERT INTO
    people(
        username,
        phone
    )
VALUES(?, ?)
`
var peopleQuery = `
SELECT
    username,
    phone
FROM
    people
`
