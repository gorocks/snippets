package elastic_test

import (
	"context"
	"log"
	"os"
	"testing"

	elastic "gopkg.in/olivere/elastic.v5"
)

type blog struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

var (
	ctx    context.Context
	client *elastic.Client
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	c, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetTraceLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}
	client = c
	code := m.Run()
	c.Stop()
	os.Exit(code)
}

func deleteIndex(t *testing.T, c *elastic.Client, index string) {
	_, err := client.DeleteIndex(index).Do(ctx)
	if err != nil {
		t.Error(err)
	}
}

type doc map[string]interface{}

// https://ume.la/4xyjUg
func TestDealWithNull(t *testing.T) {
	defer deleteIndex(t, client, "my_index")

	br, err := client.Bulk().Index("my_index").Type("posts").Add(
		elastic.NewBulkIndexRequest().Id("1").Doc(doc{"tags": []string{"search"}}),
		elastic.NewBulkIndexRequest().Id("2").Doc(doc{"tags": []string{"search", "open_source"}}),
		elastic.NewBulkIndexRequest().Id("3").Doc(doc{"other_field": "some data"}),
		elastic.NewBulkIndexRequest().Id("4").Doc(doc{"tags": nil}),
		elastic.NewBulkIndexRequest().Id("5").Doc(doc{"tags": []interface{}{"search", nil}}),
	).Do(ctx)
	for _, v := range br.Indexed() {
		t.Log(v.Status)
	}

	sr, err := client.Search().Index("my_index").Type("posts").Query(elastic.NewConstantScoreQuery(elastic.NewExistsQuery("tags"))).Do(ctx)
	t.Log(sr.TotalHits(), err)

	sr, err = client.Search().Index("my_index").Type("posts").Query(elastic.NewConstantScoreQuery(elastic.NewBoolQuery().MustNot(elastic.NewExistsQuery("tags")))).Do(ctx)
	t.Log(sr.TotalHits(), err)
}
