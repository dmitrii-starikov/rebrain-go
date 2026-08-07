package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"text/template"
)

type FieldInfo struct {
	Name       string
	Tag        string
	IsExported bool
}

func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, inFilePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	var structFields []FieldInfo
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Name.Name != structName {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range structType.Fields.List {
			if len(field.Names) == 0 {
				continue
			}

			for _, name := range field.Names {
				if !name.IsExported() {
					continue
				}

				tag := ""
				if field.Tag != nil {
					tagValue := field.Tag.Value
					tagValue = tagValue[1 : len(tagValue)-1]

					if keyName := getTagValue(tagValue, "keyname"); keyName != "" {
						tag = keyName
					} else if envName := getTagValue(tagValue, "env"); envName != "" {
						tag = envName
					}
				}

				structFields = append(structFields, FieldInfo{
					Name:       name.Name,
					Tag:        tag,
					IsExported: name.IsExported(),
				})
			}
		}
		return true
	})

	if len(structFields) == 0 {
		return fmt.Errorf("struct %s not found or has no fields", structName)
	}

	templateData := struct {
		Package    string
		StructName string
		Fields     []FieldInfo
	}{
		Package:    node.Name.Name,
		StructName: structName,
		Fields:     structFields,
	}

	tmpl, err := template.New("marshaller").Parse(marshallerTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if err := os.WriteFile(outFilePath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func getTagValue(tagString, key string) string {
	tag := reflect.StructTag(tagString)
	return tag.Get(key)
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
