package fileread

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/champly/dbtool/utility"
)

// MarkDownFile read database structure from markdown file
type MarkDownFile struct {
	tablePrefix    string
	tableMatchReg1 string
	tableMatchReg2 string
	fieldMatchReg  string
	escapeStr      []string
}

// NewMarkDownFileRead new MarkDownFile object
func NewMarkDownFileRead() *MarkDownFile {
	return &MarkDownFile{
		tablePrefix:    "####",
		tableMatchReg1: "^####.*?\\d+.*?([\u4e00-\u9fa5]+).*?\\[(.*?)\\]$",
		tableMatchReg2: "^####.*?\\d+.*?([\u4e00-\u9fa5]+).*?\\[(.*?)\\].*?(\\(.*?\\))$",
		fieldMatchReg:  `^\|(.*?)\|(.*?)([\(\d+\)]*)\|(.*?)\|(.*?)\|(.*?)\|(.*?)\|$`,
		escapeStr: []string{
			"|---|---|---|---|---|---|",
			"|字段名",
		},
	}
}

// ReadFileToModel read file content and build database structure
func (md *MarkDownFile) ReadFileToModel(fileName string) (resultModel []TableModel, err error) {
	content, err := md.read(fileName)
	if err != nil {
		return
	}

	resultModel, err = md.analysisContent(content)
	if err != nil {
		return
	}
	return
}

func (md *MarkDownFile) read(fileName string) (content chan string, err error) {
	_, err = os.Stat(fileName)
	if err != nil {
		err = fmt.Errorf("读取文件:%s, 异常:%s", fileName, err.Error())
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		err = fmt.Errorf("读取文件:%s, 失败:%s", fileName, err.Error())
		return
	}

	content = make(chan string, 10)
	reader := bufio.NewReader(file)
	go func() {
		for {
			contentBytes, _, e := reader.ReadLine()
			if e == io.EOF {
				file.Close()
				close(content)
				break
			}
			content <- string(contentBytes)
		}
	}()

	return
}

func (md *MarkDownFile) analysisContent(content chan string) (resultModel []TableModel, err error) {
	resultModel = []TableModel{}
	tableModel := TableModel{}

	isNew := true
	for str := range content {
		if md.escapeMatch(str) {
			continue
		}

		if strings.HasPrefix(str, md.tablePrefix) {
			if !isNew {
				resultModel = append(resultModel, tableModel)
			}
			tableModel, err = md.buildTableModel(str)
			isNew = false
		}

		if b, e := regexp.MatchString(md.fieldMatchReg, str); b && e == nil {
			f, err := md.buildFieldModel(str)
			if err != nil {
				return nil, err
			}
			tableModel.Fields = append(tableModel.Fields, f)
		}
	}

	resultModel = append(resultModel, tableModel)
	return
}

func (md *MarkDownFile) buildTableModel(tableStr string) (tableModel TableModel, err error) {
	// #### 4. 服务信息 [epg_sys_service]
	reg := regexp.MustCompile(md.tableMatchReg1)
	groupsMatch := reg.FindStringSubmatch(tableStr)

	if len(groupsMatch) == 3 {
		tableModel = TableModel{
			Name:        groupsMatch[2],
			Description: groupsMatch[1],
			Fields:      []FieldModel{},
		}
		return
	}

	reg = regexp.MustCompile(md.tableMatchReg2)
	groupsMatch = reg.FindStringSubmatch(tableStr)
	if len(groupsMatch) != 4 {
		err = fmt.Errorf("匹配错误: raw:%s, 不是标准的匹配格式:#### 1. 表描述 [表名] (表补充)", tableStr)
		return
	}

	tableModel = TableModel{
		Name:        groupsMatch[2],
		Description: fmt.Sprintf("%s%s", groupsMatch[1], groupsMatch[3]),
		Fields:      []FieldModel{},
	}
	return
}

func (md *MarkDownFile) buildFieldModel(fildStr string) (filedModel FieldModel, err error) {
	// |id|number(20)|-|否|PK|编号|
	reg := regexp.MustCompile(md.fieldMatchReg)
	groupsMatch := reg.FindStringSubmatch(fildStr)

	if len(groupsMatch) != 8 {
		fmt.Println(groupsMatch)
		err = fmt.Errorf("匹配错误: raw:%s, 不是标准的匹配格式:|字段名|类型(长度)|默认值|为空|约束|描述|", fildStr)
		return
	}

	groupsMatch[3] = strings.Trim(groupsMatch[3], "(")
	groupsMatch[3] = strings.Trim(groupsMatch[3], ")")
	length, err := utility.StringToInt(groupsMatch[3])
	if err != nil {
		err = fmt.Errorf("%s:类型长度不合法:%+v", fildStr, err)
		return
	}
	isEmpty := false
	if strings.EqualFold(groupsMatch[5], "是") {
		isEmpty = true
	}

	filedModel = FieldModel{
		Name:         groupsMatch[1],
		Type:         groupsMatch[2],
		Length:       length,
		DefaultValue: groupsMatch[4],
		IsEmpty:      isEmpty,
		Constraints:  groupsMatch[6],
		Description:  groupsMatch[7],
	}
	return
}

func (md *MarkDownFile) escapeMatch(str string) bool {
	for _, espStr := range md.escapeStr {
		if strings.HasPrefix(str, espStr) {
			return true
		}
	}
	return false
}
