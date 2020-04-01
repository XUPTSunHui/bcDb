package dbhandle

import (
	"context"
	"log"
	"strconv"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) DeleteBlockInfo(ctx context.Context, args *dbrequest.DeleteBlockInfo, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}

	stmt, err := dbi.Db.Prepare("DELETE FROM block_info where block_number = ?")
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}

	res, err := stmt.Exec(strconv.Itoa(args.DelBlkNum))
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}
	defer stmt.Close()

	rowCnt, _ := res.RowsAffected()
	reply.Status = 1
	reply.AffeRowCnt = int(rowCnt)

	log.Printf("编号%v区块记录已删除", args.DelBlkNum)

	return err

}
