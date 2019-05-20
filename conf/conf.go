package conf

import (
	"bee-demo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"))

	// set default database
	_ = orm.RegisterDataBase("default", "mysql", db, 30)

	// register model
	orm.RegisterModel(new(models.User))

	// create table
	_ = orm.RunSyncdb("default", false, true)
}
