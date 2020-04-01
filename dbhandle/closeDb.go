package dbhandle

import (
	"context"
	"log"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/pkg/errors"
)

func (dbi *Dbop) CloseDb(ctx context.Context, args *string, reply *dbrequest.Reply) error {
	if dbi == nil || dbi.Db == nil {
		return errors.New("mysql未打开")
	}

	err := dbi.Db.Close()
	if err != nil {
		log.Fatalf("Close失败: %v", err)
		reply.Status = -1
		return err
	}
	log.Println("db已正常关闭")

	reply.Status = 1
	return nil
}
