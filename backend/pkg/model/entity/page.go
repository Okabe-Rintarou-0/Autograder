package entity

import "autograder/pkg/model/dbm"

type Page struct {
	PageSize int
	PageNo   int
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
