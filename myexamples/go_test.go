package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	var eventGroup uint32 = 0
	var mapGroup = map[uint32]string{
		// 1: "运营服务组",
		2: "信令服务组",
		3: "接入服务组",
		4: "媒体服务组",
	}
	sig, access, media, other := 0, 0, 0, 0
	for i := 0; i < 100; i++ {
		groupName, ok := mapGroup[eventGroup]
		if !ok {
			var sliceGroup []string
			for _, v := range mapGroup {
				sliceGroup = append(sliceGroup, v)
			}
			// 没有配置分组，随机选一个分组。
			rand.Seed(time.Now().Unix())
			n := rand.Intn(len(mapGroup))
			groupName = sliceGroup[n]
			//return nil
		}
		if groupName == "信令服务组" {
			sig++
		} else if groupName == "接入服务组" {
			access++
		} else if groupName == "媒体服务组" {
			media++
		} else {
			other++
		}
	}
	fmt.Printf("sig=%d, access=%d, media=%d, other=%d", sig, access, media, other)

}
