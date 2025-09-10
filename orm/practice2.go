package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
	"gorm.io/gorm/clause"
)

type Accounts struct {
	Id int32
	Balance float64
	gorm.Model
}

type Transactions struct {
	Id uint
	FromAccountId int32
	ToAccountId int32
	Amount float64
	gorm.Model
}

var DB *gorm.DB
var Ctx context.Context


func transfer(source int32, target int32, amount float64) Accounts {
	var b Accounts
	result := DB.Transaction(func(tx * gorm.DB) error {
		result := DB.Model(Accounts{}).Where("Id = ? and balance - ? > 0", source, amount).UpdateColumn("balance", gorm.Expr("balance - ?", amount));
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected <= 0 {
			panic("更新失败：没有行受影响")
		}
		DB.Model(&b).Clauses(clause.Returning{}).Where("id = ?", target).UpdateColumn("balance", gorm.Expr("balance + ?", amount));
		gorm.G[Transactions](DB).Create(Ctx,&Transactions{FromAccountId:source,ToAccountId:target,Amount: amount} )
		return result.Error
	})
	if result != nil {
		panic("更新失败")
	}
	return b
}

func main() {
	db,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	DB = db

	Ctx = context.Background()
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})

	
	a := Accounts{Id:1, Balance: 110}
	b := Accounts{Id:2, Balance: 11}
	// db.Delete(&Accounts{}, []int32{1,2})
	// gorm.G[Accounts](db).Where("1 = 1").Delete(Ctx)
	db.Exec("DELETE FROM accounts")

	err = gorm.G[Accounts](db).Create(Ctx,&a )
	err = gorm.G[Accounts](db).Create(Ctx,&b )

	b = transfer(a.Id, b.Id, 100)

	allTransactions,err := gorm.G[Transactions](db).Where("1=1").Find(Ctx)
	fmt.Printf("b balance:%v ", b)
	fmt.Printf("b trans:%v ", allTransactions)
	
}