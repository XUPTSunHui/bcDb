package dbhandle

import (
	"context"
	"database/sql"
	"log"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) QueryBlockInfo(ctx context.Context, args *dbrequest.QueryBlockInfo, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}
	var stmt *sql.Stmt
	var err error

	if args.Limit != 0 {
		stmt, err = dbi.Db.Prepare("SELECT * FROM l_blk_data ORDER BY blk_no desc LIMIT " + string(args.Limit))
		if err != nil {
			log.Fatalf("Prepare失败: %v", err)
			reply.Status = -1
			return err
		}
	} else {
		stmt, err = dbi.Db.Prepare("SELECT * FROM l_blk_data ORDER BY blk_no")
		if err != nil {
			log.Fatalf("Prepare失败: %v", err)
			reply.Status = -1
			return err
		}
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Query失败: %v", err)
		reply.Status = -1
		return err
	}
	defer rows.Close()

	var record dbrequest.BlockInfo
	for rows.Next() {
		err := rows.Scan(&record.BlockNumber, &record.ChannelID, &record.BlockHash, &record.PreviousHash, &record.DataHash, &record.TxCount, &record.TxIDs, &record.StorageTime)
		if err != nil {
			log.Fatalf("Scan失败: %v", err)
			reply.Status = -1
			return err
		}
		reply.Records = append(reply.Records, record)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Rows失败: %v", err)
		reply.Status = -1
		return err
	}
	reply.Status = 1

	log.Println("queryBlockInfo: 查询区块信息成功")

	return err
}
