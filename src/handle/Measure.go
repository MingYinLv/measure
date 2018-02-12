package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"measure/src/db/Schema"
	"measure/src/service"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func strFirstToLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func AddMeasure(c *gin.Context) {
	measure := Schema.Measure{}
	t := reflect.TypeOf(measure)
	v := reflect.ValueOf(&measure).Elem()

	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		val, ok := c.GetPostForm(strFirstToLower(name))
		if ok {
			v.FieldByName(name).Set(reflect.ValueOf(val))
		}
	}

	id, err := service.AddMeasure(measure)

	if err == nil {
		strId := strconv.Itoa(int(id))
		measure.Id = strId
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": measure,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

}

func FindList(c *gin.Context) {
	measures := service.FindList()

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": measures,
	})
}

func FindById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	measure := service.FindMeasureById(int64(id))
	if measure.Id == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": measure,
		})
	}
}

func EditMeasure(c *gin.Context) {
	measure := Schema.Measure{}
	measure.Id = c.Param("id")
	t := reflect.TypeOf(measure)
	v := reflect.ValueOf(&measure).Elem()

	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		val, ok := c.GetPostForm(strFirstToLower(name))
		if ok {
			v.FieldByName(name).Set(reflect.ValueOf(val))
		}
	}

	_, err := service.EditMeasure(measure)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
}

func DownloadMeasure(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	measure := service.FindMeasureById(int64(id))
	if measure.Id == "0" {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.Header("Content-Disposition", "attachment; filename=测量数据.txt")
		t := reflect.TypeOf(measure)
		v := reflect.ValueOf(measure)
		result := ""
		for k := 0; k < t.NumField(); k++ {
			name := t.Field(k).Name
			result += fmt.Sprintf("%s:%s\n", Schema.KeyText[strFirstToLower(name)], v.Field(k).Interface())
		}
		c.String(http.StatusOK, result)
	}
}

func DeleteMeasure(c *gin.Context){
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}

	service.DeleteMeasureById(int64(id))
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "删除成功",
	})
}
