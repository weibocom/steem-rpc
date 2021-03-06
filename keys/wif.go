package keys

import (
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type WIF btcutil.WIF

func (w *WIF) MarshalJSON() ([]byte, error) {
	return []byte(`"` + w.String() + `"`), nil
}

func (w *WIF) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "" {
		return nil
	}
	s = strings.Trim(s, "\"")
	wif, err := DecodeWIF(s)
	if err != nil {
		return err
	}
	*w = *wif
	return nil
}

func (w *WIF) PrivateKey() *PrivateKey {
	return (*PrivateKey)(w.PrivKey)
}

func (w *WIF) PublicKey() *PublicKey {
	return (*PublicKey)(w.PrivKey.PubKey())
}

// Serialize can be used to turn WIF into a raw private key (32 bytes).
func (w *WIF) Serialize() []byte {
	return w.PrivKey.Serialize()
}

func (w *WIF) String() string {
	return ((*btcutil.WIF)(w)).String()
}

func DecodeWIF(wif string) (*WIF, error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return nil, err
	}
	return (*WIF)(w), nil
}

func NewWIF(pk *PrivateKey) (*WIF, error) {
	w, err := btcutil.NewWIF((*btcec.PrivateKey)(pk), &chaincfg.MainNetParams, false)
	if err != nil {
		return nil, err
	}
	return (*WIF)(w), nil
}

func GenerateWIF() (*WIF, error) {
	pk, err := GenerateKey()
	if err != nil {
		return nil, err
	}
	return NewWIF(pk)
}

func ParseSignature(sigStr []byte) (*btcec.Signature, error) {
	return btcec.ParseSignature(sigStr, btcec.S256())
}
