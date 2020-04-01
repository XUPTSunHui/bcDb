package dbhandle

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/XUPTSunHui/bcDb/dbrequest"
)

//PreHandle define
func (dbi *Dbop) PreHandle(ctx context.Context, args dbrequest.PreHandle, reply *dbrequest.Reply) error {
	log.Printf("解析%v的json参数", args.FuncName)
	req := args.ReqParam.(string)
	d := json.NewDecoder(strings.NewReader(req))
	d.UseNumber()

	var err error

	switch args.FuncName {
	case "InsertBlockInfo":
		{
			target := dbrequest.BlockInfo{}
			err = d.Decode(&target)
			if err != nil {
				log.Fatalf("PreHandle解析json出错: %v", err)
			}

			err = dbi.InsertBlockInfo(ctx, target, reply)
		}
	default:
		{
			log.Printf("PreHandle中未发现 %v", args.FuncName)

		}
	}

	return err
}
