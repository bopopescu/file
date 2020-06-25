package controller

import (
	"net/http"
	"strconv"
	"strings"

	"context"
	"reflect"

	"Mycrwaler/crawler-frontend/model"
	"Mycrwaler/crawler-frontend/view"
	"Mycrwaler/crawler-single/engine"
	"regexp"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(
		req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q

	resp, err := h.client.
		Search("myelas2").Type("zhenai").
	//Search(config.ElasticIndex).
		Query(elastic.NewQueryStringQuery(
		rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from

	//for _,v :=range resp.Each(
	//	reflect.TypeOf(engine.Item{})){
	//		item:=v.(engine.Item)
	//}


	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom =
			(result.Start - 1) /
				pageSize * pageSize
	}



	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

//Sex:女  AND Age:<22
//localhost:9200/myelas2/zhenai/_search?q=Payload.Age:<30 AND Payload.Sex:女
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
