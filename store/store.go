package store

import (
	"errors"
	"strings"

	"github.com/weibocom/ipc/model"
)

var (
	ErrNonExist       = errors.New("not exist")
	ErrExist          = errors.New("already exist")
	ErrNotImplemented = errors.New("not implemented yet")
)

type Store interface {
	Account
	Post
	Close() error
}

type Account interface {
	ExistAccount(name string) (bool, error)
	SaveAccount(a *model.Account) error
	LoadAccount(name string) (*model.Account, error)
	GetAccounts(company string, offset int, limit int) ([]*model.Account, error)
	GetAccountCount() (int, error)
}

type Post interface {
	GetPostCount() (int, error)
	ExistPost(dna model.DNA) (bool, error)
	SavePost(p *model.Post) error
	LoadPost(dna model.DNA) (*model.Post, error)
	GetLatestPost() (*model.Post, error)
	GetPostByMsgID(author string, mid int64) (*model.Post, error)
	GetPostByDNA(dna model.DNA) (*model.Post, error)
	GetPostByAuthor(author string, offset int, limit int) ([]*model.Post, error)
	LookupSimilarPosts(dna string, keywords string, offset int, limit int) ([]*model.Post, error)
	GetAccountPostCount(name string) (int, error)
}

func getCompany(name string) string {
	fields := strings.SplitN(name, "-", 2)
	if len(fields) == 2 {
		return fields[0]
	}

	return ""
}
