package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-kms-wrapping/wrappers/awskms"
)

func main() {
	start := time.Now()

	runs := 0
	for i := 0; i < 1000; i++ {
		_, err := encript("test")
		if err != nil {
			fmt.Printf("error: %v", err)

		}
		runs += i
	}

	// } else {
	// 	fmt.Printf("The result after encript and decript is: %s", text)
	// }

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

}

func encript(data string) (string, error) {
	// Context used in this library is passed to various underlying provider
	// libraries; how it's used is dependent on the provider libraries
	ctx := context.Background()

	wrapper := awskms.NewWrapper(nil)
	_, err := wrapper.SetConfig(map[string]string{"kms_key_id": "eb337768-42c4-41d7-a6ae-6d8da3426e7a"})
	if err != nil {
		return "error", err
	}
	blobInfo, err := wrapper.Encrypt(ctx, []byte(data), nil)
	if err != nil {
		return "error", err
	}

	//
	// Do some things...
	//

	plaintext, err := wrapper.Decrypt(ctx, blobInfo, nil)
	if err != nil {
		return "error", err
	}

	return string(plaintext), nil

}
