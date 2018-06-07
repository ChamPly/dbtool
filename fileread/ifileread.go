package fileread

// TableModel table model
type TableModel struct {
	Name        string
	Description string
	Fields      []FieldModel
}

// FieldModel filed property
type FieldModel struct {
	Name         string
	Type         string
	Length       int
	DefaultValue string
	IsEmpty      bool
	Constraints  string
	Description  string
}

// IFileRead read file interface
type IFileRead interface {
	ReadFileToModel(fileName string) (resultModel []TableModel, err error)
}
