package convertor

import (
	"fmt"
	"reflect"
)

// MapToStruct fills a struct from a map using reflection.
// It uses the "keyname" struct tag to match map keys.
func MapToStruct(mp map[string]interface{}, item interface{}) error {
	if mp == nil {
		return nil
	}

	val := reflect.ValueOf(item)

	// Must be a pointer to a struct
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("item must be a pointer to struct, got %s", val.Kind())
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("item must be a pointer to struct, got pointer to %s", val.Kind())
	}

	valType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := valType.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		key := fieldType.Tag.Get("keyname")
		if key == "" {
			key = fieldType.Name
		}

		mapValue, ok := mp[key]
		if !ok {
			continue
		}

		if mapValue == nil {
			continue
		}

		// nested structs
		if field.Kind() == reflect.Struct {
			// check if the value is already a struct of the right type
			if reflect.TypeOf(mapValue) == field.Type() {
				field.Set(reflect.ValueOf(mapValue))
			} else if mapVal, ok := mapValue.(map[string]interface{}); ok {
				// map -> recursively fill
				newStruct := reflect.New(field.Type())
				if err := MapToStruct(mapVal, newStruct.Interface()); err != nil {
					return fmt.Errorf("field %s: %w", fieldType.Name, err)
				}
				field.Set(newStruct.Elem())
			} else {
				return fmt.Errorf("field %s: expected struct or map, got %T", fieldType.Name, mapValue)
			}
		} else if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			if reflect.TypeOf(mapValue) == field.Type().Elem() {
				ptr := reflect.New(field.Type().Elem())
				ptr.Elem().Set(reflect.ValueOf(mapValue))
				field.Set(ptr)
			} else if mapVal, ok := mapValue.(map[string]interface{}); ok {
				newStruct := reflect.New(field.Type().Elem())
				if err := MapToStruct(mapVal, newStruct.Interface()); err != nil {
					return fmt.Errorf("field %s: %w", fieldType.Name, err)
				}
				field.Set(newStruct)
			} else {
				return fmt.Errorf("field %s: expected struct or map, got %T", fieldType.Name, mapValue)
			}
		} else {
			// Simple type assignment
			if err := setField(field, mapValue); err != nil {
				return fmt.Errorf("field %s: %w", fieldType.Name, err)
			}
		}
	}

	return nil
}

// StructToMap converts a struct to a map[string]interface{} using reflection.
// It uses the "keyname" struct tag to determine map keys.
func StructToMap(item interface{}) map[string]interface{} {
	if item == nil {
		return nil
	}

	val := reflect.ValueOf(item)

	// Pointer -> get the underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Not a struct -> return nil
	if val.Kind() != reflect.Struct {
		return nil
	}

	valType := val.Type()
	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := valType.Field(i)

		// Handle only public fields
		if !fieldType.IsExported() {
			continue
		}

		// Get the key name from tag or use field name
		key := fieldType.Tag.Get("keyname")
		if key == "" {
			key = fieldType.Name
		}

		// Handle nested structs - save structs, don't convert to map
		if field.Kind() == reflect.Struct {
			result[key] = field.Interface()
		} else if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			if !field.IsNil() {
				result[key] = field.Interface()
			} else {
				result[key] = nil
			}
		} else {
			result[key] = field.Interface()
		}
	}

	return result
}

// setField sets a field value with type conversion
func setField(field reflect.Value, value interface{}) error {
	if !field.CanSet() {
		return fmt.Errorf("field can`t be set")
	}

	val := reflect.ValueOf(value)

	// value is nil, set zero value
	if !val.IsValid() {
		field.Set(reflect.Zero(field.Type()))
		return nil
	}

	// if types match directly
	if val.Type().AssignableTo(field.Type()) {
		field.Set(val)
		return nil
	}

	// Try to convert
	switch field.Kind() {
	case reflect.String:
		field.SetString(fmt.Sprintf("%v", value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch v := value.(type) {
		case int:
			field.SetInt(int64(v))
		case int64:
			field.SetInt(v)
		case float64:
			field.SetInt(int64(v))
		default:
			return fmt.Errorf("cannot convert %T to int", value)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch v := value.(type) {
		case int:
			field.SetUint(uint64(v))
		case int64:
			field.SetUint(uint64(v))
		case float64:
			field.SetUint(uint64(v))
		default:
			return fmt.Errorf("cannot convert %T to uint", value)
		}
	case reflect.Float32, reflect.Float64:
		switch v := value.(type) {
		case float64:
			field.SetFloat(v)
		case int:
			field.SetFloat(float64(v))
		case int64:
			field.SetFloat(float64(v))
		default:
			return fmt.Errorf("cannot convert %T to float", value)
		}
	case reflect.Bool:
		switch v := value.(type) {
		case bool:
			field.SetBool(v)
		default:
			return fmt.Errorf("cannot convert %T to bool", value)
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
