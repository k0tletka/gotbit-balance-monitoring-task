//+build mage

package main

import (
    "path/filepath"
    "net/http"
    "os/exec"
    "os"
    "log"

    "github.com/magefile/mage/sh"
    "github.com/magefile/mage/mg"
)

const (
    projectName = "github.com/k0tletka/gotbit-balance-monitoring-task"
    contractInterfacePackage = "tether_contract"
)

func GenerateABIInterface() error {
    // Check if generated file already exist
    contractInterfaceLocation := filepath.Join(contractInterfacePackage, contractInterfacePackage + ".go")
    if _, err := os.Stat(contractInterfaceLocation); !os.IsNotExist(err) {
        return nil
    }

    // Install abigen
    log.Println("Downloading abigen...")
    if err := sh.RunV("go", "install", "github.com/ethereum/go-ethereum/cmd/abigen@latest"); err != nil {
        return err
    }

    // Download contract ABI json
    log.Println("Downloading contract ABI...")
    resp, err := http.Get(contractABIUrl)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    os.MkdirAll(contractInterfacePackage, 0775)

    // Run abigen
    abigenProcess := exec.Command(
        "abigen",
        "--abi", "-",
        "--pkg", contractInterfacePackage,
        "--out", contractInterfaceLocation,
        "--alias", "totalSupply=total",
    )

    abigenProcess.Stdin = resp.Body
    abigenProcess.Stdout = os.Stdout

    log.Println("Generating contarct interface...")
    return abigenProcess.Run()
}

func Tidy() error {
    return sh.RunV("go", "mod", "tidy")
}

func Build() error {
    mg.Deps(GenerateABIInterface)
    mg.Deps(Tidy)

    return sh.RunV("go", "build", "-o", filepath.Join(buildPath, executableName), projectName)
}
