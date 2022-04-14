//+build mage

package main

const (
    // Path where to save compiled binaries
    buildPath = "build"

    // Executable name
    executableName = "balance-monitoring"

    // URL to raw contract ABI
    contractABIUrl = "http://api.etherscan.io/api?module=contract&action=getabi&address=0xdAC17F958D2ee523a2206206994597C13D831ec7&format=raw"
)
