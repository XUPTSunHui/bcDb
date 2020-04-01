package dbrequest

//PreHandle call server端的PreHandle函数时,先将参数进行json序列化
//server端接收时进行json反序列化，再手动调具体的函数
type PreHandle struct {
	FuncName string
	ReqParam interface{} //请求参数结构体
}
