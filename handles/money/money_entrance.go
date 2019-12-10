package money

import (
	"weagent/gconst"
	"weagent/pb"
	"weagent/server"

	"github.com/golang/protobuf/proto"
)

type entranceRsp struct {
	Total       int64 `json:"total"`       // 总收益
	Money       int64 `json:"money"`       // 当前账户余额
	GetoutTotal int64 `json:"getouttotal"` // 总提现金额
}

func entranceHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "money.entranceHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	// req := &pb.HelloReq{}
	// if err := json.Unmarshal(c.Body, req); err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
	// 	httpRsp.Msg = proto.String("请求信息解析失败")
	// 	log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
	// 	return
	// }

	// log.Info("entranceHandle enter, req:", string(c.Body))

	// conn := c.RedisConn
	// playerid := c.UserID

	// // redis multi get
	// conn.Send("MULTI")
	// redisMDArray, err := redis.Values(conn.Do("EXEC"))
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
	// 	httpRsp.Msg = proto.String("统一获取缓存操作失败")
	// 	log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
	// 	return
	// }

	// // do something

	// // redis multi set
	// conn.Send("MULTI")
	// _, err = conn.Do("EXEC")
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
	// 	httpRsp.Msg = proto.String("统一存储缓存操作失败")
	// 	log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
	// 	return
	// }

	// // rsp
	// rsp := &pb.HelloRsp{}
	// data, err := json.Marshal(rsp)
	// if err != nil {
	// 	httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
	// 	httpRsp.Msg = proto.String("返回信息marshal解析失败")
	// 	log.Errorf("code:%d msg:%s json marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
	// 	return
	// }
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	// httpRsp.Data = data

	// log.Info("entranceHandle rsp, rsp:", string(data))

	return
}
