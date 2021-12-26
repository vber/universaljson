package universaljson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type UniversalJSON struct {
	rawData interface{}
}

func (json *UniversalJSON) getValue(fieldname string) reflect.Value {
	root := reflect.ValueOf(json.rawData)

	for _, k := range root.MapKeys() {
		s_k := k.String()

		if k.String() == fieldname {
			value := reflect.ValueOf(s_k)
			return reflect.ValueOf(root.MapIndex(value).Interface())
		}
	}
	return reflect.ValueOf(nil)
}

func (json *UniversalJSON) GetFloat64(fieldname string) (float64, error) {
	ret := json.getValue(fieldname)
	if ret == reflect.ValueOf(nil) {
		return 0, errors.New(fmt.Sprintf("Field:%s not found.", fieldname))
	}
	return ret.Float(), nil
}

func (json *UniversalJSON) GetInt64(fieldname string) (int64, error) {
	ret := json.getValue(fieldname)
	if ret == reflect.ValueOf(nil) {
		return 0, errors.New(fmt.Sprintf("Field:%s not found.", fieldname))
	}
	if ret.Kind() == reflect.Float64 {
		return int64(ret.Float()), nil
	} else {
		return 0, errors.New(fmt.Sprintf("Field %s type is mismatch!", fieldname))
	}
}

func (json *UniversalJSON) GetObject(fieldname string) *UniversalJSON {
	var (
		next_json *UniversalJSON
	)
	ret := json.getValue(fieldname)
	if ret == reflect.ValueOf(nil) {
		return nil
	}

	if ret.Kind() == reflect.Map {
		next_json = new(UniversalJSON)
		next_json.rawData = ret.Interface()
	}

	return next_json
}

func (json *UniversalJSON) GetArray(fieldname string) []*UniversalJSON {
	var (
		next_json []*UniversalJSON
	)

	ret := json.getValue(fieldname)
	if ret == reflect.ValueOf(nil) {
		return nil
	}

	if ret.Kind() == reflect.Slice {
		nums := ret.Len()
		next_json = make([]*UniversalJSON, nums)
		for i := 0; i < nums; i++ {
			next_json[i] = new(UniversalJSON)
			next_json[i].rawData = ret.Index(i).Interface()
		}
	}

	return next_json
}

func (json *UniversalJSON) GetString(fieldname string) (string, error) {
	ret := json.getValue(fieldname)
	if ret == reflect.ValueOf(nil) {
		return "", errors.New(fmt.Sprintf("Field:%s not found.", fieldname))
	}
	return ret.String(), nil
}

func ParseJSON(v interface{}) *UniversalJSON {
	var (
		uj  *UniversalJSON
		x   interface{}
		err error
	)
	uj = new(UniversalJSON)

	switch t := v.(type) {
	case string:
		if err = json.Unmarshal([]byte(t), &x); err != nil {
			return nil
		}
		uj.rawData = x
	case *string:
		if err = json.Unmarshal([]byte(*t), &x); err != nil {
			return nil
		}
		uj.rawData = x
	case []byte:
		if err = json.Unmarshal(t, &x); err != nil {
			return nil
		}
		uj.rawData = x
	case map[string]interface{}:
		uj.rawData = t
		return uj
	default:
		return nil
	}
	return uj
}
