package main

import (
	"hdg.com/auto-demo/src/service"
	_ "hdg.com/auto-demo/src/common"
	"fmt"
	"time"
)

func main() {
	fmt.Println("===>开始<===")
	start := time.Now()
	dispatch := service.NewServerService()
	dispatch.Dispatch()
	end := time.Now()
	fmt.Println("===>结束<===")
	fmt.Println("耗时:", end.Sub(start))
}
