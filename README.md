## Archie Chain
* Network Name - ArchieChain
* Coin Name - ARC
* Coin Symbol - ARC
* Supply - 10 billion
* Blocktime - 2 Seconds
* Consensus - PoS
* P2P Port - 1478
* JSON-RPC Port - 8545 
* ChainID Main - 1243
* ChainID Test - 1244
* EVM Compatible

## Official Links
* Website - https://archiechain.io
* Mainnet Explorer - https://archiescan.io
* Testnet Explorer - https://testnet.archiescan.io
* Mainnet RPC 1 - https://rpc-main-1.archiechain.io
* Mainnet RPC 2 - https://rpc-main-2.archiechain.io
* Testnet RPC 1 - https://rpc-test-1.archiechain.io
* Testnet RPC 2 - https://rpc-test-2.archiechain.io

### Build from Source (Ubuntu 20.04)
Requirements - `Go >=1.18.x`

#### Setup Go Path
```
sudo nano ~/.profile
```
Paste this into the bottom of the file
```
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
```
source ~/.profile
```

#### Install Go
```
wget https://go.dev/dl/go1.18.7.linux-amd64.tar.gz
sudo tar -xvf go1.18.7.linux-amd64.tar.gz
sudo mv go /usr/local && rm go1.18.7.linux-amd64.tar.gz
```
Check that it's installed
```
go version
```
You should see something like this:
```
go version go1.18.7 linux/amd64
```

#### Build archiechain
```
git clone https://github.com/archieneko/archiechain.git
cd archiechain/
go build -o archie main.go
mkdir ~/.archiechain
```

#### Running a Full Validating Node
After you have [downloaded](https://github.com/archieneko/archiechain/releases/latest) the binaries or [built from source](https://github.com/archieneko/archiechain#build-from-source), follow [this guide](ValidatorGuide.md) to setup your validator.

#### Running a non-Validating node
```
mkdir ~/.archiechain
```
```
./archie server --data-dir ~/.archiechain --chain mainnet-genesis.json --libp2p 0.0.0.0:1478 --nat <public_or_private_ip>
```

---
```
Copyright 2022 Archie Chain

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
