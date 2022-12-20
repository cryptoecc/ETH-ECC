# ETH-ECC client launch and LVE mainnet connection
Project PI : Heung-No Lee

Email : heungno@gist.ac.krk

Writer : Seungmin Kim(김승민), Hyoungsung Kim(김형성)

Email : seungminkim@gist.ac.kr

Github for this example : https://github.com/cryptoecc/ETH-ECC

For more information : [INFONET](https://heungno.net/)

- Eth-ECC is an Ethereum blockchain which INFONET has made public at github. The new characteristic of this blockchain is to enable ECCPoW a new protocol for time-varying proof-of-work generation system. 
- This package is designed to run in Linux-environment. For this you need to download and install Linux-mint (see below). If you are Windows users, you need to install the Linux-mint first. If you are Linux users, you may skip this part. 
- Under the assumption that you are in a Linux environment, you may follow this note to proceed. This note is to illustrate how to locate the Eth-ECC package at the github link, download, install, and run Eth-ECC in a local computer.

---

Contents

1. [Environment](#1-environment)
   1. [Download](#11-download)
   2. [Install](#12-install)
2. [Run Eth-ECC and connect to LVE mainnet](#2-run-eth-ecc-in-your-local-computer)

3. [Test Eth-ECC](#3-test-eth-ecc)
   1. [Basic tests](#31-basic-tests)
   2. [Make a transaction for testing private network](#32-make-a-transaction-for-testing-private-network)
---

## 1. Environment
For Windows, please visit [Windows instruciton](https://github.com/cryptoecc/ETH-ECC/blob/master/docs/eccpow%20windows%20instuction/Windows%20install%20instruction.md) before start.
Eth-ECC package uses the follow two environment

- Amazon Linux 2 Kernel 5.10 or Linux Ubuntu 18.04
- Go (version 1.17 or later) develope language

You can follow two step below to download Eth-ECC and install(build)

### 1.1 Download

Install GO

```
$ sudo apt update
$ sudo apt install snapd
$ sudo snap install go --classic
```

```
$ go verison
go version go1.19.4 linux/amd64
```

Download can be done as follows

- First `clone` Eth-ECC repository to a directory that you want to locate Eth-ECC
- For `clone`, open terminal in the driectory and type command below

```
$ git clone https://github.com/cryptoecc/ETH-ECC
```

Then Eth-ECC will be downloaded.

### 1.2 Installation of Eth-ECC

- Installation of Eth-ECC can be done as follows

After download is completed, move to `ETH-ECC` directory, open terminal and type this command

```
$ make all
```

If build is successful then you can see `geth` in /ETH-ECC/build/bin

![you can see the geth in bin folder](./img/geth-directory.png)

## 2. Run Eth-ECC and connect to LVE mainnet

Now we have to make a directory to store information. I made `geth-test` folder. You can make it anywhere. my directory is `/home/hskim/Documents/geth-test`

Move to `/Eth-ECC/build/bin`, open terminal and follow it

```
$ ./geth --lve --datadir Your_own_storage console
```

In my case,

```
(EXAMPLE) $ ./geth --lve --datadir /home/hskim/Documents/geth-test console
```

Then you can see

```
INFO [12-20|16:23:02.796] Starting Geth on Lve ... 
INFO [12-20|16:23:02.798] Maximum peer count                       ETH=50 LES=0 total=50
INFO [12-20|16:23:02.799] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
INFO [12-20|16:23:02.802] Set global gas cap                       cap=50,000,000
INFO [12-20|16:23:02.805] Allocated trie memory caches             clean=154.00MiB dirty=256.00MiB
INFO [12-20|16:23:02.805] Allocated cache and file handles         database=/home/lvminer3/seungmin/test/geth/chaindata cache=512.00MiB handles=524,288
INFO [12-20|16:23:03.266] Opened ancient database                  database=/home/lvminer3/seungmin/test/geth/chaindata/ancient/chain readonly=false
INFO [12-20|16:23:03.267] Initialising Ethereum protocol           network=12345 dbversion=<nil>
INFO [12-20|16:23:03.267] Writing custom genesis block 
INFO [12-20|16:23:03.293] Persisted trie from memory database      nodes=354 size=50.23KiB time=2.209488ms gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [12-20|16:23:03.295]  
INFO [12-20|16:23:03.296] --------------------------------------------------------------------------------------------------------------------------------------------------------- 
.
.
.
>
```

At last line, you can see console pointer(>).
we can check ChainID is 12345.

You are now connected to the LVE network!

## 3. Test Eth-ECC

1. Basic tests
   - We will make an account, mine block and check result of mining
2. Make a transaction
3. (Appendix) Block generation time log

### 3.1 Basic tests

In this test, we will follow 3 steps

- Generate account
- Check account's balance
- Mining

Now let's test lve network

```
> eth.accounts
[]
```

`eth.account` check accounts of network. There are no accounts.

#### Generate account

Let's generate account

```
> personal.newAccount("Alice")
INFO [08-06|21:33:36.241] Your new key was generated               address=0xb8C941069cC2B71B1a00dB15E6E00A200d387039
WARN [08-06|21:33:36.241] Please backup your key file!             path=/home/hskim/Documents/geth-test/keystore/UTC--2019-08-06T12-33-34.442823142Z--b8c941069cc2b71b1a00db15e6e00a200d387039
WARN [08-06|21:33:36.241] Please remember your password! 
"0xb8c941069cc2b71b1a00db15e6e00a200d387039"
```

We just generated the address of Alice:`0xb8C941069cC2B71B1a00dB15E6E00A200d387039`. We can check using geth

```
> eth.accounts
["0xb8c941069cc2b71b1a00db15e6e00a200d387039"]
```

we will use it as miner's address so block generation reward will be sent to Alice's address

#### Check account's balance

Before mining, let's check Alice's balance

```
> eth.getBalance("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
0
> eth.getBalance(eth.accounts[0])
0
```

There are 2 ways to check balance.

	- First, using address directly
	- Second, using random access of `eth.accounts`. Alice's account is first account of our network. so we can call it as `eth.account[0]`

I will use first one in this example to avoid confusion.

As we expect, there is no ether at all. Let's do mining.

#### Mining

First we have to set miner's address. For this, we will use 3 commands

- miner.setEtherbase(address)
  - It sets miner's address. Mining reward will be sent to this account
- miner.start(number of threads)
  - Start mining. You can set how many threads you will use. I will use 1 thread
  - If your CPU has enough core, you can use higher number. It will work faster.
- miner.stop()
  - Stop mining

```
> miner.setEtherbase("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
true
> miner.start(1)
null
INFO [08-06|21:42:38.198] Updated mining threads                   threads=1
INFO [08-06|21:42:38.198] Transaction pool price threshold updated price=1000000000
null
> INFO [08-06|21:42:38.198] Commit new mining work                   number=1 sealhash=4bb421…3f463a uncles=0 txs=0 gas=0 fees=0 elapsed=325.066µs
INFO [08-06|21:42:40.752] Successfully sealed new block            number=1 sealhash=4bb421…3f463a hash=4b2b78…4808f6 elapsed=2.554s
INFO [08-06|21:42:40.752] 🔨 mined potential block                  number=1 hash=4b2b78…4808f6

.
.
.

INFO [08-06|21:42:56.174] 🔨 mined potential block                  number=9 hash=2faebb…8be693
INFO [08-06|21:42:56.174] Commit new mining work                   number=10 sealhash=384aa6…cb0596 uncles=0 txs=0 gas=0 fees=0 elapsed=179.463µs
> miner.stop()
null
```

We finished mining. Now let's check it worked well.

```
> eth.getBalance("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
45000000000000000000
```

Wow! we got a reward. Exactly `wei`, not `ether`. wei is small unit of ether like satoshi of bitcoin

```
1 ether = 10^18 wei
```

We can convert it to ether by command

```
> web3.fromWei(eth.getBalance("0xb8c941069cc2b71b1a00db15e6e00a200d387039"), "ether")
45
```

Yeah! we got 45 ether. But why 45 ether? To know it, we have to see the source code of geth(go-ethereum)

### 3.2 Make a transaction for testing private network

In this section, we want to generate transaction and send ether

- We will make a new account(Bob) and will send ether from miner(Alice) to new account(Bob)

Generate new account

```
> personal.newAccout("Bob")
INFO [08-06|22:00:23.416] Your new key was generated               address=0xf39Cf42Cd233261cd2b45ADf8fb1E5A1e61A6f90
WARN [08-06|22:00:23.416] Please backup your key file!             path=/home/hskim/Documents/geth-test/keystore/UTC--2019-08-06T13-00-21.621172635Z--f39cf42cd233261cd2b45adf8fb1e5a1e61a6f90
WARN [08-06|22:00:23.416] Please remember your password! 
"0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90"

> eth.getBalance("0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90")
0
```

I got account of Bob:`0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90` Alice will send ether to Bob's account

```
> eth.sendTransaction({from: "0xb8c941069cc2b71b1a00db15e6e00a200d387039", to: "0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90", value: web3.toWei(5, "ether")})
```

Let's send 5 ether to Bob's account

- Alice's account : `0xb8c941069cc2b71b1a00db15e6e00a200d387039`
- Bob's account : `0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90`

Or we can initialize these using variable

```
> from = "0xb8c941069cc2b71b1a00db15e6e00a200d387039"
> to = "0xb8c941069cc2b71b1a00db15e6e00a200d387039"
> eth.sendTransaction({from: from, to: to, value: web3.toWei(5, "ether")})
```

We have to unlock Alice's account. Let's see status of Alice's account.

```
> personal.listWallets[0].status
"Locked"
```

Yes It is locked. So we have to unlock it to send ether from Alice to Bob

```
> web3.personal.unlockAccount("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
Unlock account 0xb8c941069cc2b71b1a00db15e6e00a200d387039
```

Alice's address is `0xb8c941069cc2b71b1a00db15e6e00a200d387039`. However we have to type a `Passphrase` of Alice. `passphrase` is `Alice` cause we generate this address using `Alice`

>Remember it!
>
>```
>> personal.newAccount("Alice")
>```

```
Passphrase: Alice
true
```

Now Alice's account is unlocked. Let's go back to transaction. We can see pending transactions

```
> eth.pendingTransactions
[]
```

Until now, there is not any transaction. We just unlocked Alice's account. Let's make a transaction again.

```
> eth.sendTransaction({from: "0xb8c941069cc2b71b1a00db15e6e00a200d387039", to: "0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90", value: web3.toWei(5, "ether")})
INFO [08-06|22:16:09.274] Setting new local account                address=0xb8C941069cC2B71B1a00dB15E6E00A200d387039
INFO [08-06|22:16:09.275] Submitted transaction                    fullhash=0x926f1bb71d5b48a306e6cde2d45c01f8af2107febf94b166a7e5f8e025dc8adc recipient=0xf39Cf42Cd233261cd2b45ADf8fb1E5A1e61A6f90
"0x926f1bb71d5b48a306e6cde2d45c01f8af2107febf94b166a7e5f8e025dc8adc"
```

There is no error. Let's see a pending transactions

```
> eth.pendingTransactions
[{
    blockHash: null,
    blockNumber: null,
    from: "0xb8c941069cc2b71b1a00db15e6e00a200d387039",
    gas: 21000,
    gasPrice: 1000000000,
    hash: "0x926f1bb71d5b48a306e6cde2d45c01f8af2107febf94b166a7e5f8e025dc8adc",
    input: "0x",
    nonce: 0,
    r: "0x70484271bdc85f7233e715423d8d0be5c669a323385b5ec0ff080a52cf3c654c",
    s: "0x1b55a792995f61128c10a48ce1e0869893c863d38489f574d84ae3a96b031cef",
    to: "0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90",
    transactionIndex: null,
    v: "0x42",
    value: 5000000000000000000
}]
```

There is transaction.

```
> eth.getBalance("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
45000000000000000000
> eth.getBalance("0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90")
0
```

We didn't mine any block. so There is no change of balance yet. Let's mine again!

```
> miner.start(1)
INFO [08-06|22:19:53.061] Updated mining threads                   threads=1
INFO [08-06|22:19:53.061] Transaction pool price threshold updated price=1000000000
null
> INFO [08-06|22:19:53.062] Commit new mining work                   number=10 sealhash=f69cfb…273c0d uncles=0 txs=0 gas=0 fees=0 elapsed=265.557µs
INFO [08-06|22:19:53.062] Commit new mining work                   number=10 sealhash=a018f5…65f494 uncles=0 txs=1 gas=21000 fees=2.1e-05 elapsed=1.022ms
INFO [08-06|22:19:54.718] Successfully sealed new block            number=10 sealhash=a018f5…

.
.
.

INFO [08-06|22:20:05.086] 🔨 mined potential block                  number=16 hash=e7688a…09ed64
INFO [08-06|22:20:05.086] Commit new mining work                   number=17 sealhash=6b297d…b76b19 uncles=0 txs=0 gas=0     fees=0       elapsed=252.945µs
> miner.stop()
null
```

Last time Alice mined 9 blocks and this time mined 7 blocks more. So We can expect Alice has 75 ether (80 ether block reward - 5 ether sent to Bob = 75 ether). Let's check.

First,

```
> eth.pendingTransactions
[]
```

There is no pending transaction. Alice and Bob's transaction is done!

Let's see balance of them

```
> eth.getBalance("0xb8c941069cc2b71b1a00db15e6e00a200d387039")
75000000000000000000
> eth.getBalance("0xf39cf42cd233261cd2b45adf8fb1e5a1e61a6f90")
5000000000000000000
```

As we expected, Alice has 75 ether, Bob has 5 ether. We did it!

---

If there are errors or you want to add more details, please make a issue in my github or official Eth-ECC gitbub

Github :

- Official : https://github.com/cryptoecc/ETH-ECC
