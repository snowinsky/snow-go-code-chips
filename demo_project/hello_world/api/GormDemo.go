package api

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type GoPerson struct {
	ID             int       `gorm:"primarykey;column:id"`
	FirstName      string    `gorm:"column:first_name"`
	BirthDate      time.Time `gorm:"column:birth_date"`
	InsertDateTime time.Time `gorm:"column:insert_date"`
	CreateTime     time.Time `gorm:"column:create_timestamp"`
}

func (receiver GoPerson) TableName() string {
	return "t_go_person"
}

func GetDbConnection() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/app_vendor?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	log.Print("connect to mysql success")
	return conn
}

func GormDemoMain() {
	insertDemo()
	updateDemo()
	selectDemo()
}

func selectDemo() {
	conn := GetDbConnection()
	personArray := make([]GoPerson, 0)

	conn.Select("id", "first_name").Where("first_name=?", "One").Find(&personArray)
	log.Print(personArray)

}

func insertDemo() {
	//conn, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/app_vendor?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s")
	conn := GetDbConnection()

	//插入一条
	firstNewPerson := GoPerson{
		FirstName:      "One",
		BirthDate:      time.Date(2000, 11, 23, 0, 0, 0, 0, time.Local),
		InsertDateTime: time.Now(),
		CreateTime:     time.Now(),
	}
	var createErr = conn.Create(&firstNewPerson).Error
	if nil != createErr {
		panic("create data record error=" + createErr.Error())
	}
	//插入时ID不是程序指定，而是数据库行为，插入后，id值会被回写到对象中
	log.Print("insert complete", " id=", firstNewPerson.ID)
	//同时插入多条，注意不是批量插入多条
	var newPersonArray = []GoPerson{
		{
			FirstName:      "two",
			BirthDate:      time.Date(2001, 11, 23, 0, 0, 0, 0, time.Local),
			InsertDateTime: time.Now(),
			CreateTime:     time.Now(),
		},
		{
			FirstName:      "two",
			BirthDate:      time.Date(2001, 11, 23, 0, 0, 0, 0, time.Local),
			InsertDateTime: time.Now(),
			CreateTime:     time.Now(),
		},
	}
	createErr = conn.Create(&newPersonArray).Error
	if nil != createErr {
		panic("create data record error=" + createErr.Error())
	}
	//批量插入多条
	var newPersonArray1 = []GoPerson{
		{
			FirstName:      "three",
			BirthDate:      time.Date(2001, 11, 23, 0, 0, 0, 0, time.Local),
			InsertDateTime: time.Now(),
			CreateTime:     time.Now(),
		},
		{
			FirstName:      "three",
			BirthDate:      time.Date(2001, 11, 23, 0, 0, 0, 0, time.Local),
			InsertDateTime: time.Now(),
			CreateTime:     time.Now(),
		},
	}
	batchCreateErr := conn.CreateInBatches(&newPersonArray1, 2).Error
	if nil != batchCreateErr {
		panic("batch create date error=" + batchCreateErr.Error())
	}
}

func updateDemo() {
	conn := GetDbConnection()
	firstNewPerson := GoPerson{
		ID:             17,
		FirstName:      "One+3",
		BirthDate:      time.Date(2000, 11, 23, 0, 0, 0, 0, time.Local),
		InsertDateTime: time.Now(),
		CreateTime:     time.Now(),
	}
	//save 就是根据主键id进行更新，如果id不存在就插入，存在就更新;需要注意的是，更新的时候是全量更新，表中所有的字段都会被更新，所以需要慎重使用
	//并不是写几个字段更新几个字段
	if nil != conn.Save(firstNewPerson).Error {
		panic("save record fail")
	}

	//如果只想更新指定的几个字段
	conn.Model(&firstNewPerson).Where("id = ?", 17).UpdateColumn("first_name", "one+2+3").UpdateColumn("create_timestamp", time.Now()).Updates("first_name")

	//使用最原始的sql进行更新
	updateByNativeSql := GoPerson{}
	//select 用这个
	conn.Raw("update t_go_person p set p.first_name=? where p.id=?", "one+2+3+1", 3).Scan(&updateByNativeSql)
	//delete update insert 用这个
	conn.Exec("update t_go_person p set p.first_name=? where p.id=?", "one+2+3+4", 4)

}
