package export

// Conf export conf
type Conf struct {
	FileName string
	IPAddr   string
	Port     int
	Pwd      string
	DBName   string
	User     string
}

// IExport sql content export interface
type IExport interface {
	Do(sqlContent map[string]string, conf Conf) (err error)
}
