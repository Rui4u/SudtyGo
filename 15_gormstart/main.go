package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	ID uint `gorm:"primarykey"`
	//Name string `gorm:"column:user_name; type:varchar(50); unique"`
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:sharui@tcp(10.211.55.5:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(10.211.55.5:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	//设置全局的logger。 这个logger在我们执行sql语句的时候，会打印每一行的sql
	//尽量看到每个api背后每个sql语句是什么

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "sr_", //所有表明前面都有这个前缀
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//// 不能设置默认值，是因为之设置一个值的时候，其他的是默认值，也就是空，非零值。  设置零值要通过 sql.NullString类型。 或者用*string来解决
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段  只有updates接口才会有影响
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

	// 迁移 schema
	//db.AutoMigrate(&User{})
	//user := User{Name: "sharui"}
	//db.Create(&user)
	//db.Model(&User{ID: 1}).Update("Name", "1") // 这种可以更新
	//db.Model(&user).Update("Name", "123")
	//db.Model(&User{ID: 1}).Update("Name", "")

	db.Model(&User{ID: 1}).Updates(User{Name: ""})
	//db.Delete(&User{Name: "sharui"}, 1)
}
