package JsonService

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserBase struct {
	Direction string `json:"direction"`
	Province  string `json:"province"`
}

type UserInfo struct {
	UserBase UserBase `json:"userbase"`
	Title    []string `json:"title"`
	Name     string   `json:"name"`
}

func TestCompareJson(t *testing.T) {
	userInfo := new(UserInfo)
	userInfo.Title = []string{"今日最佳", "本周最佳"}
	userInfo.Name = "kouhaozhe"
	userInfo.UserBase.Province = "shannxi"
	userInfo.UserBase.Direction = "yanan"

	mapUserInfo := make(map[string]interface{})

	userByte, _ := json.Marshal(userInfo)

	err := json.Unmarshal(userByte, &mapUserInfo)
	if err != nil {
		panic("Unmarshal Error")
	}

	fmt.Println(true, CompareJsonDict(mapUserInfo, mapUserInfo))
}

func TestCompareJsonTwo(t *testing.T) {
	userInfo := new(UserInfo)
	userInfo.Title = []string{"今日最佳", "本周最佳"}
	userInfo.Name = "kouhaozhe"
	userInfo.UserBase.Province = "shannxi"
	userInfo.UserBase.Direction = "yanan"

	mapUserInfo := make(map[string]interface{})

	userByte, _ := json.Marshal(userInfo)

	err := json.Unmarshal(userByte, &mapUserInfo)
	if err != nil {
		panic("Unmarshal Error")
	}

	userInfoTwo := new(UserInfo)
	userInfoTwo.Title = []string{"今日最佳", "本周最佳"}
	userInfoTwo.Name = "kouhaozhe"
	userInfoTwo.UserBase.Province = "shannxi"
	userInfoTwo.UserBase.Direction = "xian"

	mapUserInfoTwo := make(map[string]interface{})

	userByte, _ = json.Marshal(userInfoTwo)

	err = json.Unmarshal(userByte, &mapUserInfoTwo)
	if err != nil {
		panic("Unmarshal Error")
	}

	fmt.Println(false, CompareJsonDict(mapUserInfo, mapUserInfoTwo))
}

func TestCompareJsonThree(t *testing.T) {
	userInfo := new(UserInfo)
	userInfo.Title = []string{"今日最佳", "本周最佳"}
	userInfo.Name = "kouhaozhe"
	userInfo.UserBase.Province = "shannxi"
	userInfo.UserBase.Direction = "yanan"

	mapUserInfo := make(map[string]interface{})

	userByte, _ := json.Marshal(userInfo)

	err := json.Unmarshal(userByte, &mapUserInfo)
	if err != nil {
		panic("Unmarshal Error")
	}

	userInfoTwo := new(UserInfo)
	userInfoTwo.Title = []string{"今日最佳", "本周最佳"}
	userInfoTwo.Name = "haozhe"
	userInfoTwo.UserBase.Province = "shannxi"
	userInfoTwo.UserBase.Direction = "yanan"

	mapUserInfoTwo := make(map[string]interface{})

	userByte, _ = json.Marshal(userInfoTwo)

	err = json.Unmarshal(userByte, &mapUserInfoTwo)
	if err != nil {
		panic("Unmarshal Error")
	}

	fmt.Println(false, CompareJsonDict(mapUserInfo, mapUserInfoTwo))
}
