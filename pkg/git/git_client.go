package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Client struct {
	options *git.CloneOptions
}

// Creates a new instance of gitClient.
func NewGitClient() *Client {
	return &Client{}
}

// Initializes the gitClient with the specified repository path and branch.
func (c *Client) PrepareClient(path string, branch string) {
	c.options = &git.CloneOptions{
		URL:           path,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
	}
}

// Clone a Git repository to the specified path.
func (c *Client) Clone(path string) (*git.Repository, error) {
	repo, err := git.PlainClone(path, false, c.options)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
