package log

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func InitMlog()  {
	initDir()
	initMlog()
}

func initDir()  {
	var err error
	err = os.MkdirAll(LOGPATH, 0666)
	if err != nil {
		log.Fatalln("Failed to create log dir! Error: " + err.Error())
	}
	err = os.MkdirAll(MCONFIGPATH, 0666)
	if err != nil {
		log.Fatalln("Failed to create conf dir! Error: " + err.Error())
	}
}

func initMlog()  {
	hook := lumberjack.Logger{
		Filename: LOGFILE,
		MaxSize: MAXSIZE,
		MaxBackups: 31,
		MaxAge: 90,
		Compress: true,
		LocalTime: true,
	}

	var syncer zapcore.WriteSyncer

	syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))

	rawJson, err := ioutil.ReadFile(CONFIGFILE)
	if err != nil {
		panic(err)
	}

	var cfg zap.Config
	if err := json.Unmarshal(rawJson, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeTime = formatEncodingTime
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	level := cfg.Level.Level()
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), syncer, level)
	Mlog = zap.New(core)
}

func formatEncodingTime(t time.Time, enc zapcore.PrimitiveArrayEncoder)  {
	enc.AppendString(fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}