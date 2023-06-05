package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

var (
	dsn   = "root:AOOB.Pxc.db@tcp(81.71.119.20:6556)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)

type test struct {
	ID      int
	Name    string `gorm:"type:varchar(255);not null"`
	Sex     string `gorm:"type:varchar(255);not null"`
	Iphone  string `gorm:"type:varchar(255);not null"`
	Email   string `gorm:"type:varchar(255);not null"`
	Address string `gorm:"type:varchar(255);not null"`
	Remark  string `gorm:"type:varchar(255);not null"`
}

func (test) TableName() string {
	return "test"
}

func Test_Install(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 5000; i++ {
				user := test{
					Name: "test" + strconv.Itoa(i),
					Sex:  "男",
				}
				db.Create(&user)
			}
		}()
	}
	wg.Wait()
}

func Test_Select(t *testing.T) {
	var users []test
	db.Limit(100000).Find(&users)
}

func Test_Select_Sync(t *testing.T) {
	runtime.GOMAXPROCS(100)
	wg := sync.WaitGroup{}
	var users [20][]test
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			db.Offset(index * 5000).Limit(5000).Find(&users[index])
		}(i, &wg)
	}
	wg.Wait()
	t.Log(len(users[20-1]))
}
