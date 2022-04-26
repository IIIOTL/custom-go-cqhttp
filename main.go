package main

import (
	"github.com/IIIOTL/custom-go-cqhttp/cmd/gocq"

	_ "github.com/IIIOTL/custom-go-cqhttp/db/leveldb"    // leveldb
	_ "github.com/IIIOTL/custom-go-cqhttp/modules/mime"  // mime检查模块
	_ "github.com/IIIOTL/custom-go-cqhttp/modules/pprof" // pprof 性能分析
	_ "github.com/IIIOTL/custom-go-cqhttp/modules/silk"  // silk编码模块
)

func main() {
	gocq.Main()
}
