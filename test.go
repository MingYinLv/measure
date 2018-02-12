package main

import (
	"fmt"
	"measure/src/db/Schema"
	"reflect"
)

func main() {
	u := Schema.Measure{}
	t := reflect.TypeOf(u)
	//v := reflect.ValueOf(u)
	sql := "insert into users("
	values := "("
	for k := 0; k < t.NumField(); k++ {
		sql = fmt.Sprintf("%s%s,", sql, t.Field(k).Name)
		values = fmt.Sprintf("%s%s,", values, "?")
	}
	sql = fmt.Sprintf("%s)", sql[:len(sql)-1])
	values = fmt.Sprintf("%s)", values[:len(values)-1])

	fmt.Println(sql)
	fmt.Println(values)

	sql = "SELECT "
	for k := 0; k < t.NumField(); k++ {
		sql = fmt.Sprintf("%s%s,", sql, t.Field(k).Name)
	}
	sql = fmt.Sprintf("%s", sql[:len(sql)-1])
	sql = fmt.Sprintf("%s %s", sql, "FROM measure WHERE id=?")
	fmt.Println(sql)
}
