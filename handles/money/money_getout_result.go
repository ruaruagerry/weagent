package money

import (
	"encoding/json"
	"weagent/gconst"
	"weagent/pb"
	"weagent/rconst"
	"weagent/server"
	"weagent/tables"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type getoutResultReq struct {
	Rid    int64 `json:"rid"`
	Status int32 `json:"status"`
}

func getoutResultHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "money.getoutResultHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &getoutResultReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("getoutResultHandle enter, req:", string(c.Body))

	db := c.DbConn
	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("HGET", rconst.HashMoneyPrefix+playerid, rconst.FieldMoneyNum)
	conn.Send("HGET", rconst.HashMoneyPrefix+playerid, rconst.FieldMoneyGetout)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	moneynum, _ := redis.Int64(redisMDArray[0], nil)
	getout, _ := redis.Int64(redisMDArray[1], nil)

	// do something
	getoutrecord := tables.Getoutrecord{}
	has, err := db.Where("rid = ?", req.Rid).Get(&getoutrecord)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrDB))
		httpRsp.Msg = proto.String("查询提现记录失败")
		log.Errorf("code:%d msg:%s query getoutrecord err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	if !has {
		httpRsp.Result = proto.Int32(int32(gconst.ErrDB))
		httpRsp.Msg = proto.String("未找到提现记录")
		log.Errorf("code:%d msg:%s not find rid getoutrecord, rid:%s", httpRsp.GetResult(), httpRsp.GetMsg(), req.Rid)
		return
	}

	// 提现成功
	if req.Status == tables.GetoutStatusSuccess {
		moneynum -= getoutrecord.Money
		getout += getoutrecord.Money

		// redis multi set
		conn.Send("MULTI")
		conn.Send("HSET", rconst.HashMoneyPrefix+playerid, rconst.FieldMoneyNum, moneynum)
		conn.Send("HSET", rconst.HashMoneyPrefix+playerid, rconst.FieldMoneyGetout, getout)
		_, err = conn.Do("EXEC")
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
			httpRsp.Msg = proto.String("统一存储缓存操作失败")
			log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}

	// 更新
	getoutrecord.Status = req.Status
	_, err = db.Where("rid = ?", req.Rid).AllCols().Update(&getoutrecord)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrDB))
		httpRsp.Msg = proto.String("更新提现记录失败")
		log.Errorf("code:%d msg:%s update getoutrecord err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("getoutResultHandle rsp, result:", httpRsp.GetResult())

	return
}
