package fileread

// FileReadRegister fileRead register method
var FileReadRegister map[string]IFileRead

// register different file read method to FileReadRegister
func init() {
	FileReadRegister = map[string]IFileRead{}

	FileReadRegister["md"] = NewMarkDownFileRead()
}
