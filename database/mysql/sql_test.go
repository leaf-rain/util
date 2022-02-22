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
		Addr:       "192.168.1.111:3306",
		User:       "root",
		Password:   "master",
		DbName:     "vchat_main",
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

	rows, err := client.Table("like0").Select("id", "like_uid").Where("uid = 20").Rows()
	fmt.Println("init err ===> ", err)
	defer rows.Close()

	for rows.Next() {
		var likeModelS = Like{}
		if err := rows.Scan(&likeModelS.Id, &likeModelS.LikeUid); err != nil {
			fmt.Println("err ===>", err)
		}
		fmt.Println(likeModelS)
	}
	if rows.Next() {

	} else {
		fmt.Println("====")
	}

}
