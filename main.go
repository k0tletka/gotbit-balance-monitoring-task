package main

import (
    "os/signal"
    "context"
    "syscall"
    "log"
    "time"

    "github.com/k0tletka/gotbit-balance-monitoring-task/tether_contract"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
    ethMainnetEndpoint = "https://speedy-nodes-nyc.moralis.io/ef1d4b780ef0a230b52eeb0a/eth/mainnet"
    outputInterval = 30 * time.Second
)

var (
    contractAddress = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
    userAddress = common.HexToAddress("0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852")
)

func main() {
    ctx, _ := signal.NotifyContext(context.Background(),
        syscall.SIGINT,
        syscall.SIGTERM,
        syscall.SIGQUIT,
    )

    timeoutCtx, _ := context.WithTimeout(ctx, 10 * time.Second)

    client, err := ethclient.DialContext(timeoutCtx, ethMainnetEndpoint)
    if err != nil {
        log.Fatalln(err)
    }

    contractCaller, err := tether_contract.NewTetherContract(contractAddress, client)
    if err != nil {
        log.Fatalln(err)
    }

    outputTimer := time.NewTicker(outputInterval)

    for {
        timeoutCtx, _ = context.WithTimeout(ctx, 10 * time.Second)

        if balance, err := contractCaller.BalanceOf(&bind.CallOpts{Context: timeoutCtx}, userAddress); err == nil {
            log.Println("Balance of account: ", balance)
        } else {
            log.Fatalln(err)
        }

        select {
        case <-outputTimer.C:
            continue
        case <-ctx.Done():
            break
        }
    }
}
