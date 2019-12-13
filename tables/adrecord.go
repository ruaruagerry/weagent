package tables

import "time"

// Adrecord 广告收益记录
type Adrecord struct {
	Rid        int64     `xorm:"pk autoincr BIGINT(20) <-"`
	ID         string    `xorm:"id"`       // 用户ID
	Name       string    `xorm:"name"`     // 用户昵称
	Earnings   int64     `xorm:"earnings"` // 收益
	AdMoney    int64     `xorm:"admoney"`  // 当前余额
	CreateTime time.Time `xorm:"created"`  // 创建时间
}
