package main

import (
	"fmt"
	"os"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

func main() {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("MyFirstTable")

	// Create
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	if err := table.Put(item).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	// Read
	var readResult Item
	if err := table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&readResult); err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
	}

	// Update
	var updateResult Item
	text := "My Second Text"
	if err := table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("MyText", text).Value(&updateResult); err != nil {
		fmt.Printf("Filed to update item[%v]\n", err)
	}

	// Conditional Check
	if err := table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).If("MyText = ?", "some word").Run(); err != nil {
		fmt.Printf("failed to delete item with conditional check[%v]", err)
	}

	// Delete
	if err := table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run(); err != nil {
		fmt.Printf("Failed to delete item[%v]", err)
	}

}
