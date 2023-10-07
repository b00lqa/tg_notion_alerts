package main

import (
	"context"

	"github.com/jomei/notionapi"
	"github.com/spf13/viper"
)

type notionClient struct {
	config map[string]string
	client *notionapi.Client
}

func createNotionClient() notionClient {
	config := viper.GetStringMapString("Notion")
	client := notionapi.NewClient(
		notionapi.Token(config["key"]),
	)

	return notionClient{
		config: config,
		client: client,
	}
}

func (nc *notionClient) getDB() (*notionapi.Database, error) {
	db, err := nc.client.Database.Get(
		context.TODO(),
		notionapi.DatabaseID(nc.config["db_id"]),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (nc *notionClient) getTasksList(db *notionapi.Database, filter notionapi.Filter) (
	[]notionapi.Page,
	error,
) {
	qresp, err := nc.client.Database.Query(
		context.TODO(),
		notionapi.DatabaseID(db.ID),
		&notionapi.DatabaseQueryRequest{
			Filter: notionapi.PropertyFilter{},
		},
	)
	if err != nil {
		return []notionapi.Page{}, err
	}

	return qresp.Results, nil
}
