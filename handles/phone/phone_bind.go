package phone

import (
	"encoding/json"
	"weagent/gconst"
	"weagent/pb"
	"weagent/rconst"
	"weagent/server"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
)

type bindReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func bindHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "phone.bindHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &bindReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("bindHandle enter, req:", string(c.Body))

	conn := c.RedisConn
	playerid := c.UserID

	// redis multi get
	conn.Send("MULTI")
	conn.Send("GET", rconst.StringPhoneCodePrefix+playerid)
	conn.Send("HGET", rconst.HashAccountPrefix+playerid, rconst.FieldAccPhone)
	conn.Send("HGET", rconst.HashAccountPrefix+playerid, rconst.FieldAccOpenID)
	conn.Send("SISMEMBER", rconst.SetPhoneHasBinded, req.Phone)
	redisMDArray, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一获取缓存操作失败")
		log.Errorf("code:%d msg:%s redisMDArray Values err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// 验证码
	msgbyte, _ := redis.Bytes(redisMDArray[0], nil)
	myphone, _ := redis.String(redisMDArray[1], nil)
	openid, _ := redis.String(redisMDArray[2], nil)
	isinhasbinded, _ := redis.Bool(redisMDArray[3], nil)

	// do something
	if isinhasbinded {
		httpRsp.Result = proto.Int32(int32(gconst.ErrPhoneHasBinded))
		httpRsp.Msg = proto.String("手机号码已被绑定")
		log.Errorf("code:%d msg:%s, phoneno has binded", httpRsp.GetResult(), httpRsp.GetMsg())
		return
	}

	if len(msgbyte) == 0 {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("您输入的验证码错误")
		log.Errorf("code:%d msg:%s, redis bytes err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	if myphone != "" {
		httpRsp.Result = proto.Int32(int32(gconst.ErrPhoneAlreadyBind))
		httpRsp.Msg = proto.String("已绑定手机号码")
		log.Errorf("code:%d msg:%s, already bind phoneno, phonno:%s", httpRsp.GetResult(), httpRsp.GetMsg(), myphone)
		return
	}

	msg := &rconst.PhoneMsg{}
	err = json.Unmarshal(msgbyte, msg)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("验证码信息unmarshal解析失败")
		log.Errorf("code:%d msg:%s, msgbyte unmarshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// 验证输入信息
	log.Info("msg:", msg)
	if msg.Phone != req.Phone || msg.Code != req.Code {
		httpRsp.Result = proto.Int32(int32(gconst.ErrPhoneCode))
		httpRsp.Msg = proto.String("验证码错误")
		log.Errorf("code:%d msg:%s, phoneno:%s code:%s", httpRsp.GetResult(), httpRsp.GetMsg(), msg.Phone, msg.Code)
		return
	}

	// redis multi set
	conn.Send("MULTI")
	if openid != "" {
		conn.Send("HSET", rconst.HashAccountOpenIDPrefix+openid, rconst.FieldAccOpenIDPhone, req.Phone)
	}
	conn.Send("HSET", rconst.HashAccountPrefix+playerid, rconst.FieldAccPhone, req.Phone)
	conn.Send("SADD", rconst.SetPhoneHasBinded, req.Phone)
	conn.Send("DEL", rconst.StringPhoneCodePrefix+playerid)
	_, err = conn.Do("EXEC")
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrRedis))
		httpRsp.Msg = proto.String("统一存储缓存操作失败")
		log.Errorf("code:%d msg:%s exec err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("bindHandle rsp, result:", httpRsp.GetResult())

	return
}
