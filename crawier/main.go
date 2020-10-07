package main

import (
	"goWorkspace/crawier/engine"
	"goWorkspace/crawier/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
