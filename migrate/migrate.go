package main

import (
	"github.com/takahiromitsui/go-task-manager/model"
	"github.com/takahiromitsui/go-task-manager/util"
)

func init() {
	util.LoadEnv()
	util.ConnectToDB()
}

func main() {
	util.DB.AutoMigrate(&model.Task{})
}