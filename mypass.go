package mypass

import (
	"encoding/hex"
	"encoding/json"

	"github.com/emrekeskinmac/mypass/account"
	"github.com/emrekeskinmac/mypass/encyrpt"
)

type MyPass struct {
	provider Provider
	username string
	secret   []byte
}

func New(username, password string, provider Provider) *MyPass {
	mp := &MyPass{
		provider: provider,
		username: username,
	}
	mp.secret = account.MakeSecret(username, password)
	return mp
}

func (mp *MyPass) Save(account account.Account) error {
	data, err := json.Marshal(account)
	if err != nil {
		return err
	}
	encrypted, err := encyrpt.Encrypt(data, mp.secret)
	if err != nil {
		return err
	}
	toHex := hex.EncodeToString(encrypted)
	return mp.provider.Save(mp.username, account.Name, toHex)
}

func (mp *MyPass) Find(name string) (accounts []account.Account, err error) {
	hxs, err := mp.provider.Find(mp.username, name)
	if err != nil {
		return accounts, err
	}
	for _, hx := range hxs {
		data, err := hex.DecodeString(hx)
		if err != nil {
			return accounts, err
		}
		decrypted, err := encyrpt.Decrypt(data, mp.secret)
		if err != nil {
			return accounts, err
		}
		var account account.Account
		err = json.Unmarshal(decrypted, &account)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (mp *MyPass) Delete(name string) (accounts []account.Account, err error) {
	return accounts, err
}
