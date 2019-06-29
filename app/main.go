package main
import (
	"fmt"
	"github.com/olivere/elastic"
	"context"
)
func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()
	resp, err := client.Get().
		Index("index").
		Type("Type").
		Id("Id").
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Print(resp)
}