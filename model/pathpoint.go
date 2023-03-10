package model

import (
	"time"
)

type Pathpoint struct {
	Id            uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	User_id       uint64    `json:"-"`
	Grid_point_id uint64    `json:"-"`
	Createdate    time.Time `gorm:"column:createdate"`
	Updatedate    time.Time `gorm:"column:updatedate"`
}

// 向数据库插入路径点
func (pathpoint *Pathpoint) Create() error {
	return DB.Mysql.Create(&pathpoint).Error
}

// func (pathpoint *Pathpoint) GetId() uint64 {
// 	t := &Referencepoint{}
// 	// TODO 添加查询失败时的处理
// 	DB.Mysql.Where("place_id = ? AND grid_point_id = ?",
// 		referencepoint.Place_id, referencepoint.Grid_point_id).Find(&t)
// 	return t.Id
// }

// TODO 经纬度和地址验证
// 结构体属性合法性校验
// 目前仅校验Username,Password,Usertype
// func (place *Place) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(user)
// }
