package main

import (
	_ "embed"
	"math/rand"
	"time"

	"github.com/duc-cnzj/mars/v4/cmd"

	_ "github.com/duc-cnzj/mars/v4/plugins/domainmanager"
	_ "github.com/duc-cnzj/mars/v4/plugins/gitserver/github"
	_ "github.com/duc-cnzj/mars/v4/plugins/gitserver/gitlab"
	_ "github.com/duc-cnzj/mars/v4/plugins/picture"
	_ "github.com/duc-cnzj/mars/v4/plugins/wssender/memory"
	_ "github.com/duc-cnzj/mars/v4/plugins/wssender/nsq"
	_ "github.com/duc-cnzj/mars/v4/plugins/wssender/redis"
)

//go:embed config_example.yaml
var configFile []byte

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute(configFile)
}
