package main

import (
	"go-homework/week9"
)

func main() {
	//// 第二周作业
	//_ = week2.HandleError()
	//
	//// 第三周作业
	//s := week3.NewServer()
	//cgiS := week3.NewCgiService()
	//logS := week3.NewLogService()
	//s.Register(cgiS.Name, &cgiS)
	//s.Register(logS.Name, &logS)
	//if err := s.Start(); err != nil {
	//	log.Fatal(err)
	//}

	week9.Serve()
}
