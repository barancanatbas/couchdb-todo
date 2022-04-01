package config

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func Connect() *gocb.Cluster {
	cluster, err := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "123456",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
	return cluster
}
