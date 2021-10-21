package main

import "github.com/lshsuper/go-webboot/src/core"

func main()  {


	app:=core.NewWebBootServer(":10086")
	defer app.Stop()

	//选择自动注册
	app.AutoRegister()
	//启动
	app.Start()


}
