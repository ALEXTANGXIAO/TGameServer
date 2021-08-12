package tserver

import (
	"fmt"
)

var EventerMap = make(map[string][]func(interface{}))

func ResigterEvent(eventKey string, callback func(interface{})) {
	list := EventerMap[eventKey]
	list = append(list, callback)
	EventerMap[eventKey] = list
}

func CallEvent(eventKey string, param interface{}) {
	list := EventerMap[eventKey]
	for _, callback := range list {
		callback(param)
	}
}

func RemoveEvent(eventKey string) {
	_, ok := EventerMap[eventKey]

	if ok {
		delete(EventerMap, eventKey)
	} else {
		fmt.Print("Had Not Keys", eventKey)
	}
}

////////////////////////////////////////////////////////////////
// 	ResigterEvent("Test",test)
// 	CallEvent("Test",nil)
////////////////////////////////////////////////////////////////
