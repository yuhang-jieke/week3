package model

import (
	__ "week3/wei/user-server/handler/proto"

	"gorm.io/gorm"
)

type Goodss struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30);comment:商品名称"`
	Price float32 `gorm:"type:decimal(10,2);comment:商品价格"`
	Stock int     `gorm:"type:int(11);comment:商品库存"`
}

func (g *Goodss) GoodsAdd(db *gorm.DB) error {
	return db.Create(&g).Error
}
func (g *Goodss) GoodsShow(db *gorm.DB, id int) ([]*__.List, error) {
	var list []*__.List
	err := db.Model(&g).Where("id=?", id).Find(&list).Error
	return list, err
}
