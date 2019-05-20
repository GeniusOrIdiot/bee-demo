package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
	Age  int
}

func Create(user User) error {
	o := orm.NewOrm()

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	return err
}

func Get(id int) (u *User, e error) {
	o := orm.NewOrm()

	// query
	u = &User{Id: id}
	e = o.Read(u)

	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(u.Id, u.Name)
	}

	return u, e
}
