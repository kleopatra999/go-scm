package scm

// Generated by gen-enumtype. DO NOT EDIT.

import "fmt"

type SecurityOptionsType uint

var SecurityOptionsTypeSsh SecurityOptionsType = 0
var SecurityOptionsTypeAccessToken SecurityOptionsType = 1

var securityOptionsTypeToString = map[SecurityOptionsType]string{
	SecurityOptionsTypeSsh: "ssh",
	SecurityOptionsTypeAccessToken: "accessToken",
}

var stringToSecurityOptionsType = map[string]SecurityOptionsType{
	"ssh": SecurityOptionsTypeSsh,
	"accessToken": SecurityOptionsTypeAccessToken,
}

func AllSecurityOptionsTypes() []SecurityOptionsType {
	return []SecurityOptionsType{
		SecurityOptionsTypeSsh,
		SecurityOptionsTypeAccessToken,
	}
}

func SecurityOptionsTypeOf(s string) (SecurityOptionsType, error) {
	securityOptionsType, ok := stringToSecurityOptionsType[s]
	if !ok {
		return 0, newErrorUnknownSecurityOptionsType(s)
	}
	return securityOptionsType, nil
}

func (this SecurityOptionsType) String() string {
	if int(this) < len(securityOptionsTypeToString) {
		 return securityOptionsTypeToString[this]
	}
	panic(newErrorUnknownSecurityOptionsType(this).Error())
}

type SecurityOptions interface {
	Type() SecurityOptionsType
}

func (this *SSHSecurityOptions) Type() SecurityOptionsType {
	return SecurityOptionsTypeSsh
}

func (this *AccessTokenSecurityOptions) Type() SecurityOptionsType {
	return SecurityOptionsTypeAccessToken
}

func SecurityOptionsSwitch(
	securityOptions SecurityOptions,
	sSHSecurityOptionsFunc func(sSHSecurityOptions *SSHSecurityOptions) error,
	accessTokenSecurityOptionsFunc func(accessTokenSecurityOptions *AccessTokenSecurityOptions) error,
) error {
	switch securityOptions.Type() {
	case SecurityOptionsTypeSsh:
		return sSHSecurityOptionsFunc(securityOptions.(*SSHSecurityOptions))
	case SecurityOptionsTypeAccessToken:
		return accessTokenSecurityOptionsFunc(securityOptions.(*AccessTokenSecurityOptions))
	default:
		return newErrorUnknownSecurityOptionsType(securityOptions.Type())
	}
}

func (this SecurityOptionsType) NewSecurityOptions(
	sSHSecurityOptionsFunc func() (*SSHSecurityOptions, error),
	accessTokenSecurityOptionsFunc func() (*AccessTokenSecurityOptions, error),
) (SecurityOptions, error) {
	switch this {
	case SecurityOptionsTypeSsh:
		return sSHSecurityOptionsFunc()
	case SecurityOptionsTypeAccessToken:
		return accessTokenSecurityOptionsFunc()
	default:
		return nil, newErrorUnknownSecurityOptionsType(this)
	}
}

func (this SecurityOptionsType) Produce(
	securityOptionsTypeSshFunc func() (interface{}, error),
	securityOptionsTypeAccessTokenFunc func() (interface{}, error),
) (interface{}, error) {
	switch this {
	case SecurityOptionsTypeSsh:
		return securityOptionsTypeSshFunc()
	case SecurityOptionsTypeAccessToken:
		return securityOptionsTypeAccessTokenFunc()
	default:
		return nil, newErrorUnknownSecurityOptionsType(this)
	}
}

func (this SecurityOptionsType) Handle(
	securityOptionsTypeSshFunc func() error,
	securityOptionsTypeAccessTokenFunc func() error,
) error {
	switch this {
	case SecurityOptionsTypeSsh:
		return securityOptionsTypeSshFunc()
	case SecurityOptionsTypeAccessToken:
		return securityOptionsTypeAccessTokenFunc()
	default:
		return newErrorUnknownSecurityOptionsType(this)
	}
}

