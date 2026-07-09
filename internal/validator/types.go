package validator

type MinerBid struct {
	IP     string  `bson:"ip"`
	UID    string  `bson:"uid"`
	Payout float64 `bson:"payout"`
	Count  int     `bson:"count"`
}

type MinerInfo struct {
	Core      *Core    `bson:"inline"`
	Block     int      `bson:"block"`
	Scores    []uint16 `bson:"scores,omitempty"`
	Timestamp int64    `bson:"timestamp,omitempty"`
	Weights   Weights  `bson:"weights,omitempty"`
}

type Weights struct {
	UIDs       []uint16  `bson:"uids"`
	Incentives []float64 `bson:"incentives"`
}
