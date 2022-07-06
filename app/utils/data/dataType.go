package data

import (
	"encoding/json"
	"strconv"
	"strings"
)

func StringToInt(string2 string) int {
	int, err := strconv.Atoi(string2)
	if err!=nil{
		return 0
	}
	return  int
}
func IntToString(int2 int) string {
	return strconv.Itoa(int2)
}
func UintToString(int2 uint) string {
	return strconv.Itoa(int(int2))
}

func MapToStruct(mapData interface{},structData interface{}) error {
	//集合转json
	tmp,err:=json.Marshal(mapData)
	if err!=nil{
		return err
	}
	err = json.Unmarshal(tmp,structData)
	if err!=nil{
		return err
	}
	return nil
}

func StringToMap(string2 string,sep string) []string  {
	return strings.Split(string2,sep)
}