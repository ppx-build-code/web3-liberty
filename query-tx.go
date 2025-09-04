package main

import (
    "context"
    "fmt"
    "log"
	// "encoding/json"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

func test() {
	fmt.Printf("end....")
}


func main() {

	fmt.Println("开始了。。。")

    // 连接以太坊节点（可以是 Infura / Alchemy / 本地 geth）
    client, err := ethclient.Dial("http://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }

	// // 获取最新区块号
    // header, err := client.HeaderByNumber(context.Background(), nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println("最新区块号:", header.Number.String())

	// // 获取整个区块
    // block, err := client.BlockByNumber(context.Background(), header.Number)
    // if err != nil {
    //     log.Fatal(err)
    // }

	// fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
    // fmt.Printf("交易数量: %d\n", len(block.Transactions()))

	// // 遍历区块内交易
    // chainID, _ := client.NetworkID(context.Background())
    // for _, tx := range block.Transactions() {
    //     from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
    //     if err != nil {
    //         log.Println("无法解析发送者:", err)
    //         continue
    //     }

    //     txInfo := map[string]interface{}{
    //         "From":     from.Hex(),
    //         "To":       tx.To().Hex(),
    //         "Value":    tx.Value().String(),
    //         "Gas":      tx.Gas(),
    //         "GasPrice": tx.GasPrice().String(),
    //         "Hash":     tx.Hash().Hex(),
    //     }

    //     jsonData, _ := json.Marshal(txInfo)
    //     fmt.Println(string(jsonData))
    // }

    // 交易哈希
    txHash := common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")

    // 查询交易
    tx,pending, err := client.TransactionByHash(context.Background(), txHash)
    if err != nil {
        log.Fatal(err)
    }

	// 获取 chain ID（必须）
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // 使用 EIP155Signer 或 HomesteadSigner 恢复发送者
    signer := types.NewEIP155Signer(chainID)
    from, err := types.Sender(signer, tx)
    if err != nil {
        log.Fatal(err)
    }

	
    fmt.Println("From:", from.Hex(), pending)
    fmt.Println("To:", tx.To().Hex())
    fmt.Println("Value:", tx.Value())

    // fmt.Println("交易信息:")
    // fmt.Println("Hash:", tx.Hash().Hex())
    // fmt.Println("Value:", tx.Value().String())
    // fmt.Println("From:", tx.From()) // 注意: 需要签名者信息才能获取
    // fmt.Println("To:", tx.To().Hex())
    // fmt.Println("Gas:", tx.Gas())
    // fmt.Println("GasPrice:", tx.GasPrice().String())
    // fmt.Println("是否Pending:", isPending)
}
