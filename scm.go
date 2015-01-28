package scm

import (
	"io"

	"github.com/peter-edge/exec"
)

type GitCheckoutOptions struct {
	User            string
	Host            string
	Path            string
	Branch          string
	CommitId        string
	SecurityOptions *GitSecurityOptions
}

type GithubCheckoutOptions struct {
	User            string
	Repository      string
	Branch          string
	CommitId        string
	SecurityOptions *GithubSecurityOptions
}

type HgCheckoutOptions struct {
	User            string
	Host            string
	Path            string
	ChangesetId     string
	SecurityOptions *HgSecurityOptions
}

type ClientOptions struct {
	IgnoreCheckoutFiles bool
}

type Client interface {
	CheckoutGitTarball(*GitCheckoutOptions) (io.Reader, error)
	CheckoutGithubTarball(*GithubCheckoutOptions) (io.Reader, error)
	CheckoutHgTarball(*HgCheckoutOptions) (io.Reader, error)
}

func NewClient(executorReadFileManagerProvider exec.ExecutorReadFileManagerProvider, clientOptions *ClientOptions) Client {
	return newClient(executorReadFileManagerProvider, clientOptions)
}
