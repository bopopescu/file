package engine

import (
	"fmt"

	"thehappymouse/ccmouse/crawler/fetcher"
	"github.com/rs/zerolog/log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	fmt.Printf("Fetching %s\n", r.Url)
	fmt.Println()
	if err != nil {
		log.Error().Msgf("请求[%s]失败：%s", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parse.Parse(body,r.Url), nil
}
