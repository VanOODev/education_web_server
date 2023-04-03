package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"server/storages/list"
	"server/storages/slice"
	"strconv"
)

type Storage interface {
	Add(data int64) (index int64)
	Get(index int64) (data int64)
	Delete(index int64)
	String() string
}

type Config struct {
	Databasest struct {
		DatabaseUrl string `hcl:"Database_url"`
	} `hcl:"Database,block"`
	StorageTypest struct {
		StorageType string `hcl:"Storage_type"`
	} `hcl:"Storagetype,block"'`
}

var cfg Config
var l Storage = list.NewList()
var sl Storage = slice.NewSliceStorage()

func add(w http.ResponseWriter, r *http.Request) {
	// Получить значения a и b из запроса
	dataString := r.FormValue("data")
	data, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid value for parameter 'a'", http.StatusBadRequest)
		return
	}

	err = hclsimple.DecodeFile("config.hcl", nil, &cfg)
	if err != nil {
		fmt.Printf("Ошибка загрузки файла конфигурации: %v\n", err)
		return
	}

	conn, err := pgx.Connect(context.Background(), cfg.Databasest.DatabaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "mytable").Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Создаем таблицу, если она не существует
		conn.Exec(context.Background(), "CREATE TABLE mytable (id SERIAL PRIMARY KEY, data integer NOT NULL)")
	}
	switch r.FormValue("type") {

	case "Лист":
		fmt.Println("Введите значение")
		l.Add(data)

	case "Срез":
		fmt.Println("Введите значение")
		sl.Add(data)
	case "База данных":

		fmt.Println("Введите значение")
		fmt.Scan(&data)
		stmt := `INSERT INTO mytable (data) VALUES ($1)`
		_, err = conn.Exec(context.Background(), stmt, data)
		if err != nil {
			fmt.Println("Error inserting item:", err)
			return
		}
		fmt.Println("Item inserted successfully")
	}

}

func main() {
	conn, err := pgx.Connect(context.Background(), cfg.Databasest.DatabaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "mytable").Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Создаем таблицу, если она не существует
		conn.Exec(context.Background(), "CREATE TABLE mytable (id SERIAL PRIMARY KEY, data integer NOT NULL)")
	}

	// Регистрация обработчика add на URL-адресе /add
	http.HandleFunc("/add", add)

	// Запуск HTTP-сервера на порту 8080
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
