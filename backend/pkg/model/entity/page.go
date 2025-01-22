package entity

import "autograder/pkg/model/dbm"

type Page struct {
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
	PageNo   int `json:"page_no" form:"page_no" binding:"required"`
}

func (p *Page) ToDBM() *dbm.Page {
	if p == nil {
		return nil
	}
	return &dbm.Page{
		PageSize: p.PageSize,
		PageNo:   p.PageNo,
	}
}
