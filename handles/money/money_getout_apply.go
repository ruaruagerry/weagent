package money

import (
	"encoding/json"
	"time"
	"weagent/gconst"
	"weagent/pb"
	"weagent/server"
	"weagent/tables"

	"github.com/golang/protobuf/proto"
)

type getoutApplyReq struct {
	GetoutMoney int64 `json:"getoutmoney"`
}

func getoutApplyHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "money.getoutApplyHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &getoutApplyReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("getoutApplyHandle enter, req:", string(c.Body))

	// 查看提现金额是否合法
	_, has := getoutMoneyConfig[req.GetoutMoney]
	if !has {
		httpRsp.Result = proto.Int32(int32(gconst.ErrMoneyInvalidGetout))
		httpRsp.Msg = proto.String("提现金额错误")
		log.Errorf("code:%d msg:%s getoutmoney is invalid, getoutmoney:%d", httpRsp.GetResult(), httpRsp.GetMsg())
		return
	}

	db := c.DbConn
	playerid := c.UserID
	nowtime := time.Now()

	// 插入收益记录
	go func() {
		getoutrecord := &tables.Getoutrecord{
			ID:         playerid,
			Money:      req.GetoutMoney,
			CreateTime: nowtime,
			Status:     tables.GetoutStatusReview,
		}
		_, err := db.Insert(getoutrecord)
		if err != nil {
			httpRsp.Result = proto.Int32(int32(gconst.ErrDB))
			httpRsp.Msg = proto.String("提现记录插入失败")
			log.Errorf("code:%d msg:%s getoutrecord insert err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
			return
		}
	}()

	// rsp
	httpRsp.Result = proto.Int32(int32(gconst.Success))

	log.Info("getoutApplyHandle rsp, result:", httpRsp.GetResult())

	return
}
