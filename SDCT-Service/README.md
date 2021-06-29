# 可监管机密交易服务

该文档说明如何部署/测试可监管审计的机密交易服务以及代码结构

## 安装环境

测试环境：

* 操作系统：macOS 10.13.6
* Go：go1.15.1 darwin/amd64
* Ganache：v2.4.0

Go语言安装:
请在[golang中国](https://studygolang.com/dl)或者[官方链接](https://golang.org/dl/)下载相应的Go语言安装包，按照教程进行安装。安装结束后运行：

```bash
# 查看安装Go语言版本/检测是否正确安装
go version
# 因本项目使用go mod进行管理，请使用较高的Go语言版本(可直接下载最新版本)
```

如能得到`go version go1.15.1 darwin/amd64`类似输出，则表明Go语言已正确安装，同时请注意设置好了`GOPATH, GOROOT`环境变量.

Ganache以太坊测试链客户端安装：
请在[ganache](https://github.com/trufflesuite/ganache/releases)上下载对应的安装包安装。

## 部署/测试

### 启动Ganache

1. 启动Ganache客户端，单击QUICKSTART快速启动:

2. 进入程序后，单击右上角设置按钮进行SERVER设置, **注意端口设置必须为8545**:

3. ACCOUNT/KEYS设置，**保证账户资金充足**，mnemonic中字符串为`three stock swap matter mutual okay virus guess river behave recall decrease`

4. 设置ACCOUNT DEFAULT BALANCE为128, 设置TOTAL ACCOUNTS TO GENERATE为3

5.	单击CHAIN, **Gas Limit设置为8000000, 与以太坊主网保持一致**

6.	设置GAS PRICE为0

7.	单击SAVE&RESTART



### 部署/测试

目前已使用Go语言实现了可监管机密交易服务的部署与测试，可直接通过终端命令运行.

### 测试流程说明

```
cd $sdct/cmd
go build -o SDCT_test
./SDCT_test
```
注意这是一个交互式的命令, 单击回车自行进入下一步流程
运行时需要导入hashmap, 需要约10s
请注意, 第一次运行时会自动生成hashMap, 大约需要1分钟

1. 使用SDCT客户端部署SDCT合约(出现Send init SDCTVerifier tx succeeds表明已经完成部署，回车进行下一步)

2. 使用SDCT客户端生成Alice、Bob和Carol三个账户 (回车进行下一步) 

3.	Alice、Bob、Carol使用SDCT-ETH对接模块将ETH Coin兑换为SDCT Token (出现Deposit account succeeds表明已经成功, 可以演示ganache界面，回车进行下一步)

4.	Carol使用SDCT客户端转账128 SDCT Token给Bob (出现CTx transfer succeeds表明转账已成功, 回车进行下一步)

5.	Bob使用SDCT客户端转账128 SDCT Token给Alice (出现CTx transfer succeeds表明转账已成功, 回车进行下一步）

6.	Alice和Bob使用SDCT-ETH对接模块将SDCT Token兑换回ETH Coin (出现Burn tx succeeds表明交易已成功, 可演示ganache界面, 回车进行下一步)
(Carol余额为0，无需做转换)
7.	SDCT监管模块对隐私交易实施穿透式监管

交易均成功执行，演示完毕

## Go语言部分的代码结构

/client: client.go: 以太坊客户端

/cmd: main.go sdct交互命令行工具

/contracts: 编译生成的Go语言版本合约binding

/curve: 椭圆曲线

* bn128.go: bn128椭圆曲线
* bn128_test.go: 测试文件
* curve.go: 多个椭圆曲线支持

/deployer: deployer.go: SDCT智能合约部署

/proof: proof/PKE模块

* aggrangeproof: bullet proof
* aggrangeproof_test.go: 测试文件
* calculate_dlog.go: 生成hashMap，解密msg
* calculate_dlog_test.go: 测试文件
* ct_valid: 证明拥有twisted ELGamal明文和随机数的知识
* ct_valid_test.go: 测试文件
* elgamal_ct.go: twisted ElGamal PKE 
* elgamal_test.go: 测试文件
* inner_product.go: bullet proof内联结构
* inner_product_test.go: 测试文件
* nizk_plaintext_quality.go: 证明3个twisted ElGamal加密明文一致
* nizk_plaintext_quality_test.go: 测试文件
* sigma_protocol.go: 证明离散对数关系一致
* sigma_protocol_test.go: 测试文件

/solidity: SDCT合约
* contracts: 合约代码
    * SDCTSetup.sol: SDCT公共参数设置
    * SDCTSystem.sol: SDCT验证执行模块入口
    * SDCTVerifier.sol: SDCT交易验证
    * Migrations.sol: 测试文件
    * Token.sol: ERC20 token
    * TokenConverter.sol: ETH/ERC20 token与SDCT token兑换比例计算

/sdctsys:
* account.go: SDCT账户管理
* sdctsys.go: SDCT-ETH对接模块
* sdctsys_test.go: 测试文件
* sdctsyssolidity_test.go: 测试文件

/utils: 通用模块
* common.go: 挑战/哈希计算
* common_test.go: 测试文件
* ecpoint.go: 椭圆曲线点操作
* field_vector.go: field vector操作
* file.go: 文件操作
* generator_vector.go: 生成元向量操作

/generator.go: 生成Go语言版本合约binding

/generator_test.go: 测试文件

/go.mod: Go语言模块管理文件

/go.sum: Go语言模块管理文件(由go.mod自动生成)

/Readme.md: SDCT markdown格式文档

/Readme.pdf: SDCT pdf格式文档