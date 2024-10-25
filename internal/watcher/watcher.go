package watcher

import (
	"context"
	"fmt"
	deployment "github.com/terrapi-solution/protocol/deployment/v1"
	"github.com/terrapi-solution/runner/internal/client"
	"github.com/terrapi-solution/runner/internal/git"
	"github.com/terrapi-solution/runner/internal/wrapper"
)

func Start(ctx context.Context) {
	c := client.NewClient()
	defer c.Close()

	d, _ := c.Deployment.Get(ctx, &deployment.GetRequest{Id: 1})
	fmt.Println(d.Request)

	gitClient := git.NewGitClient()
	gitClient.PrepareClient(d.Module.Address, "master")
	_, err := gitClient.Clone("c:/temp/terraform")
	if err != nil {
		fmt.Println(err)
	}
	defer gitClient.Remove("c:/temp/terraform")

	executor := wrapper.New("terraform")
	executorParam := wrapper.NewInitParams()

	executor.SetWorkingDirectory("c:/temp/terraform")
	t := executor.Init(executorParam).Initialise()
	t.InitLogger(&wrapper.OutputLog{})

	err = t.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(t.Dir)
}

func Processing(identifier uint, ctx context.Context) {

}
