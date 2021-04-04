package base

import (
	"encoding/json"
	"log"
)

// Row 拓展信息
type Row struct {
	Row string `gorm:"column:row"` // 拓展信息
}

// RowToMap 从Row中根据JSON格式的数据反序列化出map
func (row *Row) RowToMap() (mapData map[string]string, err error) {
	mapData = make(map[string]string)
	err = json.Unmarshal([]byte(row.Row), &mapData)
	if err != nil {
		log.Printf("[Row][RowToMap] unmarshal error, err:%v", err)
		return map[string]string{}, err
	}
	return mapData, nil
}

// MapToRow 根据Map序列化出Row
func (row *Row) MapToRow(mapData map[string]string) (err error) {
	jsonString, err := json.Marshal(mapData)
	if err != nil {
		log.Printf("[Row][MapToRow] marshal error, err:%v", err)
		return err
	}
	row.Row = string(jsonString)
	return nil
}

// Set 插入或则覆盖一个键值对
func (row *Row) Set(key string, value string) (err error) {
	mapData, err := row.RowToMap()
	if err != nil {
		return err
	}
	mapData[key] = value
	return row.MapToRow(mapData)
}

// Delete 删除一个键
func (row *Row) Delete(key string) (err error) {
	mapData, err := row.RowToMap()
	if err != nil {
		return err
	}
	delete(mapData, key)
	return row.MapToRow(mapData)
}
