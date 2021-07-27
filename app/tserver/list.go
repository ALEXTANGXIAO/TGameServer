package tserver

import (
	"errors"
	"fmt"
)

func main() {
	values := []interface{}{1, 2, 3, 5}
	newValues, err := Add(values, 4, 3)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//期望输出0，1，2，3，4，5
	for i, v := range newValues {
		fmt.Printf("下标:%d,值:%d\n", i, v)
	}
}

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {
	//校验检查
	if index < 0 || index > len(values) {
		return nil, errors.New("index out of Range!")
	}

	//方式一使用子切片，子切片会扩容，只是写法简短，但是没有性能上的好处
	// res := values[:index]

	//方式二 普通循环
	res := []interface{}{}

	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}

	res = append(res, val)

	for i := index; i < len(values); i++ {
		v := values[i]
		res = append(res, v)
	}

	return res, nil
}

func RemoveC(values []*Client, val *Client) []*Client {
	if len(values) <= 0 {
		return values
	}

	res := []*Client{}

	for i := 0; i < len(values); i++ {
		if values[i] == val {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}

func Remove(values []interface{}, val interface{}) []interface{} {

	if len(values) <= 0 {
		return values
	}

	res := []interface{}{}

	for i := 0; i < len(values); i++ {
		if values[i] == val {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}

func Delete(values []interface{}, index int) []interface{} {

	if index < 0 || index > len(values) {
		return values
	}

	res := []interface{}{}

	for i := 0; i < len(values); i++ {
		if index == i {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}
