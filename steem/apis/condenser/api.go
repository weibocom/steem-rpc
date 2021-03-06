package condenser

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/weibocom/ipc/interfaces"
	"github.com/weibocom/ipc/internal/call"
	"github.com/weibocom/ipc/steem/types"
)

const (
	APIID = "condenser_api"
)

var emptyParams = []string{}

type API struct {
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) *API {
	return &API{caller}
}

/*
   // Accounts
   (get_accounts)
   (get_account_references)
   (lookup_account_names)
   (lookup_accounts)
   (get_account_count)
   (get_conversion_requests)
   (get_account_history)
*/

func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	raw, err := call.Raw(api.caller, APIID+".get_accounts", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) LookupAccountNames(accountNames []string) ([]*Account, error) {
	raw, err := call.Raw(api.caller, APIID+".lookup_account_names", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	raw, err := call.Raw(api.caller, APIID+".lookup_accounts", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetAccountCount() (uint32, error) {
	raw, err := call.Raw(api.caller, APIID+".get_account_count", emptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, err
	}
	return resp, nil
}

func (api *API) GetAccountHistory(account string, from int64, limit uint32) ([]*types.OperationObject, error) {
	raw, err := call.Raw(api.caller, APIID+".get_account_history", []interface{}{account, from, limit})
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(*raw), &tmp1); err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errm
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

/*
   // Witnesses
   (get_witnesses)
   (get_witness_by_account)
   (get_witnesses_by_vote)
   (lookup_witness_accounts)
   (get_witness_count)
   (get_active_witnesses)
   (get_miner_queue)
*/

// GetWitnesses implements get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witnesses", [][]uint32{id})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetLegacyWitnesses(id []uint32) ([]*Witness, error) {
	var resp []*Witness
	err := api.caller.Call("get_witnesses", [][]uint32{id}, &resp)

	return resp, err
}

func (api *API) GetWitnessesByAccount(author string) (*Witness, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witness_by_account", []string{author})
	if err != nil {
		return nil, err
	}
	var resp *Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	raw, err := call.Raw(api.caller, APIID+".lookup_witness_accounts", []string{author})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witness_count", emptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, err
	}
	return resp, nil
}

func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := call.Raw(api.caller, APIID+".get_active_witnesses", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetFeedEntriesRaw(
	accountName string,
	entryID uint32,
	limit uint16,
) (*json.RawMessage, error) {

	if limit == 0 {
		limit = 500
	}

	params := []interface{}{accountName, entryID, limit}

	raw, err := call.Raw(api.caller, APIID+".get_feed_entries", params)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (api *API) GetFeedEntries(
	accountName string,
	entryID uint32,
	limit uint16,
) ([]*FeedEntry, error) {

	raw, err := api.GetFeedEntriesRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}

	var resp []*FeedEntry
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(
			err, "go-steem/rpc: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

func (api *API) GetContent(author, permlink string) (*Content, error) {
	var resp Content
	if err := api.caller.Call(APIID+".get_content", []string{author, permlink}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
