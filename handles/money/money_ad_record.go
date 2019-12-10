package money

import (
	"encoding/json"
	"weagent/gconst"
	"weagent/pb"
	"weagent/server"
	"weagent/tables"

	"github.com/golang/protobuf/proto"
)

type adRecordReq struct {
	Start int32 `json:"start"`
	End   int32 `json:"end"`
}

type adRecordRsp struct {
	AdRecords []*tables.Adrecord `json:"adrecords"`
}

func adRecordHandle(c *server.StupidContext) {
	log := c.Log.WithField("func", "money.adRecordHandle")

	httpRsp := pb.HTTPResponse{
		Result: proto.Int32(int32(gconst.UnknownError)),
	}
	defer c.WriteJSONRsp(&httpRsp)

	// req
	req := &adRecordReq{}
	if err := json.Unmarshal(c.Body, req); err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("请求信息解析失败")
		log.Errorf("code:%d msg:%s json Unmarshal err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	log.Info("getoutRecordHandle enter, req:", string(c.Body))

	start := int(req.Start)
	end := int(req.End)
	if start >= end {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParam))
		httpRsp.Msg = proto.String("请求参数错误")
		log.Errorf("code:%d msg:%s req param err, start:%d end:%d", httpRsp.GetResult(), httpRsp.GetMsg(), start, end)
		return
	}

	db := c.DbConn
	playerid := c.UserID

	adrecords := []*tables.Adrecord{}
	err := db.Where("playerid = ?", playerid).Limit(end, start).Find(&adrecords)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrDB))
		httpRsp.Msg = proto.String("查询广告收益记录失败")
		log.Errorf("code:%d msg:%s get adrecords err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}

	// rsp
	rsp := &adRecordRsp{
		AdRecords: adrecords,
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		httpRsp.Result = proto.Int32(int32(gconst.ErrParse))
		httpRsp.Msg = proto.String("返回信息marshal解析失败")
		log.Errorf("code:%d msg:%s json marshal err, err:%s", httpRsp.GetResult(), httpRsp.GetMsg(), err.Error())
		return
	}
	httpRsp.Result = proto.Int32(int32(gconst.Success))
	httpRsp.Data = data

	log.Info("getoutRecordHandle rsp, rsp:", string(data))

	return
}
