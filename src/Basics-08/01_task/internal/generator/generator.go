package generator

const (
	marshallStructKey       = "marshallStruct"
	marshallStructFieldsKey = "marshallStructFields"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	return nil
}

// Task03 - структура для yaml конфигурации
// Если нужно расширить yaml файл, тогда эту структуру нужно также расширить
// нужными параметрами, yaml файл и эта структура должны быть идентичными.
type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

// Task03 - функция генерации конфига
func ConfigGenerate(tmpl string, outFilePath string) error {
	return nil
}

// Основная функция которая должна производить генерацию
func generate(tmpl string, outfilePath string, fields interface{}) error {
	return nil
}
