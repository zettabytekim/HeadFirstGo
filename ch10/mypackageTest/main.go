package main

import "HeadFirstGo/ch10/mypackage"

func main() {
	value := mypackage.MyType{}
	value.ExportedMethod()
}
