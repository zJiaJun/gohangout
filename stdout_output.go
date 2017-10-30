package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

type StdoutOutput struct {
}

func (outputPlugin *StdoutOutput) emit(event map[string]interface{}) {
	buf, err := json.Marshal(event)
	if err != nil {
		glog.Errorf("marshal %v error:%s", event, err)
	}
	fmt.Println(string(buf))
}