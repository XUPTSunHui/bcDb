package main

import (
	"github.com/XUPTSunHui/bcDb/dbrequest"
)

var (
	insertBlockInfo1 *dbrequest.BlockInfo = &dbrequest.BlockInfo{
		BlockNumber:  4,
		ChannelID:    "channel1",
		BlockHash:    "2hd836fhwi",
		PreviousHash: "638hf2389rfn",
		DataHash:     "08h2s0m",
		TxCount:      2,
		TxIDs:        "tx1wdefg" + "," + "tx2kdjirj",
		//StorageTime: 存的时候自动给出
	}

	insertBlockInfo2 *dbrequest.BlockInfo = &dbrequest.BlockInfo{
		BlockNumber:  5,
		ChannelID:    "channel1",
		BlockHash:    "7hd636fhwi",
		PreviousHash: "6siu8hf2389rfn",
		DataHash:     "8yghh2s0m",
		TxCount:      2,
		TxIDs:        "tx14rfefg" + "," + "tx298jirj",
		//StorageTime: 存的时候自动给出
	}

	queryBlockInfo *dbrequest.QueryBlockInfo = &dbrequest.QueryBlockInfo{
		// Limit: 20,
	}

	queryBlockRow *dbrequest.QueryBlockRow = &dbrequest.QueryBlockRow{
		// QBlkNum: 4,
	}

	updateBlockInfo *dbrequest.UpdateBlockInfo = &dbrequest.UpdateBlockInfo{
		UpBlkNum: 5,
		UpItem:   "channel_id",
		UpValue:  "channel2",
	}

	freeOp1 *dbrequest.FreeOp = &dbrequest.FreeOp{
		SQLCMD:    "select * from block_info",
		IsQuery:   true,
		TableName: "block_info",
	}

	deleteBlockInfo1 *dbrequest.DeleteBlockInfo = &dbrequest.DeleteBlockInfo{
		DelBlkNum: 4,
	}

	deleteBlockInfo2 *dbrequest.DeleteBlockInfo = &dbrequest.DeleteBlockInfo{
		DelBlkNum: 5,
	}

	preHandle1 *dbrequest.PreHandle = &dbrequest.PreHandle{
		FuncName: "InsertBlockInfo",
		//ReqParam: interface{}, json序列化之后再赋值
	}
)
