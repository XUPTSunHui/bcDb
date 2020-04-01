package dbhandle

import (
	"context"
	"log"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) UpdateBlockInfo(ctx context.Context, args *dbrequest.UpdateBlockInfo, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}

	stmt, err := dbi.Db.Prepare("UPDATE block_info set " + string(args.UpItem) + "= ? where block_number = ?")
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}
	res, err := stmt.Exec(args.UpValue, args.UpBlkNum)
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}
	defer stmt.Close()

	rowCnt, _ := res.RowsAffected()
	reply.Status = 1
	reply.AffeRowCnt = int(rowCnt)

	log.Printf("编号%v区块信息已更新", args.UpBlkNum)

	return err

}
