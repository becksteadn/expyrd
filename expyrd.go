package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"time"
)

type Key struct {
	UserName string
	AccessKeyID string
	CreateDate time.Time
	LastAccessedDate time.Time
}

func handle(err) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create AWS SDK session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := iam.New(sess)

	keys := make(map([string][]Key))

	// Get all users
	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})
	handle(err)

	for _, user := range result.Users {
		if user.PasswordLastUsed != nil {
			fmt.Println(*user.UserName, user.PasswordLastUsed)
		} else {
			keys, err := svc.ListAccessKeys(&iam.ListAccessKeysInput{
				UserName: aws.String(*user.UserName),
			})
			handle(err)

			fmt.Println(*user.UserName, keys.AccessKeyMetadata[0].CreateDate)
			fmt.Printf("%T\n", keys.AccessKeyMetadata[0].CreateDate)
		}
	} 
}