package register

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var zkConn *zk.Conn

// Prepare connect to zk
func Prepare() {
	conn, _, err := zk.Connect([]string{ZOOKSERVER}, time.Second)
	if err == nil {
		zkConn = conn
	} else {
		panic(err)
	}
}

// RandomSelectDataServer random policy
func RandomSelectDataServer() (string, error) {
	if zkConn == nil {
		return "", fmt.Errorf("zk conn lose")
	}
	instances, _, err := zkConn.Children("/objects/dataservers")
	if err != nil {
		return "", err
	}
	return instances[rand.Intn(len(instances))], nil
}
