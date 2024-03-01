package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"runtime"
	"time"
)

var subfix = "888"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go a(i)
	}

	time.Sleep(time.Hour * 666)
}

func generateEthWallet() (*ecdsa.PrivateKey, common.Address, error) {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}

	// 从私钥中获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	// 从公钥中获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, address, nil
}

func a(i int) {

	for {
		privateKey, address, err := generateEthWallet()
		if err != nil {
			fmt.Println(err)
		}

		if is777777Address(address, fmt.Sprintf("%d", i)) {
			fmt.Printf("Private Key: %x\n", crypto.FromECDSA(privateKey))
			fmt.Printf("Address: %s\n", address.Hex())
			panic("生成好了")
			return
		}
	}

}

func is777777Address(address common.Address, a string) bool {
	return address.Hex()[len(address.Hex())-len(subfix):] == subfix
}
