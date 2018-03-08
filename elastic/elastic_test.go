package elastic_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/olivere/elastic"
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
	c, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	client = c
	setup(client)
	code := m.Run()
	shutdown(client)
	os.Exit(code)
}

func setup(c *elastic.Client) {
	_, err := c.Index().Index("blog").Type("article").Id("1").BodyJson(&blog{
		Title:   "New version of Elasticsearchan released!",
		Content: "Version 1.0 released today!",
		Tags:    []string{"announce", "elasticsearch", "release"},
	}).Do(ctx)
	if err != nil {
		panic(err)
	}
}

func shutdown(c *elastic.Client) {
	_, err := client.DeleteIndex("blog").Do(ctx)
	if err != nil {
		panic(err)
	}
	c.Stop()
}

func TestGetADocument(t *testing.T) {
	resp, err := client.Get().Index("blog").Type("article").Id("1").Do(ctx)
	if err != nil {
		t.Error(err)
	}
	var b blog
	json.Unmarshal(*resp.Source, &b) // nolint: gas
	t.Logf("%v %#v %v", resp, b, *resp.Version)
}

func TestUpdateADocument(t *testing.T) {
	resp, err := client.Update().Index("blog").Type("article").Id("1").Script(elastic.NewScriptInline("ctx._source.content = \"new content1\"")).Do(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
