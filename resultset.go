package db

import (
	"reflect"
	"strconv"
	"time"
)

//Resultset ...
type Resultset map[string]interface{}

//IntNil ...
func (r Resultset) IntNil(name string) *int {
	if val := r.getValue(name); val != nil {
		switch val := val.Interface().(type) {
		case **uint64:
			if *val == nil {
				return nil
			}
			intVal := int(**val)
			return &intVal
		case **int64:
			if *val == nil {
				return nil
			}
			intVal := int(**val)
			return &intVal
		case **float64:
			if *val == nil {
				return nil
			}
			intVal := int(**val)
			return &intVal
		case *[]uint8:
			if intVal, e := strconv.Atoi(string(*val)); e == nil {
				return &intVal
			}
			return nil
		case **string:
			if *val == nil {
				return nil
			}
			if intVal, e := strconv.Atoi(**val); e == nil {
				return &intVal
			}
			return nil
		default:

		}
	}
	return nil
}

//Int ...
func (r Resultset) Int(name string) int {
	if val := r.IntNil(name); val != nil {
		return *val
	}
	return 0
}

//IntOr ...
func (r Resultset) IntOr(name string, defValue int) int {
	if val := r.IntNil(name); val != nil {
		return *val
	}
	return defValue
}

//TimeNil ...
func (r Resultset) TimeNil(name string) *time.Time {
	if val := r.getValue(name); val != nil {
		if val, ok := val.Interface().(**time.Time); ok {
			if *val == nil {
				return nil
			}
			return *val
		}
	}
	return nil
}

//Time ...
func (r Resultset) Time(name string) time.Time {
	if val := r.TimeNil(name); val != nil {
		return *val
	}
	return time.Time{}
}

func (r Resultset) getValue(name string) *reflect.Value {
	if val, ok := r[name]; ok {
		val := reflect.ValueOf(val)
		return &val
	}
	return nil
}

//FloatNil ...
func (r Resultset) FloatNil(name string) *float64 {
	if val := r.getValue(name); val != nil {
		switch val := val.Interface().(type) {
		case **uint64:
			if *val == nil {
				return nil
			}
			floatVal := float64(int(**val))
			return &floatVal
		case **int64:
			if *val == nil {
				return nil
			}
			floatVal := float64(int(**val))
			return &floatVal
		case **float64:
			return *val
		}
	}
	return nil
}

//Float ...
func (r Resultset) Float(name string) float64 {
	if val := r.FloatNil(name); val != nil {
		return *val
	}
	return 0
}

//FloatOr ...
func (r Resultset) FloatOr(name string, defValue float64) float64 {
	if val := r.FloatNil(name); val != nil {
		return *val
	}
	return defValue
}

//StringNil ...
func (r Resultset) StringNil(name string) *string {
	if val := r.getValue(name); val != nil {
		if reflect.ValueOf(val.Interface()).Elem().IsNil() {
			return nil
		}
		str := ``
		switch val := val.Interface().(type) {
		case **string:
			str = **val
		case *[]byte:
			str = string(*val)
		case **uint64:
			str = strconv.FormatUint(**val, 10)
		case **int64:
			str = strconv.FormatInt(**val, 10)
		case **float64:
			str = strconv.FormatUint(uint64(**val), 10)
		case *time.Time:
			str = val.String()
		case **time.Time:
			v := *val
			str = v.String()
		default:
			println(`unable to parse string from ` + reflect.TypeOf(val).String())
		}
		return &str
	}
	return nil
}

//Interface ...
func (r Resultset) Interface(name string) interface{} {
	if val := r.getValue(name); val != nil {
		return val.Interface()
	}
	return nil
}

//Bytes ...
func (r Resultset) Bytes(name string) []byte {
	if val := r.getValue(name); val != nil {
		if reflect.ValueOf(val.Interface()).Elem().IsNil() {
			return nil
		}
		switch val := val.Interface().(type) {
		case *[]byte:
			return *val
		case **uint64:
			return []byte(strconv.FormatUint(**val, 10))
		case **int64:
			return []byte(strconv.FormatInt(**val, 10))
		case **float64:
			return []byte(strconv.FormatUint(uint64(**val), 10))
		default:
			return []byte(``)
		}
	}
	return nil
}

//String ...
func (r Resultset) String(name string) string {
	if val := r.StringNil(name); val != nil {
		return *val
	}
	return ``
}

//StringOr ...
func (r Resultset) StringOr(name string, defValue string) string {
	if val := r.StringNil(name); val != nil {
		return *val
	}
	return defValue
}
