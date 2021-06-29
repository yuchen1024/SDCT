package sdct

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sdct/utils"
)

const (
	// ContractsPath is the location of truffle contract build files.
	ContractsPath = "./solidity/build/contracts"

	// ABI is the special field in contract json files which specify
	// contract's abi.
	ABI = "abi"
	// BIN is the special filed in contract json files which specify
	// contract's initial binary code to be deployed.
	BIN = "bytecode"
)

var contractsName = []map[string]string{
	{
		"name": "SDCTSetup",
	},
	{
		"name": "SDCTSystem",
	},
	{
		"name": "SDCTVerifier",
	},
	{
		"name": "Token",
	},
	{
		"name": "TokenConverter",
	},
}

// GenerateGoCode automatically generates go files for contract defined in truffle contracts.
func GenerateGoCode(newFlag bool) {
	for _, v := range contractsName {
		path := filepath.Join(ContractsPath, v["name"]+".json")
		files := generateGoCode(path, v["name"], newFlag)
		for _, f := range files {
			utils.Delete(f)
		}
	}
}

// GenerateJavaCode automatically generates java code to call/send contract tx.
func GenerateJavaCode(flag bool) {
	for _, v := range contractsName {
		if v["name"] != "SDCT" {
			continue
		}
		path := filepath.Join(ContractsPath, v["name"]+".json")
		files := generateJavaCode(path, v["name"])
		if flag {
			for _, f := range files {
				utils.Delete(f)
			}
		}
	}
}

func generateGoCode(path, name string, newFlag bool) []string {
	raw := utils.Read(path)

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	abiRaw, err := json.Marshal(data[ABI])
	binRaw, err := json.Marshal(data[BIN])
	if err != nil {
		panic(err)
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file
	name = strings.ToLower(name)
	abiName := name + ".abi"
	binName := name + ".bin"
	var outName string
	var dir string
	if newFlag {
		outName = filepath.Join("./", "new", name, name+".go")
	} else {
		dir = filepath.Join("./contracts", name)
		outName = filepath.Join("./contracts", name, name+".go")
	}

	os.MkdirAll(dir, 0755)

	utils.Write(abiName, abiRaw)
	utils.Write(binName, trimBin)

	cmd := "./abigen"
	command := exec.Command(cmd, "--bin", binName, "--abi", abiName, "--pkg", name, "--out", outName, "--type", name)
	fmt.Printf(command.String())
	fmt.Print(outName, "\n")
	err = command.Run()
	if err != nil {
		panic(err)
	}

	return []string{
		filepath.Join("./", abiName),
		filepath.Join("./", binName),
	}
}

func generateJavaCode(path, name string) []string {
	raw := utils.Read(path)

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	abiRaw, err := json.Marshal(data[ABI])
	binRaw, err := json.Marshal(data[BIN])
	if err != nil {
		panic(err)
	}

	// Trim ""
	trimBin := []byte(strings.Trim(string(binRaw), "\""))

	// Write abi and bin file
	abiName := name + ".abi"
	binName := name + ".bin"

	utils.Write(abiName, abiRaw)
	utils.Write(binName, trimBin)

	cmd := "/home/ubuntu/web3j-3.6.0/bin/web3j"
	command := exec.Command(cmd, "solidity", "generate", binName, abiName, "-p", "contracts", "-o", "./contracts/java")
	err = command.Run()
	if err != nil {
		panic(err)
	}

	return []string{
		filepath.Join("./", abiName),
		filepath.Join("./", binName),
	}
}
