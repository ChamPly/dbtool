package conversion

import (
	"fmt"
	"strings"

	"github.com/champly/dbtool/fileread"
)

// MySQLConversion mysql conversion
type MySQLConversion struct {
}

// NewMySQLConversion new mysql conversion object
func NewMySQLConversion() *MySQLConversion {
	return &MySQLConversion{}
}

// ModelToSQL convertion model to mysql sql
func (m *MySQLConversion) ModelToSQL(tables []fileread.TableModel) (result map[string]string, err error) {
	result = map[string]string{}

	for _, table := range tables {
		result[table.Name] = m.buildSQL(table)
	}

	return
}

func (m *MySQLConversion) buildSQL(table fileread.TableModel) (sql string) {
	// -- ----------------------------
	// -- Table structure for kc_delivery_channel
	// -- ----------------------------
	// DROP TABLE IF EXISTS `kc_delivery_channel`;
	// CREATE TABLE `kc_delivery_channel` (

	//   PRIMARY KEY (`channel_id`)
	// ) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='发货渠道';

	tableDescSQLStr := fmt.Sprintf(`
-- ----------------------------
-- Table structure for %s
-- ----------------------------`, table.Name)
	dropTableSQLStr := fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", table.Name)
	createTablePreStr := fmt.Sprintf("CREATE TABLE `%s` (", table.Name)

	pkFields := []string{}
	fieldSqls := []string{}
	for _, field := range table.Fields {
		if strings.EqualFold(strings.ToUpper(field.Constraints), "PK") {
			pkFields = append(pkFields, fmt.Sprintf("`%s`", field.Name))
		}

		fieldSqls = append(fieldSqls, m.buildFildSQL(field))
	}

	pkSQLStr := ""
	if len(pkFields) > 0 {
		pkSQLStr = fmt.Sprintf("\r\n\tPRIMARY KEY (%s)", strings.Join(pkFields, ","))
	}
	fieldsSQLStr := strings.Join(fieldSqls, "\r\n\t")

	endTableStr := fmt.Sprintf(") ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='%s';", table.Description)

	sql = fmt.Sprintf("%s\r\n%s\r\n%s\r\n\t%s%s\r\n%s", tableDescSQLStr, dropTableSQLStr, createTablePreStr, fieldsSQLStr, pkSQLStr, endTableStr)

	return
}

func (m *MySQLConversion) buildFildSQL(field fileread.FieldModel) (filedSQL string) {
	// `notify_url` varchar(256) CHARACTER SET utf8 DEFAULT NULL COMMENT '通知地址'
	fileName := field.Name

	typeStr := field.Type
	if strings.EqualFold(field.Type, "number") {
		typeStr = "int"
	}

	if field.Length > 0 {
		typeStr += fmt.Sprintf("(%d)", field.Length)
	}
	if strings.EqualFold(field.Type, "varchar") {
		typeStr += " CHARACTER SET utf8"
	}

	isNULL := ""
	if !field.IsEmpty {
		isNULL = "NOT NULL "
	}
	if strings.EqualFold(strings.ToUpper(field.Constraints), "PK") && (field.Type == "int" || field.Type == "number") {
		isNULL += " AUTO_INCREMENT"
	}

	defaultValue := "DEFAULT"
	if strings.EqualFold(field.DefaultValue, "") {
		defaultValue += " NULL"
	} else {
		defaultValue += fmt.Sprintf(" '%s'", field.DefaultValue)
	}
	if !strings.EqualFold(field.Type, "varchar") {
		defaultValue = ""
	}

	desc := ""
	if !strings.EqualFold(field.Description, "") {
		desc = fmt.Sprintf("COMMENT '%s'", field.Description)
	}

	filedSQL = fmt.Sprintf("`%s` %s %s%s %s,", fileName, typeStr, isNULL, defaultValue, desc)
	return
}
