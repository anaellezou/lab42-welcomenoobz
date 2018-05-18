package main

import (
	"database/sql"
	"fmt"
)

type student struct {
	name string
	techs string
	pizza string
	music string
	role string
	inspirationnal_quote string
}

const table_name = "students"

func TestDB() {
	db, err := sql.Open("sqlite3", "./lab42.db")
	if err != nil {
		fmt.Printf("Error DB: %v\n", err)
	}
	sqlStmt := `
	create table students (id integer not null primary key,
		                  name text,
		                  techs text,
		                  pizza text,
		                  music text,
		                  role text,
		                  inspirationnal_quote text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Printf("%v\n", db)
	defer db.Close()
}

func CreateStudent(stud student) {
	sqlStmt := StrFormat(`insert into students ('name', 'techs', 'pizza', 'music', 'role', 'inspirationnal_quote')
		VALUES('{name}', '{techs}', '{pizza}', '{music}', '{role}', '{inspirationnal_quote}')`,
		"name", stud.name,
		"techs", stud.techs,
		"pizza", stud.pizza,
		"music", stud.music,
		"role", stud.role,
		"inspirationnal_quote", stud.inspirationnal_quote)

	db, err := sql.Open("sqlite3", "./lab42.db")
	if err != nil {
		fmt.Printf("Error DB: %v\n", err)
	}
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Printf("%v\n", db)
	defer db.Close()
}

func FindStudentBy(field string, value string) (student, error) {
	sqlStmt := StrFormat(`select * from {table_name} where {table_name}.{field}="{value}" limit 1;`,
		"table_name", table_name,
		"field", field,
		"value", value)
	db, err := sql.Open("sqlite3", "./lab42.db")
	if err != nil {
		fmt.Printf("Error DB: %v\n", err)
	}

	rows, err := db.Query(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		panic(err)
	}
	fmt.Printf("%v\n", db)
	for rows.Next() {
		var (
			id int
			name string
			techs string
			pizza string
			music string
			role string
			inspirationnal_quote string
		)

		if err := rows.Scan(&id, &name, &techs, &pizza, &music, &role, &inspirationnal_quote); err != nil {
            panic(err)
        }
        stud := student{name, techs, pizza, music, role, inspirationnal_quote}
		return stud, nil
	}
	defer db.Close()
	return student{}, nil
}

func GetAllStudents() ([]student, error) {
    sqlStmt := `select * from students;`
    db, err := sql.Open("sqlite3", "./lab42.db")
    if err != nil {
        fmt.Printf("Error DB: %v\n", err)
    }

    rows, err := db.Query(sqlStmt)
    if err != nil {
        fmt.Printf("%q: %s\n", err, sqlStmt)
        panic(err)
    }
    fmt.Printf("%v\n", db)
    var array_studs []student;
    for rows.Next() {
        var (
            id int
            name string
            techs string
            pizza string
            music string
            role string
            inspirationnal_quote string
        )

        if err := rows.Scan(&id, &name, &techs, &pizza, &music, &role, &inspirationnal_quote); err != nil {
            panic(err)
        }
        stud := student{name, techs, pizza, music, role, inspirationnal_quote}
        array_studs = append(array_studs, stud)
    }
    defer db.Close()
    return array_studs, nil
}