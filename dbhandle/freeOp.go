package dbhandle

import (
	"context"
	"database/sql"
	"log"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) FreeOp(ctx context.Context, args *dbrequest.FreeOp, reply *dbrequest.Reply) error {
	if dbi == nil {
		return errors.New("mysql未打开")
	}

	var err error
	var stmt *sql.Stmt

	stmt, err = dbi.Db.Prepare(args.SQLCMD)
	if err != nil {
		log.Fatalf("Prepare失败: %v", err)
		reply.Status = -1
		return err
	}
	defer stmt.Close()

	if args.IsQuery == false {

		res, err := stmt.Exec()
		if err != nil {
			log.Fatalf("Prepare失败: %v", err)
			reply.Status = -1
			return err
		}
		rowCnt, _ := res.RowsAffected()
		reply.Status = 1
		reply.AffeRowCnt = int(rowCnt)

		log.Printf("FreeOp: %v已执行", args.SQLCMD)

	} else {
		rows, err := stmt.Query()
		if err != nil {
			log.Fatalf("Query失败: %v", err)
			reply.Status = -1
			return err
		}
		defer rows.Close()

		switch args.TableName {
		case "block_info":
			{
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
			}
		default:
			{
				log.Println("暂不支持FreeOp查询表格" + args.TableName)
			}
		}
		log.Printf("FreeOp: %v已查询", args.SQLCMD)

	}

	return err
}
