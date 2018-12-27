package register

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// ZOOKSERVER zookeeper server
var ZOOKSERVER = os.Getenv("ZOOK_SERVER")

// DataServer register ip:port
func DataServer(key string) error {
	conn, _, err := zk.Connect([]string{ZOOKSERVER}, time.Second)
	if err != nil {
		return err
	}

	exist, _, err := conn.Exists("/objects")
	if err != nil {
		return err
	}

	if !exist {
		_, err = conn.Create("/objects", []byte{}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}

		_, err = conn.Create("/objects/dataservers", []byte{}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}
	}
	_, err = conn.Create(fmt.Sprintf("/objects/dataservers/%s", key), []byte{}, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}

	return nil
}
