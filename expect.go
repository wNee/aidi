package aidi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// ExpectFunc function type used as argument to Expect()
type ExpectFunc func(a *Aidi) (bool, string)

// Expect Checks according to the given function, which allows you to describe any kind of assertion.
func (a *Aidi) Expect(foo ExpectFunc) *Aidi {
	//Global.NumAsserts++
	if ok, err_str := foo(a); !ok {
		a.AddError(err_str)
	}
	return a
}

// Checks for header and if values match
func (a *Aidi) ExpectHeader(key, value string) *Aidi {
	Global.NumAsserts++
	chk_val := a.Resp.Header.Get(key)
	if chk_val == "" {
		err_str := fmt.Sprintf("Expected Header %q, but it was missing", key)
		a.AddError(err_str)
	} else if chk_val != value {
		err_str := fmt.Sprintf("Expected Header %q to be %q, but got %q", key, value, chk_val)
		a.AddError(err_str)
	}
	return a
}

// ExpectBodyJson checks if the body of the response
// equal the json
//
// bodyJson the except response body json string
func (a *Aidi) ExpectBodyJson(bodyJson string) *Aidi {
	Global.NumAsserts++

	buf := new(bytes.Buffer)
	buf.ReadFrom(a.Resp.Body)
	s := buf.String()
	equal, err := areEqualJSON(bodyJson, s)
	if err != nil {
		a.AddError(err.Error())
		return a
	}
	if !equal {
		a.AddError(fmt.Sprintf("ExpectBody equality test failed for %s, got value: %s", bodyJson, s))
	}
	return a
}

// ExpectBodyContainJson checks if the body of the response
// contain the json
//
// bodyJson the except response body json string
func (a *Aidi) ExpectBodyContainJson(bodyJson string) *Aidi {
	Global.NumAsserts++

	buf := new(bytes.Buffer)
	buf.ReadFrom(a.Resp.Body)
	s := buf.String()
	equal, err := containsJSON(bodyJson, s)
	if err != nil {
		a.AddError(err.Error())
		return a
	}
	if !equal {
		a.AddError(fmt.Sprintf("ExpectBody equality test failed for %s, got value: %s", bodyJson, s))
	}
	return a
}

func containsJSON(s1, s2 string) (bool, error) {
	var m1 interface{}
	var m2 interface{}
	if err := json.Unmarshal([]byte(s1), &m1); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(s2), &m2); err != nil {
		return false, err
	}

	return deepContrains(m1, m2), nil
}

func deepContrains(x, y interface{}) bool {
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)

	if v1.Type() != v2.Type() {
		return false
	}
	return deepValueContains(v1, v2, 0)
}

func deepValueContains(v1, v2 reflect.Value, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return !v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}
	// 10层以上深度的默认为false
	if depth > 10 {
		return false
	}

	switch v2.Kind() {
	case reflect.Array:
		for i := 0; i < v2.Len(); i++ {
			if !deepValueContains(v1.Index(i), v2.Index(i), depth+1) {
				return false
			}
		}
		return true
	case reflect.Slice:
		if v1.IsNil() || v2.IsNil() {
			return v2.IsNil()
		}
		if v1.Len() < v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v2.Len(); i++ {
			if !deepValueContains(v1.Index(i), v2.Index(i), depth+1) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, v2.NumField(); i < n; i++ {
			if !deepValueContains(v1.Field(i), v2.Field(i), depth+1) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() || v2.IsNil() {
			return v2.IsNil()
		}
		if v1.Len() < v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v2.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueContains(val1, val2, depth+1) {
				return false
			}
		}
		return true
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v2.IsNil()
		}
		return deepValueContains(v1.Elem(), v2.Elem(), depth+1)
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return v1.Uint() == v2.Uint()
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	default:
		return  false
	}
}

func areEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string %s :: %s", s1, err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string %s :: %s", s2, err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

// Prints a report for the Aidi Object
//
// If there are any errors, they will all be printed as well
func (a *Aidi) PrintReport() *Aidi {
	if len(a.Errs) == 0 {
		fmt.Printf("Pass  [%s]\n", a.Name)
	} else {
		fmt.Printf("FAIL  [%s]\n", a.Name)
		for _, e := range a.Errs {
			fmt.Println("        - ", e)
		}
	}

	return a
}

// Prints a report for the Aidi Object in go_test format
//
// If there are any errors, they will all be printed as well
func (a *Aidi) PrintGoTestReport() *Aidi {
	if len(a.Errs) == 0 {
		fmt.Printf("=== RUN   %s\n--- PASS: %s (%.2fs)\n", a.Name, a.Name, a.ExecutionTime)
	} else {
		fmt.Printf("=== RUN   %s\n--- FAIL: %s (%.2fs)\n", a.Name, a.Name, a.ExecutionTime)
		for _, e := range a.Errs {
			fmt.Println("	", e)
		}
	}
	return a
}

// Checks the response status code
func (a *Aidi) ExpectStatus(code int) *Aidi {
	Global.NumAsserts++
	status := a.Resp.StatusCode
	if status != code {
		err_str := fmt.Sprintf("Expected Status %d, but got %d: %q", code, status, a.Resp.Status)
		a.AddError(err_str)
	}
	return a
}
