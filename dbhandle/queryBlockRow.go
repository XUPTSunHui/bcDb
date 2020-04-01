package dbhandle

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) QueryBlockRow(ctx context.Context, args *dbrequest.QueryBlockRow, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}
	var stmt *sql.Stmt
	var err error
	if args.QBlkNum != 0 {
		stmt, err = dbi.Db.Prepare("SELECT * FROM block_info where block_number = " + strconv.Itoa(args.QBlkNum))
		if err != nil {
			log.Fatalf("Prepare失败: %v", err)
			reply.Status = -1
			return err
		}
	} else {
		stmt, err = dbi.Db.Prepare("SELECT * FROM block_info where block_number = (SELECT max(block_number) FROM block_info)")
		if err != nil {
			log.Fatalf("Prepare失败: %v", err)
			reply.Status = -1
			return err
		}
	}
	defer stmt.Close()

	var record dbrequest.BlockInfo
	err = stmt.QueryRow().Scan(&record.BlockNumber, &record.ChannelID, &record.BlockHash, &record.PreviousHash, &record.DataHash, &record.TxCount, &record.TxIDs, &record.StorageTime)
	if err != nil {
		log.Fatalf("Scan失败: %v", err)
		reply.Status = -1
		return err
	}
	reply.Row = record
	reply.Status = 1

	log.Println("queryBlockRow: 查询区块信息成功")

	return err

}
