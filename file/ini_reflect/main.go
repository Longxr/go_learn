package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini配置文件解析

// MysqlConfig ...
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Test     bool   `ini:"test"`
}

// Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//打开配置文件读取数据
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a pointer")
		return
	}
	//按行读取数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Printf("%#v\n", lineSlice)
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//；#开头是注释
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//[开头是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//根据sectionName去data里找对应结构体
			// v := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//找到对应的结构体
					structName = field.Name
					fmt.Printf("找到%s对应的结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//=分割的是键值对
			equalIndex := strings.Index(line, "=")
			if equalIndex == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			key := strings.TrimSpace(line[:equalIndex])
			value := strings.TrimSpace(line[equalIndex+1:])
			//根据structName去data里取字段
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //嵌套结构体的值
			sType := sValue.Type()                     //嵌套结构体的类型信息，包括tag
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段不是结构体", structName)
				return
			}
			//遍历嵌套结构体
			var fieldName string
			var field reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field = sType.Field(i) //结构体的字段
				if field.Tag.Get("ini") == key {
					//找到了对应的字段
					fieldName = field.Name
					break
				}
			}
			//未找到对应字段
			if len(fieldName) == 0 {
				continue
			}
			//根据字段取出结构体中的字段
			fileObj := sValue.FieldByName(fieldName) //结构体字段对应的值
			fmt.Println(fieldName, field.Type.Kind())
			switch field.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// fmt.Println(cfg.Address, cfg.Port, cfg.UserName, cfg.Password)
	fmt.Printf("%#v\n", cfg)
}
