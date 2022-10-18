package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

var client *gorm.DB

func TestMain(m *testing.M) {
	client = NewMysql(&Config{
		Addr:       "127.0.0.1:3306",
		User:       "root",
		Password:   "password",
		DbName:     "gateway",
		Parameters: "charset=utf8mb4&parseTime=true&loc=Local",
	})
	now := time.Now()
	m.Run()
	fmt.Println("代码执行耗时：", time.Since(now))
}

type Like struct {
	Id      int32  `gorm:"column:id; type:int; not null; primary_key" json:"id"`
	Uid     uint64 `gorm:"column:uid; type:int; not null; defalut:0" json:"uid"`
	LikeUid uint64 `gorm:"column:like_uid; type:int; not null; defalut:0" json:"like_uid"`
}

func TestMysql(t *testing.T) {
	var sql = "select * from gateway.method"
	rows, err := client.Raw(sql).Rows()
	fmt.Println("init err ===> ", err)
	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows)
	}
	if rows.Next() {
		fmt.Println(rows)
	} else {
		fmt.Println("====")
	}

}
