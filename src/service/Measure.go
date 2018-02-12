package service

import (
	"database/sql"
	"fmt"
	"measure/src/db"
	"measure/src/db/Schema"
	"reflect"
)

func FindMeasureById(id int64) Schema.Measure {
	measure := Schema.Measure{}
	t := reflect.TypeOf(measure)
	v := reflect.ValueOf(&measure).Elem()
	params := []interface{}{}
	sqlSelect := "SELECT "
	for k := 0; k < t.NumField(); k++ {
		sqlSelect = fmt.Sprintf("%s%s,", sqlSelect, t.Field(k).Name)
		var temp string
		params = append(params, &temp)
	}
	sqlSelect = fmt.Sprintf("%s %s", sqlSelect[:len(sqlSelect)-1], " FROM measure WHERE id=?")
	stms, err := db.DB.Prepare(sqlSelect)
	if err != nil {
		panic(err.Error())
	}

	row := stms.QueryRow(id)
	err = row.Scan(params...)
	if err == sql.ErrNoRows {
		return Schema.Measure{Id: "0"}
	} else if err != nil {
		panic(err.Error())
	}
	stms.Close()
	for i := 0; i < len(params); i++ {
		s := *params[i].(*string)
		v.FieldByName(t.Field(i).Name).Set(reflect.ValueOf(s))
	}
	return measure
}

func FindList() []Schema.MeasureSimple {
	var result []Schema.MeasureSimple
	stms, err := db.DB.Prepare("SELECT id,name,company,phone,address FROM measure")
	if err != nil {
		panic(err)
	}
	defer stms.Close()
	rows, err := stms.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, company, address, phone string
		err = rows.Scan(&id, &name, &company, &phone, &address)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, Schema.MeasureSimple{id, name, company, address, phone})
	}
	return result

}

func AddMeasure(measure Schema.Measure) (int64, error) {
	t := reflect.TypeOf(measure)
	v := reflect.ValueOf(measure)
	sql := "insert into measure("
	values := "values("
	params := []string{}
	for k := 0; k < t.NumField(); k++ {
		sql = fmt.Sprintf("%s%s,", sql, t.Field(k).Name)
		values = fmt.Sprintf("%s'%s',", values, v.Field(k).Interface())
		params = append(params, fmt.Sprint(v.Field(k).Interface()))
	}
	sql = fmt.Sprintf("%s)", sql[:len(sql)-1])
	values = fmt.Sprintf("%s)", values[:len(values)-1])
	stms, err := db.DB.Prepare(fmt.Sprintf("%s %s", sql, values))
	if err != nil {
		panic(err)
	}
	defer stms.Close()

	result, err := stms.Exec()

	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func EditMeasure(measure Schema.Measure) (int64, error) {
	t := reflect.TypeOf(measure)
	v := reflect.ValueOf(measure)
	sql := "update measure set "
	params := []interface{}{}
	for k := 0; k < t.NumField(); k++ {
		if t.Field(k).Name != "id" {
			sql = fmt.Sprintf("%s%s = ?,", sql, t.Field(k).Name)
			params = append(params, fmt.Sprint(v.Field(k).Interface()))
		}
	}
	params = append(params, measure.Id)
	sql = fmt.Sprintf("%s WHERE id=?", sql[:len(sql)-1])

	stms, err := db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(params...)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer stms.Close()
	return result.RowsAffected()
}

func DeleteMeasureById(id int64) (int64, error) {
	stms, err := db.DB.Prepare("delete from measure WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(id)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	stms.Close()
	return result.RowsAffected()
}
