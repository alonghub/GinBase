// log 日志模块
package log

import (
    "fmt"
    "github.com/alonghub/GinBase/utils/config"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "strings"
)

var Logger *zap.Logger

// ZapLoggerInit 使用uber的ZAP日志模块初始化
func ZapLoggerInit(mode, serviceName string) {
    logDir := strings.TrimRight(config.ServerSetting.LogDir, "/")
    if logDir == "" {
        logDir = "."
    }
    hook := lumberjack.Logger{
        Filename:   fmt.Sprintf("%s/logs/td.log", logDir), // 日志文件路径
        MaxSize:    1,                                     // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: 10,                                    // 日志文件最多保存多少个备份
        MaxAge:     7,                                     // 文件最多保存多少天
        Compress:   false,                                 // 是否压缩
    }

    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "linenum",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        //EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeCaller: zapcore.ShortCallerEncoder, // 全路径编码器
        EncodeName:   zapcore.FullNameEncoder,
    }

    atomicLevel := zap.NewAtomicLevel()
    if mode == "debug" {
        atomicLevel.SetLevel(zap.DebugLevel)
    } else {
        atomicLevel.SetLevel(zap.InfoLevel)
    }

    // 设置日志级别

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
        //zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件
        atomicLevel,                                         // 日志级别
    )

    // 开启开发模式，堆栈跟踪
    caller := zap.AddCaller()
    // 开启文件及行号
    envLevel := zap.Development()
    // 设置初始化字段
    filed := zap.Fields(zap.String("serviceName", serviceName))
    // 构造日志
    if mode == "debug" {
        Logger = zap.New(core, caller, envLevel, filed)
    } else {
        Logger = zap.New(core, filed)
    }
    // log.Logger.Info("reboot_info", zap.Strings("a list", []string{}), zap.Bool("is_pm", bool))
    Logger.Info("log init done.")
}
