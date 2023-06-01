package main

import (
	"errors"
	"fmt"
	"reflect"
)

func USAGE_METHOD_OF_TypeOf() {
	type favContextKey string
	k := favContextKey("language")
	fmt.Println(reflect.TypeOf(k)) //main.favContextKey
}

func HOW_TO_CHECK_STRUCT_MATCH() {
	type DataA struct {
		Name string
		Age  int
	}

	type DataB struct {
		Name  string
		Topic string
	}

	type DataC struct {
		Name   string
		Age    int
		Role   string
		Active bool
	}

	type Person struct {
		Name   string `json:"username"`
		Age    int    `json:"age"`
		Role   string `json:"role"`
		Active bool   `json:"active,omitempty"`
		// Description   string `json:"-"`
		// NotIncludeVal string
	}

	checkStructMatch := func(schema interface{}, obj interface{}) error {
		// 스키마와 객체가 모두 struct가 아니면 에러 반환
		s := reflect.ValueOf(schema).Type()
		o := reflect.ValueOf(obj).Type()
		if s.Kind() != reflect.Struct || o.Kind() != reflect.Struct {
			return fmt.Errorf("invalid argument type, expected struct, got %T and %T", schema, obj)
		}

		// for i := 0; i < s.NumField(); i++ {
		// 	schemaField := s.Field(i)
		// 	objField, ok := o.FieldByName(schemaField.Name)
		// 	if !ok {
		// 		return fmt.Errorf("need data field: %s", schemaField.Name)
		// 	}
		// 	// TODO: need research
		// 	if !objField.IsValid() || objField.IsZero() {
		// 		return fmt.Errorf("field '%s' is empty", objField.Name)
		// 	}
		// }

		return nil
	}

	p := Person{Name: "John", Age: 30, Role: "admin", Active: true}
	a := DataA{Name: "Alice", Age: 25}
	b := DataB{Name: "Bob", Topic: "technology"}
	c1 := DataC{Name: "Charlie", Age: 35, Role: "user", Active: false}
	c2 := DataC{Name: "Charlie", Age: 35, Role: "user"}

	err1 := checkStructMatch(p, p)
	err2 := checkStructMatch(p, a)
	err3 := checkStructMatch(p, b)
	err4 := checkStructMatch(p, c1)
	err5 := checkStructMatch(p, c2)

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)
	fmt.Println(err5)
}

func HOW_TO_CHECK_STRUCT_VALUES() {

	type Person struct {
		Name   string `json:"username"`
		Age    int    `json:"age"`
		Role   string `json:"role"`
		Active bool   `json:"active,omitempty"`
		// Description   string `json:"-"`
		// NotIncludeVal string
	}

	checkEmptyFields := func(obj interface{}) error {
		v := reflect.ValueOf(obj)
		// if v.Kind() == reflect.Pointer {

		// }

		// obj의 값이 struct가 아니면 에러 반환
		if v.Kind() != reflect.Struct {
			return errors.New("not a struct")
		}

		// struct 객체의 각 필드를 검사하여 값이 비어있으면 에러 반환
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Interface() == reflect.Zero(field.Type()).Interface() {
				return errors.New("empty field found")
			}
		}
		return nil
	}

	p1 := Person{Name: "John", Age: 30, Role: "IAM", Active: false}
	p2 := Person{Name: "Alice", Age: 24, Role: "ADMIN"}
	p3 := Person{Name: "Bob", Age: 43, Role: "MEMBER", Active: true}
	p4 := Person{Name: "", Age: 0}

	err1 := checkEmptyFields(p1)
	err2 := checkEmptyFields(p2)
	err3 := checkEmptyFields(p3)
	err4 := checkEmptyFields(p4)

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)
}

func main() {
	HOW_TO_CHECK_STRUCT_MATCH()
	// HOW_TO_CHECK_STRUCT_VALUES()
}
