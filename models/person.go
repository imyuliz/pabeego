package models

import (
	"github.com/astaxie/beego/orm"
)

//一页显示的记录数
const PAGESIZE = 3

type Person struct {
	Pid      int    `orm:"pk;auto"`
	NickName string `orm:"size(30);index"`
	UserSex  string `orm:"size(5);null"`
	//用户邮箱
	UserEmail string `orm:"null"`
}

func FindAll() []*Person {
	o := orm.NewOrm()
	allper := make([]*Person, 0)
	o.Raw("select * FROM person").QueryRows(&allper)
	return allper
}

// FindCount 返回符合条件的总记录数
func FindCount() int64 {
	o := orm.NewOrm()
	var count int64
	o.Raw("SELECT COUNT(*) FROM person").QueryRow(&count)
	return count
}
// FindInfo 返回分页的数据
func FindInfo(page int) []*Person {
	o := orm.NewOrm()
	somePel := make([]*Person, 0)
	o.Raw("SELECT * FROM person ORDER BY pid DESC LIMIT ?,?",page,PAGESIZE).QueryRows(&somePel)
	
	return somePel
}

func FindInofoTwo(offset,limit int) []*Person {
	o:=orm.NewOrm()
	somePel := make([]*Person, 0)
	o.Raw("SELECT * FROM person ORDER BY pid DESC LIMIT ?,?",offset,limit).QueryRows(&somePel)
	
	return somePel
}

// type Page struct {
// 	//当前页数
// 	PageNo int
// 	//一页显示的记录数据
// 	PageSize int
// 	// 总共有多少页
// 	TotalPage int
// 	//一共有多少条记录
// 	TotalCount int
// 	//首页
// 	FirstPage bool
// 	//最后一页
// 	LastPage bool
// 	//每页显示的具体数据
// 	List interface{}
// }

// // PageUtil 分页函数
// // count总记录数
// // pageNo传过来的页面
// // pageSize 一页显示的记录数
// func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
// 	tp := count / pageSize
// 	if count%pageSize > 0 {
// 		tp = count/pageSize + 1
// 	}
// 	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
// }

/**
以上除了前台传递的pageNo，count又数据查到，pageSize由用户指定，list由数据库查询到的数据
**/

// 1是偏移量
// 5是显示多少条数据
// SELECT * FROM person   LIMIT 1,5

//符合条件的总记录数
// SELECT COUNT(*) FROM person
