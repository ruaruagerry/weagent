package rconst

const (
	// LockTime 锁的时限
	LockTime = 10

	/* ---------- setup ---------- */

	// StringLockRealModifyHandlePrefix 实名认证锁
	StringLockRealModifyHandlePrefix = "weagent:lock:realmodifyhandle:"

	/* ---------- money ---------- */

	// StringLockMoneyAdSeePrefix 广告上报锁
	StringLockMoneyAdSeePrefix = "weagent:lock:moneyadsee:"
	// StringLockMoneyAdClickPrefix 广告点击锁
	StringLockMoneyAdClickPrefix = "weagent:lock:moneyadclick:"
	// StringLockMoneyGetoutApplyPrefix 申请提现锁
	StringLockMoneyGetoutApplyPrefix = "weagent:lock:moneygetoutapply:"
	// StringLockMoneyGetoutResultPrefix 提现审核锁
	StringLockMoneyGetoutResultPrefix = "weagent:lock:moneygetoutresult:"
)
