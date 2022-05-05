package model

import (
	"database/sql"
	"time"
	"top.lel/ginapp01/tool"
)

// UserInfo 实体类
// 默认使用snake_name 蛇形命名法来和表字段映射
// 也可以使用 gorm:"column:field"指定对应的列
type UserInfo struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name"`
	Email    *string    `json:"email"`
	Age      uint8      `json:"age"`
	Birthday *time.Time `json:"birthday"`
	// no read,write,migration permission...
	MemberNumber sql.NullString `gorm:"-:all" json:"-"`
	ActivatedAt  tool.Date      `json:"activatedAt" gorm:"column:activated_at"`
	//ActivatedAt  sql.NullTime   `json:"activatedAt" gorm:"column:activated_at"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// equals: id, crt_at, upd_at, del_at
	//gorm.Model
}

// TableName 表名, 实现tabNm方法
// 此方法不支持动态表名, 如果使用动态表名, 使用Scope
// 约定大于配置 https://gorm.io/docs/conventions.html
func (UserInfo) TableName() string {
	return "user_info"
}

/*func (u *UserInfo) Scan(v interface{}) error {
	switch res := v.(type) {
	case UserInfo:
		res.MemberNumber.Valid = false
		res.MemberNumber.String = ""

		*u = res
	default:
		return errors.New("类型错误")
	}
	return nil
}*/
