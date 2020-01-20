package main

import (
	"log"
	"thehappymouse/ccmouse/crawler_distributed/rpcsupport"
	"fmt"
	"thehappymouse/ccmouse/crawler_distributed/config"
	"thehappymouse/ccmouse/crawler_distributed/worker"
	"flag"
)

var port = flag.Int("port", config.WorkerPort0, "请输入工作端口号(默认10086）")

func main() {
	flag.Parse()
	port2 := fmt.Sprintf(":%d", *port)

	fmt.Println("Worker Server Start At:", port2)
	log.Fatal(rpcsupport.ServeRpc(port2,
		worker.CrawlService{}))
}
