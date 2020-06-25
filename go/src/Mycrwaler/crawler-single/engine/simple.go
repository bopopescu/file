package engine

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// 单任务版引擎
type SimpleEngine struct {
}

func (e *SimpleEngine) 	Run(queue ...Request) {
	var count = 0
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fmt.Print(123)
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		results, err := Worker(r)


		if err != nil {
			continue
		}


		fmt.Print(321)


		for _, r := range results.Requests {

			fmt.Println(r)
			//queue = append(queue, r)
		}

		//queue = append(queue, results.Requests...)

		fmt.Print(results)
		for _, item := range results.Items {
			count++
			fmt.Println(1111)
			log.Info().Msgf("Got Item: $%d %v", count, item)
		}
	}
}
