package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID                        int       `gorm:"primaryKey;autoIncrement" json:"id"`           // ID 是项目的唯一标识符，自增主键
	SaleStart                 time.Time `gorm:"type:timestamp" json:"saleStart"`              // SaleStart 是销售开始的时间
	SaleEnd                   time.Time `gorm:"type:timestamp" json:"saleEnd"`                // SaleEnd 是销售结束的时间
	RegistrationTimeEnds      time.Time `gorm:"type:timestamp" json:"registrationTimeEnds"`   // RegistrationTimeEnds 是注册结束时间
	RegistrationTimeStarts    time.Time `gorm:"type:timestamp" json:"registrationTimeStarts"` // RegistrationTimeStarts 是注册开始时间
	CreateTime                time.Time `gorm:"autoCreateTime" json:"createTime"`             // CreateTime 是记录创建时间，自动设置
	UpdateTime                time.Time `gorm:"autoUpdateTime" json:"updateTime"`             // UpdateTime 是记录更新时间，自动设置
	Tge                       time.Time `gorm:"type:timestamp" json:"tge"`                    // Tge 是指项目的某个重要时间点
	UnlockTime                time.Time `gorm:"type:timestamp" json:"unlockTime"`             // UnlockTime 是解锁时间，通常用于代币解锁
	VestingPortionsUnlockTime string    `gorm:"type:text" json:"vestingPortionsUnlockTime"`   // VestingPortionsUnlockTime 是分段解锁时间的 JSON 字符串
	VestingPercentPerPortion  string    `gorm:"type:text" json:"vestingPercentPerPortion"`    // VestingPercentPerPortion 是每个解锁部分的解锁百分比的 JSON 字符串
}

// Method to convert JSON strings to slices when retrieving data
func (p *Project) AfterFind(tx *gorm.DB) (err error) {
	var unlockTimes []int64
	var percentPerPortions []int64
	if err := json.Unmarshal([]byte(p.VestingPortionsUnlockTime), &unlockTimes); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(p.VestingPercentPerPortion), &percentPerPortions); err != nil {
		return err
	}
	return nil
}
