package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

const (
	marshallStructKey       = "marshallStruct"
	marshallStructFieldsKey = "marshallStructFields"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	return nil
}

// Config структура для yaml конфигурации
type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

// ConfigGenerate - подготовка данных и генерация конфига
func ConfigGenerate(tmpl string, outFilePath string) error {
	data := Config{
		Name:       "my-service",
		Port:       "8080",
		ReplicaSet: 3,
		ImageName:  "myapp",
		Tag:        "latest",
		EnvPath:    ".env",
	}

	return generate(tmpl, outFilePath, data)
}

// generate - функция генерации конфига
func generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("config").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, fields); err != nil {
		return fmt.Errorf("failed to fill template: %w", err)
	}

	if err := os.WriteFile(outfilePath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
