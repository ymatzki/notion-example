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

	// Create new page
	newBlock, newBlockOp := client.SetNewRecordOp(userResp.User.ID, root, notionapi.BlockPage)
	ops := []*notionapi.Operation{newBlockOp}
	ops = append(ops, newBlock.SetTitleOp(fmt.Sprintf("new page %d", now())))
	ops = append(ops, root.ListAfterContentOp(newBlock.ID, ""))

	// Add text in new page
	childBlock, childBlockOp := client.SetNewRecordOp(userResp.User.ID, newBlock, notionapi.BlockText)
	ops = append(ops, childBlockOp)
	ops = append(ops, childBlock.SetTitleOp(fmt.Sprintf("new text %d", now())))
	ops = append(ops, newBlock.ListAfterContentOp(childBlock.ID, ""))

	err = client.SubmitTransaction(ops)
	if err != nil {
		panic(err)
	}
}

func now() int64 {
	return time.Now().Unix() * 1000
}
