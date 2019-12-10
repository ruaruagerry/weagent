package tables

import "time"

const (
	// GetoutStatusReview 审核中
	GetoutStatusReview = 0
	// GetoutStatusPass 审核通过
	GetoutStatusPass = 1
	// GetoutStatusRefused 审核拒绝
	GetoutStatusRefused = 2
	// GetoutStatusSuccess 提现成功
	GetoutStatusSuccess = 3
	// GetoutStatusFailed 提现失败
	GetoutStatusFailed = 4
)

// Getoutrecord 提现记录
type Getoutrecord struct {
	Rid        int64     `xorm:"pk autoincr BIGINT(20) <-"`
	ID         string    `xorm:"id"`       // 用户ID
	Earnings   int64     `xorm:"earnings"` // 收益
	CreateTime time.Time `xorm:"created"`  // 创建时间
	Status     int32     `xorm:"status"`   // 提现状态
}
