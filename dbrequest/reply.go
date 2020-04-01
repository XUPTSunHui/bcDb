package dbrequest

type Reply struct {
	Status      int //-1失败 1成功
	AffeRowCnt int
	Records     []interface{}
	Row         interface{}
	JStrRecords string
}

type BlockInfo struct {
	BlockNumber  int
	ChannelID    string
	BlockHash    string
	PreviousHash string
	DataHash     string
	TxCount      int
	TxIDs        string
	StorageTime  string
}