func newErrorUnknownSecurityOptionsType(value interface{}) error {
	return fmt.Errorf("scm: UnknownSecurityOptionsType: %v", value)
}
type CheckoutOptionsType uint

var CheckoutOptionsTypeGit CheckoutOptionsType = 0
var CheckoutOptionsTypeGithub CheckoutOptionsType = 1
var CheckoutOptionsTypeHg CheckoutOptionsType = 2
var CheckoutOptionsTypeBitbucketGit CheckoutOptionsType = 3
var CheckoutOptionsTypeBitbucketHg CheckoutOptionsType = 4

var checkoutOptionsTypeToString = map[CheckoutOptionsType]string{
	CheckoutOptionsTypeGit: "git",
	CheckoutOptionsTypeGithub: "github",
	CheckoutOptionsTypeHg: "hg",
	CheckoutOptionsTypeBitbucketGit: "bitbucketGit",
	CheckoutOptionsTypeBitbucketHg: "bitbucketHg",
}

var stringToCheckoutOptionsType = map[string]CheckoutOptionsType{
	"git": CheckoutOptionsTypeGit,
	"github": CheckoutOptionsTypeGithub,
	"hg": CheckoutOptionsTypeHg,
	"bitbucketGit": CheckoutOptionsTypeBitbucketGit,
	"bitbucketHg": CheckoutOptionsTypeBitbucketHg,
}

func AllCheckoutOptionsTypes() []CheckoutOptionsType {
	return []CheckoutOptionsType{
		CheckoutOptionsTypeGit,
		CheckoutOptionsTypeGithub,
		CheckoutOptionsTypeHg,
		CheckoutOptionsTypeBitbucketGit,
		CheckoutOptionsTypeBitbucketHg,
	}
}

func CheckoutOptionsTypeOf(s string) (CheckoutOptionsType, error) {
	checkoutOptionsType, ok := stringToCheckoutOptionsType[s]
	if !ok {
		return 0, newErrorUnknownCheckoutOptionsType(s)
	}
	return checkoutOptionsType, nil
}

func (this CheckoutOptionsType) String() string {
	if int(this) < len(checkoutOptionsTypeToString) {
		 return checkoutOptionsTypeToString[this]
	}
	panic(newErrorUnknownCheckoutOptionsType(this).Error())
}

type CheckoutOptions interface {
	Type() CheckoutOptionsType
}

func (this *GitCheckoutOptions) Type() CheckoutOptionsType {
	return CheckoutOptionsTypeGit
}

func (this *GithubCheckoutOptions) Type() CheckoutOptionsType {
	return CheckoutOptionsTypeGithub
}

func (this *HgCheckoutOptions) Type() CheckoutOptionsType {
	return CheckoutOptionsTypeHg
}

func (this *BitbucketGitCheckoutOptions) Type() CheckoutOptionsType {
	return CheckoutOptionsTypeBitbucketGit
}

func (this *BitbucketHgCheckoutOptions) Type() CheckoutOptionsType {
	return CheckoutOptionsTypeBitbucketHg
}

func CheckoutOptionsSwitch(
	checkoutOptions CheckoutOptions,
	gitCheckoutOptionsFunc func(gitCheckoutOptions *GitCheckoutOptions) error,
	githubCheckoutOptionsFunc func(githubCheckoutOptions *GithubCheckoutOptions) error,
	hgCheckoutOptionsFunc func(hgCheckoutOptions *HgCheckoutOptions) error,
	bitbucketGitCheckoutOptionsFunc func(bitbucketGitCheckoutOptions *BitbucketGitCheckoutOptions) error,
	bitbucketHgCheckoutOptionsFunc func(bitbucketHgCheckoutOptions *BitbucketHgCheckoutOptions) error,
) error {
	switch checkoutOptions.Type() {
	case CheckoutOptionsTypeGit:
		return gitCheckoutOptionsFunc(checkoutOptions.(*GitCheckoutOptions))
	case CheckoutOptionsTypeGithub:
		return githubCheckoutOptionsFunc(checkoutOptions.(*GithubCheckoutOptions))
	case CheckoutOptionsTypeHg:
		return hgCheckoutOptionsFunc(checkoutOptions.(*HgCheckoutOptions))
	case CheckoutOptionsTypeBitbucketGit:
		return bitbucketGitCheckoutOptionsFunc(checkoutOptions.(*BitbucketGitCheckoutOptions))
	case CheckoutOptionsTypeBitbucketHg:
		return bitbucketHgCheckoutOptionsFunc(checkoutOptions.(*BitbucketHgCheckoutOptions))
	default:
		return newErrorUnknownCheckoutOptionsType(checkoutOptions.Type())
	}
}

