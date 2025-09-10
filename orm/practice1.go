package main

import(
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

type Student struct {
	gorm.Model
	Id int32
	Name string
	Age int32
	Grade string
}

func first() {
	db,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	student := Student{Id:1,Name:"张三", Age: 10, Grade: "三年级"}
	ctx := context.Background()
	db.AutoMigrate(&Student{})

	err = gorm.G[Student](db).Create(ctx,&student )

	result := gorm.WithResult()
	student.Id = 2
	err = gorm.G[Student](db, result).Create(ctx, &student)
	
	// fmt.Printf("student,%v\n",student)
	// fmt.Printf("student,%v,%v\n",student.Id,result)

	studentRes, err := gorm.G[Student](db).Where("id = ?", 1).First(ctx)
	fmt.Printf("studentRes,%v\n",studentRes)
	studentRes, err = gorm.G[Student](db).Where("id = ?", 2).First(ctx)
	fmt.Printf("studentRes,%v\n",studentRes)
	studentRes, err = gorm.G[Student](db).Where("id = ?", 3).First(ctx)
	fmt.Printf("studentRes,%v\n",studentRes)

	gorm.G[Student](db).Where("id = ?", 1).Update(ctx, "Name", "王朝")
	studentRes, err = gorm.G[Student](db).Where("id = ?", 1).First(ctx)
	fmt.Printf("studentRes,%v\n",studentRes)

	gorm.G[Student](db).Where("id = ?", 1).Updates(ctx, Student{Age: 20, Grade: "四年级"})
	gorm.G[Student](db).Where("id = ?", 2).Updates(ctx, Student{Age: 24, Grade: "⑤年级"})
	studentRes, err = gorm.G[Student](db).Where("id = ?", 1).First(ctx)
	fmt.Printf("studentRes,%v\n",studentRes)

	gtEighteenStudent, err := gorm.G[Student](db).Where("age > 18").Find(ctx)
	fmt.Printf("gtEighteenStudent,%v\n",gtEighteenStudent)
}

func main() {

	first()
	
}