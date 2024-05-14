package consumer

import "fmt"

func ReceiveAndInsertActivity(message []byte) {
	fmt.Println(string(message))
}