package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type HuToolUtils interface {
	CopyProperties(itemOne interface{}, itemTwo interface{})
	IsEmpty(interface{}, interface{}) bool
	ParseList(strList string, newList interface{})
}

type BeanUtils struct{}

func NewBeanUtils() HuToolUtils {
	return &BeanUtils{}
}

func (b *BeanUtils) CopyProperties(itemOne interface{}, itemTwo interface{}) {
	by, err := json.Marshal(&itemOne)
	if err != nil {
		fmt.Printf("json marshal error:%v", err)
		return
	}
	err = json.Unmarshal(by, &itemTwo)
	if err != nil {
		fmt.Printf("json unmarshal error:%v", err)
		return
	}
}

func (b *BeanUtils) IsEmpty(itemOne interface{}, itemTwo interface{}) bool {
	return reflect.DeepEqual(itemOne, itemTwo)
}

func (b *BeanUtils) ParseList(strList string, newList interface{}) {
	err := json.Unmarshal([]byte(strList), &newList)
	if err != nil {
		fmt.Printf("json unmarshal error:%v", err)
		return
	}
}
