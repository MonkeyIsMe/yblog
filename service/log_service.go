package service

import (
	"log"
	"yblog/constant"
	"yblog/model"

	"github.com/MonkeyIsMe/MyTool/time"
)

func AddLog(old, new []byte, op, obj int) error {
	logDetal := constant.LogOpDetail{
		OldInfo: string(old),
		NewInfo: string(new),
	}
	logInfo := model.Log{
		LogOperation: op,
		LogObj:       obj,
		LogFlag:      1,
		ErrorMsg:     "Success",
		LogOperator:  "admin", // todo 填入账号
		LogDetail:    logDetal,
		LogTime:      time.QueryNowTime(),
	}
	err := logInfo.AddLog()
	if err != nil {
		log.Fatalf("LogService Add Log err : [%+v]", err)
		return err
	}
	return nil
}
