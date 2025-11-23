package dynamo

import (
	"fmt"

	"github.com/guregu/dynamo"
)

// Item is struct for DynamoDB
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

func Callxxxx() {
	table := DynamoInit()
	//////////////////////
	// 単純なCRUD - Create
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	err := table.Put(item).Run()
	if err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Read
	var readResult Item
	err = table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Update
	var updateResult Item
	text := "My Second Text"
	err = table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("MyText", text).Value(&updateResult)
	if err != nil {
		fmt.Printf("Failed to update item[%v]\n", err)
	}

	//////////////////////
	// Conditional Check
	err = table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).If("MyText = ?", "some word").Run()
	if err != nil {
		fmt.Printf("Failed to delete item with confitional check[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Delete
	err = table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run()
	if err != nil {
		fmt.Printf("Failed to delete item[%v]\n", err)
	}

}