func (this CheckoutOptionsType) NewCheckoutOptions(
	gitCheckoutOptionsFunc func() (*GitCheckoutOptions, error),
	githubCheckoutOptionsFunc func() (*GithubCheckoutOptions, error),
	hgCheckoutOptionsFunc func() (*HgCheckoutOptions, error),
	bitbucketGitCheckoutOptionsFunc func() (*BitbucketGitCheckoutOptions, error),
	bitbucketHgCheckoutOptionsFunc func() (*BitbucketHgCheckoutOptions, error),
) (CheckoutOptions, error) {
	switch this {
	case CheckoutOptionsTypeGit:
		return gitCheckoutOptionsFunc()
	case CheckoutOptionsTypeGithub:
		return githubCheckoutOptionsFunc()
	case CheckoutOptionsTypeHg:
		return hgCheckoutOptionsFunc()
	case CheckoutOptionsTypeBitbucketGit:
		return bitbucketGitCheckoutOptionsFunc()
	case CheckoutOptionsTypeBitbucketHg:
		return bitbucketHgCheckoutOptionsFunc()
	default:
		return nil, newErrorUnknownCheckoutOptionsType(this)
	}
}

func (this CheckoutOptionsType) Produce(
	checkoutOptionsTypeGitFunc func() (interface{}, error),
	checkoutOptionsTypeGithubFunc func() (interface{}, error),
	checkoutOptionsTypeHgFunc func() (interface{}, error),
	checkoutOptionsTypeBitbucketGitFunc func() (interface{}, error),
	checkoutOptionsTypeBitbucketHgFunc func() (interface{}, error),
) (interface{}, error) {
	switch this {
	case CheckoutOptionsTypeGit:
		return checkoutOptionsTypeGitFunc()
	case CheckoutOptionsTypeGithub:
		return checkoutOptionsTypeGithubFunc()
	case CheckoutOptionsTypeHg:
		return checkoutOptionsTypeHgFunc()
	case CheckoutOptionsTypeBitbucketGit:
		return checkoutOptionsTypeBitbucketGitFunc()
	case CheckoutOptionsTypeBitbucketHg:
		return checkoutOptionsTypeBitbucketHgFunc()
	default:
		return nil, newErrorUnknownCheckoutOptionsType(this)
	}
}

func (this CheckoutOptionsType) Handle(
	checkoutOptionsTypeGitFunc func() error,
	checkoutOptionsTypeGithubFunc func() error,
	checkoutOptionsTypeHgFunc func() error,
	checkoutOptionsTypeBitbucketGitFunc func() error,
	checkoutOptionsTypeBitbucketHgFunc func() error,
) error {
	switch this {
	case CheckoutOptionsTypeGit:
		return checkoutOptionsTypeGitFunc()
	case CheckoutOptionsTypeGithub:
		return checkoutOptionsTypeGithubFunc()
	case CheckoutOptionsTypeHg:
		return checkoutOptionsTypeHgFunc()
	case CheckoutOptionsTypeBitbucketGit:
		return checkoutOptionsTypeBitbucketGitFunc()
	case CheckoutOptionsTypeBitbucketHg:
		return checkoutOptionsTypeBitbucketHgFunc()
	default:
		return newErrorUnknownCheckoutOptionsType(this)
	}
}

func newErrorUnknownCheckoutOptionsType(value interface{}) error {
	return fmt.Errorf("scm: UnknownCheckoutOptionsType: %v", value)
}
