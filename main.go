package main

import (
	"log"

	"go-homework/week2"
	"go-homework/week3"
	"go-homework/week9"
)

func main() {
	// 第二周作业
	_ = week2.HandleError()

	// 第三周作业
	s := week3.NewServer()
	cgiS := week3.NewCgiService()
	logS := week3.NewLogService()
	s.Register(cgiS.Name, &cgiS)
	s.Register(logS.Name, &logS)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// 第9周作业
	week9.Serve()
}
