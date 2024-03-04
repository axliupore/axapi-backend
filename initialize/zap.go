package initialize

import (
	"fmt"
	"github.com/axliupore/axapi/axapi-backend/core"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitZap() {
	if ok := utils.DirExistOrNot(global.Config.Zap.Directory); !ok {
		fmt.Printf("create %v directory\n", global.Config.Zap.Directory)
		_ = os.Mkdir(global.Config.Zap.Directory, os.ModePerm)
	}

	z := &global.Config.Zap
	cores := core.Zap.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	// 是否显示行号
	if z.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	global.Log = logger
}
