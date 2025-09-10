package main

import (
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB


type Employees struct {
	ID uint
	Name string
	Department string
	Salary float64
}

type Books struct {
	ID uint
	Title string
	Author string
	Price float64
}

func initDB () {
	var err error

	db, err = sqlx.Open("sqlite3", "./test.db")

	if err != nil {
		log.Fatalf("connect db fail: %v", err)
	}

	schema := `
	create table if not exists employees (
		id INTEGER primary key autoincrement,
		name varchar(32),
		department varchar(32),
		salary REAL
	);`

	db.MustExec(schema)

	schema = `
	create table if not exists books (
		id INTEGER primary key autoincrement,
		title varchar(32),
		author varchar(32),
		price REAL
	);`
	db.MustExec(schema)
}

func insertData(employees [] Employees, books [] Books) {
	sql := "insert into employees (name, department, salary) values (?, ?, ?)"
	bsql := "insert into books (title, author, price) values (?, ?, ?)"
	for _, employee := range employees {
		db.MustExec(sql, employee.Name, employee.Department, employee.Salary)
	}
	for _, book := range books {
		db.MustExec(bsql, book.Title, book.Author, book.Price)
	}

		
}	

func query(dp string) []Employees {
	employees := []Employees{}
	sql := "select id, name,department,salary from employees where department = ?"
	db.Select(&employees, sql, dp)
	return employees
}

func queryTopSalary() []Employees {
	employees := []Employees{}
	sql := "select id, name,department,salary from employees order by salary desc limit 1"
	db.Select(&employees, sql)
	return employees
}
func queryAssignPriceBooks(price float64) []Books {
	books := []Books{}
	sql := "select id, title,author,price from books where price > ?"
	db.Select(&books, sql, price)
	return books
}

func main() {
	initDB()
	employees := make([]Employees, 0)
	employees = append(employees, 
		Employees{Name:"张三", Department: "技术部",Salary:9898.5 },
		Employees{Name:"张四", Department: "行政部",Salary:1898.5 },
		Employees{Name:"张吴", Department: "技术部",Salary:9898.5 },
		Employees{Name:"张六", Department: "技术部",Salary:19898.5 },
		Employees{Name:"张起", Department: "公关部",Salary:5898.5 },
	)

	books := make([]Books, 0)
	books = append(books, 
		Books{Title:"张三", Author: "技术部",Price:50.5 },
		Books{Title:"张四", Author: "行政部",Price:42.51 },
		Books{Title:"张吴", Author: "技术部",Price:999.5 },
		Books{Title:"张六", Author: "技术部",Price:121.5 },
		Books{Title:"张起", Author: "公关部",Price:44.5 },
	)

	insertData(employees, books)
	result := query("技术部")
	fmt.Printf("ep:%v\n", result)
	result = queryTopSalary()
	fmt.Printf("ep:%v\n", result)

	bks := queryAssignPriceBooks(50)
	fmt.Printf("books: %v\n", bks)
}