package export

var ExportRegister map[string]IExport

func init() {
	ExportRegister = map[string]IExport{}

	ExportRegister["file"] = NewFileExport()
}
