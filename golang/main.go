package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kjk/notionapi"
)

func main() {
	authToken, exists := os.LookupEnv("NOTION_AUTH_TOKEN")
	if !exists {
		panic(errors.New("must be set NOTION_AUTH_TOKEN environment variables"))
	}
	client := &notionapi.Client{
		AuthToken: authToken,
		Logger:    os.Stdout,
		DebugLog:  true,
	}
	pageID, exists := os.LookupEnv("NOTION_PAGE_ID")
	if !exists {
		panic(errors.New("must be set NOTION_PAGE_ID environment variables"))
	}

	page, err := client.DownloadPage(pageID)
	if err != err {
		panic(err)
	}

	userResp, err := client.LoadUserContent()
	if err != nil {
		panic(err)
	}

	root := page.Root()

	newBlock, err := addBlock(client, root, userResp.User.ID, notionapi.BlockPage, fmt.Sprintf("new page %d", now()))
	if err != nil {
		panic(err)
	}
	_, err = addBlock(client, newBlock, userResp.User.ID, notionapi.BlockText, fmt.Sprintf("new text %d", now()))
	if err != nil {
		panic(err)
	}
}

func now() int64 {
	return time.Now().Unix() * 1000
}

func addBlock(client *notionapi.Client, block *notionapi.Block, userID, recordType, title string) (*notionapi.Block, error) {
	newBlock, newBlockOp := client.SetNewRecordOp(userID, block, recordType)
	ops := []*notionapi.Operation{newBlockOp}
	ops = append(ops, newBlockOp)
	ops = append(ops, newBlock.SetTitleOp(title))
	ops = append(ops, block.ListAfterContentOp(newBlock.ID, ""))
	err := client.SubmitTransaction(ops)
	if err != nil {
		return nil, err
	}
	return newBlock, nil
}
