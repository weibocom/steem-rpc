package service

import (
	"strconv"

	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
)

func AddPost(company string, uid int64, mid int64, title string, content string, currentTs int64) (model.DNA, error) {
	author := generateUniqueAccount(company, uid)
	_, err := ipcClient.LookupAccount(author)
	if err == store.ErrNonExist {
		_, err = ipcClient.CreateAccount(author, `{"company":"`+company+`"}`)
		if err != nil {
			return nil, err
		}
	}

	return ipcClient.Post(author, mid, title, []byte(content), strconv.FormatInt(mid, 10), []string{"test"})
}

func GetContentByMsgID(company string, uid int64, mid int64) (*model.Post, error) {
	author := generateUniqueAccount(company, uid)
	return ipcClient.LookupPostByMsgID(author, mid)
}

func GetContentByDNA(dna string) (*model.Post, error) {
	return ipcClient.LookupPostByDNA(model.DNA(dna))
}

func GetLatestMid() (*model.Post, error) {
	return ipcClient.GetLatestPost()
}