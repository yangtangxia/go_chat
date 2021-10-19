package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Bind(req *http.Request, obj interface{}) error {
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		return BindJson(req, obj)
	}

	if strings.Contains(strings.ToLower(contentType), "application/x-www-form-urlencoded") {
		return BindForm(req, obj)
	}

	return errors.New("当前方法暂不支持")
}

func BindJson(req *http.Request, obj interface{}) error {
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(s,obj)
	return err
}

func BindForm(req *http.Request, ptr interface{}) error {
	req.ParseForm()
	fmt.Println(req.Form.Encode())
	err := mapForm(ptr, req.Form)

	return err
}

func mapForm(ptr interface{}, form map[string][]string) error {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()

	for i := 0;i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}

		structfieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get("form")
		if inputFieldName == "" {
			inputFieldName = typeField.Name

			if structfieldKind == reflect.Struct {
				err := mapForm(structField.Addr().Interface(), form)
				if err != nil {
					return err
				}
				continue
			}
		}

		inputValue, exists := form[inputFieldName]
		if !exists {
			continue
		}

		numElems := len(inputValue)
		if structfieldKind == reflect.Slice && numElems > 0 {
			sliceOf := structField.Type().Elem().Kind()
			slice := reflect.MakeSlice(structField.Type(), numElems, numElems)

			for i := 0; i < numElems; i++ {
				if err:= setWithProperType(sliceOf, inputValue[i], slice.Index(i)); err != nil {
					return err
				}
			}
			val.Field(i).Set(slice)
		} else {
			if _, isTime := structField.Interface().(time.Time); isTime {
				if err := setTimeField(inputValue[0], typeField, structField);  err != nil{
					return err
				}
				continue
			}
			if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
				return err
			}
		}
	}
	return nil
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32,structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}

	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}

	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}

	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}

	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		timeFormat = "2021-01-01 01:01:00"
		val = strings.Replace(val, "/", "-", 0)
		num := len(strings.Split(val, " "))
		if num == 1 {
			val = val + " 00:00:00"
		} else {
			num = len(strings.Split(val, ":"))
			if num == 1 {
				val = val + ":00:00"
			} else if num == 2 {
				val = val + ":00"
			}
 		}
	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _:= strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func ensureNotPointer(obj interface{})  {
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		panic("Pointers are not accepted as binding models")
	}
}

