package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	eth       = new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)
	amount    = new(big.Int).SetUint64(90)
	amountWei = new(big.Int).Mul(amount, eth)
)

const (
	local = "http://127.0.0.1:8545"
)

// Client represents rpc client.
type Client struct {
	client *rpc.Client
}

// Transaction represents a tx.
type Transaction struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Nonce    string         `json:"nonce,omitempty"`
	Value    string         `json:"value"`
	GasLimit string         `json:"gasLimit"`
	GasPrice string         `json:"gasPrice"`
	Data     []byte         `json:"data"`
}

// Dial connects to rpc client
func Dial(url string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	client := Client{}
	client.client = c

	return &client, nil
}

// GetAccounts returns all accounts.
func (c *Client) GetAccounts() ([]common.Address, error) {
	var res []common.Address
	err := c.client.Call(&res, "eth_accounts")

	return res, err
}

// SendTx sends tx without signing it using send_transaction.
func (c *Client) SendTx(tx *Transaction) (common.Hash, error) {
	var res common.Hash

	err := c.client.Call(&res, "eth_sendTransaction", tx)

	return res, err
}

// SendFromFirstAccout sends eth to someone from test local account get by rpc.
func (c *Client) SendFromFirstAccout(to common.Address, amount *big.Int) (common.Hash, error) {
	accounts, err := c.GetAccounts()
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendETH(accounts[0], to, amount)
}

// SendETH sends ether to address.
func (c *Client) SendETH(from common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	ether := new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)
	tx := Transaction{}
	tx.From = from
	tx.To = to
	tx.Value = new(big.Int).Mul(amount, ether).String()
	tx.GasLimit = "100000"

	return c.SendTx(&tx)
}

// GetAccountWithETH sends eth to an random account.
func (c *Client) GetAccountWithETH() *bind.TransactOpts {
	account := GenerateAccount()

	if _, err := c.SendFromFirstAccout(account.From, amount); err != nil {
		panic(err)
	}

	return account
}

// GetClient returns ethclient, otherwise panic.
func GetClient(url string) *ethclient.Client {
	return getClient(url)
}

// GetLocal return client on local net.
func GetLocal() *ethclient.Client {
	return getClient(local)
}

// GetLocalRPC returns rpc client on local net.
func GetLocalRPC() *Client {
	client, err := Dial(local)
	if err != nil {
		panic(err)
	}

	return client
}

func getClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return client
}

// GetAccountWithKey returns opts for sending tx.
func GetAccountWithKey(keyhex string) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(keyhex)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(key)
}

// GenerateAccount generates a new random account.
func GenerateAccount() *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return bind.NewKeyedTransactor(key)
}

// SetNonce sets nonce to account if nil.
func SetNonce(auth *bind.TransactOpts, client *ethclient.Client) error {
	if auth.Nonce != nil {
		return nil
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return err
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	return nil
}

// WaitForTx stacks until tx is mined.
func WaitForTx(client *ethclient.Client, hash common.Hash) *types.Receipt {
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime) > time.Minute*5 {
			panic(fmt.Sprintf("wait too much time: tx %s", hash.Hex()))
		}
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		if receipt != nil && receipt.Status != 1 {
			panic("tx failed")
		}

		if receipt != nil && receipt.Status == 1 {
			return receipt
		}

		time.Sleep(time.Millisecond * 500)
	}
}
