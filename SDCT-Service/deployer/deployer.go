package deployer

import (
	"crypto/ecdsa"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/client"
	"github.com/sdct/contracts/sdctsetup"
	"github.com/sdct/contracts/sdctsystem"
	"github.com/sdct/contracts/sdctverifier"
	"github.com/sdct/contracts/token"
	"github.com/sdct/contracts/tokenconverter"
	"github.com/sdct/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeploySDCTSetup deploys param contracts.
func DeploySDCTSetup(auth *bind.TransactOpts, ethclient *ethclient.Client, pk *ecdsa.PublicKey) (common.Address, *sdctsetup.Sdctsetup) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := sdctsetup.DeploySdctsetup(auth, ethclient, pk.X, pk.Y)
	if err != nil {
		panic(err)
	}
	log.Info("Send SDCTSetup tx succeeds", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployToken deploys an erc20 token contract.
func DeployToken(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *token.Token) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := token.DeployToken(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("Send token tx succeeds", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployTokenConverter deploys contract to convert token.
func DeployTokenConverter(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *tokenconverter.Tokenconverter) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := tokenconverter.DeployTokenconverter(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("Send token converter tx succeeds", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeploySDCTVerifier deploys sdct proof verifier contract.
func DeploySDCTVerifier(auth *bind.TransactOpts, ethclient *ethclient.Client, params common.Address) (common.Address, *sdctverifier.Sdctverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := sdctverifier.DeploySdctverifier(auth, ethclient, params)
	if err != nil {
		panic(err)
	}
	log.Info("Send SDCTVerifier tx succeeds", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeploySDCTSystem deploys sdct system main contract.
func DeploySDCTSystem(auth *bind.TransactOpts, ethclient *ethclient.Client, params, sdctVerifier, tokenConverter common.Address) (common.Address, *sdctsystem.Sdctsystem) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := sdctsystem.DeploySdctsystem(auth, ethclient, params, sdctVerifier, tokenConverter)
	if err != nil {
		panic(err)
	}
	log.Info("Send SDCTSystem tx succeeds", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeploySDCTSystemAllContract deploys all contract for sdct system.
func DeploySDCTSystemAllContract(auth *bind.TransactOpts, ethclient *ethclient.Client, pk *ecdsa.PublicKey) ([]common.Address, *sdctsystem.Sdctsystem) {
	addrs := make([]common.Address, 0)

	params, _ := DeploySDCTSetup(auth, ethclient, pk)
	addrs = append(addrs, params)

	tokenConverter, _ := DeployTokenConverter(auth, ethclient)
	addrs = append(addrs, tokenConverter)

	verifierAddr, _ := DeploySDCTVerifier(auth, ethclient, params)
	addrs = append(addrs, verifierAddr)

	sdctMain, sdct := DeploySDCTSystem(auth, ethclient, params, verifierAddr, tokenConverter)
	addrs = append(addrs, sdctMain)

	return addrs, sdct
}

// InitVector inits g/h generate vector for agg range proof.
func InitVector(auth *bind.TransactOpts, ethclient *ethclient.Client, addr common.Address, bitsize int) {
	vectorsize := bitsize * 2
	initStep := 32
	step := 45
	if vectorsize <= initStep {
		return
	}
	vectorsize -= initStep
	vr, err := sdctverifier.NewSdctverifier(addr, ethclient)
	if err != nil {
		panic(err)
	}

	for vectorsize > 0 {
		tx, err := vr.Init(auth, big.NewInt(int64(step)))
		if err != nil {
			panic(err)
		}
		log.Info("Send SDCTVerifier tx succeeds", "tx", tx.Hash().Hex())
		client.WaitForTx(ethclient, tx.Hash())
		auth.Nonce.Add(auth.Nonce, utils.One)
		vectorsize -= step
	}
}
