package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"

	"github.com/XUPTSunHui/bcDb/dbrequest"
	"github.com/smallnest/rpcx/client"
)

var (
	dbaddr   = flag.String("dbaddr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_db", "prefix path")
)

func main() {
	flag.Parse()

	//d := client.NewPeer2PeerDiscovery("tcp@"+"localhost:8972", "")
	// RegisterName 和 ServicePath对应
	ed := client.NewEtcdV3Discovery(*basePath, "Dbop", []string{*etcdAddr}, nil)
	// ed := client.NewEtcdDiscovery(*basePath, "Dbop", []string{*etcdAddr}, nil)
	// xclient := client.NewXClient("Dbop", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	xclient := client.NewXClient("Dbop", client.Failtry, client.RandomSelect, ed, client.DefaultOption)

	defer xclient.Close()

	ctx := context.Background()
	reply := &dbrequest.Reply{}

	err := xclient.Call(ctx, "InsertBlockInfo", insertBlockInfo1, reply)
	if err != nil {
		log.Fatalf("failed to call InsertBlockInfo : %v", err)
	}
	//若server端对reply进行了json序列化，此处reply需要进行json反序列化
	//若reply不需要进行json序列化，可以使用mapstructure包将map格式转为struct格式
	log.Printf("insert block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "InsertBlockInfo", insertBlockInfo2, reply)
	if err != nil {
		log.Fatalf("failed to call InsertBlockInfo2 : %v", err)
	}
	log.Printf("insert block_info2, response: %v\n", reply)

	err = xclient.Call(ctx, "QueryBlockInfo", queryBlockInfo, reply)
	if err != nil {
		log.Fatalf("failed to call QueryBlockInfo : %v", err)
	}
	log.Printf("query block_info, response: %v\n", reply)

	//测试PreHandle
	jsonBytes, _ := json.Marshal(insertBlockInfo1)
	preHandle1.ReqParam = string(jsonBytes)
	err = xclient.Call(ctx, "PreHandle", preHandle1, reply)
	if err != nil {
		log.Fatalf("failed to call PreHandle : %v", err)
	}
	log.Printf("PreHandle---insert block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "QueryBlockRow", queryBlockRow, reply)
	if err != nil {
		log.Fatalf("failed to call InsertBlockRow : %v", err)
	}
	log.Printf("query block_info row, response: %v\n", reply)

	err = xclient.Call(ctx, "UpdateBlockInfo", updateBlockInfo, reply)
	if err != nil {
		log.Fatalf("failed to call updateBlockInfo : %v", err)
	}
	log.Printf("update block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "FreeOp", freeOp1, reply)
	if err != nil {
		log.Fatalf("failed to call FreOp1 : %v", err)
	}
	log.Printf("FreeOp1 block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "DeleteBlockInfo", deleteBlockInfo1, reply)
	if err != nil {
		log.Fatalf("failed to call deleteBlockInfo : %v", err)
	}
	log.Printf("delete block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "DeleteBlockInfo", deleteBlockInfo2, reply)
	if err != nil {
		log.Fatalf("failed to call deleteBlockInfo : %v", err)
	}
	log.Printf("delete block_info, response: %v\n", reply)

	err = xclient.Call(ctx, "CloseDb", "", reply)
	if err != nil {
		log.Fatalf("failed to call CloseDb : %v", err)
	}
	log.Printf("CloseDb, response: %v\n", reply)

}
