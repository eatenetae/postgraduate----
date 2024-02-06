package redis_ser

import (
	"server/global"
	"strconv"
)

type CountDB struct {
	Index string // 前缀索引
}

func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}

// SetCount 在原有基础上增加一个值
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()
	return err
}

func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

func (c CountDB) GetInfo() map[string]int {
	var Info = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		Info[id] = num
	}
	return Info
}

func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
