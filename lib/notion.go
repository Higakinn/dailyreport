package lib

import (
	"os"
	"fmt"
	"context"

	"github.com/dstotijn/go-notion"
)


func CreateNotionPage(databaseId string, pageName string) {
	apiKey := os.Getenv("NOTION_API_KEY")
	client := notion.NewClient(apiKey)
	
	queryDbParams :=  &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Name",
			Text: &notion.TextDatabaseQueryFilter{
				Equals: pageName,
			},
		},
	}
	resp, _:= client.QueryDatabase(context.Background(), databaseId ,queryDbParams)
	// 作ろうとしているページが存在しているか確認
        if len(resp.Results) == 1 {
		fmt.Println("existed page")
		os.Exit(1)
	}
	createPageParams := notion.CreatePageParams{
		ParentType: notion.ParentTypeDatabase,
		ParentID:   databaseId,
		DatabasePageProperties: &notion.DatabasePageProperties{
			"title": notion.DatabasePageProperty{
				Title: []notion.RichText{
					{
						Text: &notion.Text{
							Content: pageName,
						},
					},
				},
			},
		},
	}
	page, err := client.CreatePage(context.Background(), createPageParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(page)
	fmt.Println(databaseId)
	fmt.Println(client)
}
