package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethereum-optimism/optimism/op-program/chainconfig"
	"github.com/ethereum-optimism/optimism/op-program/verify"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	var l1RpcUrl string
	var l1RpcKind string
	var l1BeaconUrl string
	var l2RpcUrl string
	var dataDir string
	flag.StringVar(&l1RpcUrl, "l1", "", "L1 RPC URL to use")
	flag.StringVar(&l1BeaconUrl, "l1.beacon", "", "L1 Beacon URL to use")
	flag.StringVar(&l1RpcKind, "l1-rpckind", "", "L1 RPC kind")
	flag.StringVar(&l2RpcUrl, "l2", "", "L2 RPC URL to use")
	flag.StringVar(&dataDir, "datadir", "",
		"Directory to use for storing pre-images. If not set a temporary directory will be used.")
	flag.Parse()

	if l1RpcUrl == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Must specify --l1 RPC URL")
		os.Exit(2)
	}
	if l1BeaconUrl == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Must specify --l1.beacon URL")
		os.Exit(2)
	}
	if l2RpcUrl == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Must specify --l2 RPC URL")
		os.Exit(2)
	}

	goerliOutputAddress := common.HexToAddress("0xE6Dfba0953616Bacab0c9A8ecb3a9BBa77FC15c0")
	err := verify.Run(l1RpcUrl, l1RpcKind, l1BeaconUrl, l2RpcUrl, goerliOutputAddress, dataDir, "goerli", chainconfig.OPGoerliChainConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed: %v\n", err.Error())
		os.Exit(1)
	}
}
