package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id   int
	Name string
}

func connect() (*sql.DB, error) {
	return sql.Open("postgres", "user=postgres dbname=postgres password=password sslmode=disable")
}

func GetTodos() ([]Todo, error) {

	db, err := connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id,name FROM todos")
	if err != nil {
		return nil, err
	}

	var res []Todo
	var todo Todo
	for rows.Next() {
		err = rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, todo)
	}

	return res, nil
}

func PostTodo(todo Todo) (Todo, error) {

	db, err := connect()
	if err != nil {
		return Todo{}, err
	}

	row := db.QueryRow("INSERT INTO todos(name) VALUES ($1) returning id,name", todo.Name)

	err = row.Scan(&todo.Id, &todo.Name)
	return todo, err
}

func DeleteTodo(id int) error {
	db, err := connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE from todos WHERE id = $1", id)

	return err
}
