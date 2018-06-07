package conversion

import (
	"github.com/ChamPly/dbtool/fileread"
)

// IConversion 转换
type IConversion interface {
	ModelToSQL(tables []fileread.TableModel) (result map[string]string, err error)
}
