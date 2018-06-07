package server

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/ChamPly/dbtool/conversion"
	"github.com/ChamPly/dbtool/export"
	"github.com/ChamPly/dbtool/fileread"

	"github.com/ChamPly/dbtool/log"
)

var fileName *string
var outFile *string
var outType *string

func flagInit() {
	fileName = flag.String("i", "", "输入文件名")
	outFile = flag.String("C", "out.sql", "内容输出文件名")
	outType = flag.String("t", "mysql", "转换数据库类型")

	flag.Parse()
}

// Start server start
func Start() {
	flagInit()
	if strings.EqualFold(*fileName, "") {
		flag.Usage()
		return
	}

	// file to map
	method, ok := fileread.FileReadRegister[filepath.Ext(*fileName)[1:]]
	if !ok {
		log.Error("register file read not support")
		return
	}

	result, err := method.ReadFileToModel(*fileName)
	if err != nil {
		log.Error(err)
		return
	}

	// map to sql
	sqlHandle, ok := conversion.SQLConversionRegister[*outType]
	if !ok {
		log.Error("register conversion not support")
		return
	}

	sqlResult, err := sqlHandle.ModelToSQL(result)
	if err != nil {
		log.Error(err)
		return
	}

	// sql to file
	exportHandle, ok := export.ExportRegister["file"]
	if !ok {
		log.Errorf("register export not support")
		return
	}
	err = exportHandle.Do(sqlResult, export.Conf{
		FileName: *outFile,
	})
	if err != nil {
		log.Errorf("export result error:", err)
		return
	}

	log.Info("translation success!")
}
