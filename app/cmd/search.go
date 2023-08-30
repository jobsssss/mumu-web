package cmd

import (
	json2 "encoding/json"
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"mumu/app/models/book"
	"mumu/pkg/config"
	"mumu/pkg/console"
	"mumu/pkg/logger"
	"mumu/pkg/model"
)

var CmdSearch = &cobra.Command{
	Use:   "search",
	Short: "Search management",
}

var queryKey string
var CmdSearchQuery = &cobra.Command{
	Use:   "query",
	Short: "Query by arg",
	Run:   query,
}

var CmdSearchAdd = &cobra.Command{
	Use:   "add",
	Short: "Add document",
	Run:   add,
}

var delIndex string
var CmdSearchDel = &cobra.Command{
	Use:   "del",
	Short: "del meili search index",
	Run:   del,
}

func init() {
	CmdSearch.AddCommand(CmdSearchQuery, CmdSearchAdd, CmdSearchDel)

	CmdSearchQuery.Flags().StringVarP(&queryKey, "key", "k", "", "The key of search")
	err := CmdSearchQuery.MarkFlagRequired("key")
	logger.LogIf(err)

	CmdSearchDel.Flags().StringVarP(&delIndex, "key", "k", "", "The index will delete")
	err2 := CmdSearchDel.MarkFlagRequired("key")
	logger.LogIf(err2)
}

var client = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   cast.ToString(config.Env("MEILISEARCH_HOST", "")),
	APIKey: cast.ToString(config.Env("MEILISEARCH_KEY", "")),
})

func query(cmd *cobra.Command, args []string) {
	response, err := client.Index("book").Search(queryKey, &meilisearch.SearchRequest{})
	if err != nil {
		return
	}
	fmt.Println(response.Hits)
}

func add(cmd *cobra.Command, args []string) {
	var results []book.Book
	result := model.DB.FindInBatches(&results, 10, func(tx *gorm.DB, batch int) error {
		for _, result := range results {
			json, _ := json2.Marshal(result)
			_, err := client.Index("book").AddDocuments(json)
			logger.LogIf(err)
			console.Success(result.Name + "同步到搜索引擎成功")
		}
		return nil
	})
	console.Success("All sync success, total:" + cast.ToString(result.RowsAffected) + " done!")
}

func del(cmd *cobra.Command, args []string) {
	client.DeleteIndex(delIndex)
}
