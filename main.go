package main

import (
	"log"

	"go-homework/week2"
	"go-homework/week3"
)

func main() {
	// 第二周作业
	_ = week2.HandleError()

	// 第三周作业
	s := week3.NewServer()
	cgiS := week3.NewCgiService()
	logS := week3.NewLogService()
	s.Register(week3.Service(cgiS))
	s.Register(week3.Service(logS))
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
