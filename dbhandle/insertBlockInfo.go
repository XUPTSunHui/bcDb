package dbhandle

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

type Dbop struct {
	Db       *sql.DB
	DataBase string
	UserName string
	Passwd   string
}

func (dbi *Dbop) InsertBlockInfo(ctx context.Context, args dbrequest.BlockInfo, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}

	stmt, err := dbi.Db.Prepare("INSERT INTO l_blk_data (blk_no, chan_name, blk_hash, pre_hash, data_hash, tx_cnt, tx_ids, db_time) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}

	res, err := stmt.Exec(args.BlockNumber, args.ChannelID, args.BlockHash, args.PreviousHash, args.DataHash, args.TxCount, args.TxIDs, time.Now())
	if err != nil {
		log.Fatalf("Exec失败: %v", err)
		reply.Status = -1
		return err
	}
	defer stmt.Close()

	rowCnt, _ := res.RowsAffected()
	reply.Status = 1
	reply.AffeRowCnt = int(rowCnt)

	log.Printf("编号%v区块已保存", args.BlockNumber)

	return err

}
