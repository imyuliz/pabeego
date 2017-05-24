package controllers

import (
	m "PagingBybeego/models"
	// "github.com/astaxie/beego"
)

type PersonController struct {
	BaseController
}

func (li *PersonController) AllPeople() {
//每页显示的记录数
	limit := 3
	//查询符合条件的总记录数
	count := m.FindCount()

	pager := li.SetPaginator(limit, count)
	person := m.FindInofoTwo(pager.Offset(), limit)
	li.Data["Allp"] = person
	li.TplName = "page.html"
}
