# How to setup a Validator for Archie Chain (Testnet)
In order to run a validator, a user needs 100k ARC and a VPS server (or computer that can be online 24hrs a day). If you need a VPS provider to run your validator on, Amazon AWS, Digital Ocean, and Vultr are great options. For this tutorial I will use Vultr. 

Note: You can use [this link](https://www.vultr.com/?ref=8988039-8H) for $100 credit on Vultr.

**VPS Spec Requirements**
* OS - Ubuntu 20.04
* Storage - 55 GB min 
* RAM - 4 GB min

## VPS setup walkthrough
1. [Create an account](https://www.vultr.com/?ref=8988039-8H) on Vultr or login to your chosen VPS provider
2. Click the blue plus sign `+`
3. Select the following:
- Choose Server - Cloud Compute
- CPU & Storage Technology - intel Regular Performance
- Server Location - Choose any location you would like
- Server Image - Ubuntu 20.04
- Server Size - 
- Add Auto Backups - Disable (optional)
- Server Hostname & Label - Name it whatever you like

When you are done, click `Deploy Now` to create the server. 

After it's created, click the server and make note of the server's IP address and password.

## Setting up
After you have created a VPS server and collected more than 100k ARC, then you are ready to proceed. Login to your VPS server with SSH using the Terminal app (if using MacOS) or Command Prompt (if on Windows) and paste in this command. Replace `SERVER_IP` with your servers IP address.

```
ssh root@SERVER_IP
```

Hit enter, then paste in the server's password and hit enter. 

Once you were able to login, you can proceed. 

## Part 1 - Build from Source (Ubuntu 20.04)
Requirements - `Go >=1.18.x`

### Setup Go Path
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

### Install Go
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

### Build archiechain
```
git clone https://github.com/archieneko/archiechain.git
cd archiechain/
go build -o archie main.go
``` 

## Part 2 - Generate Validator Keys
After you finish `Part 1`, Create a directory to store your keys and testnet blockchain data on the VPS server. 
```
mkdir ~/.archietest
cd ~/archiechain
```

Now run this command to generate your validator keys
```
./archie secrets init  --data-dir ~/.archietest
```

You should see something like this:
```
[SECRETS INIT]
Public key (address) = 0xXXXXXXXXXX...
BLS Public key       = 0xXXXXXXXXXX...
Node ID              = 16UXXXXXXXXX...
```

Save that somewhere and continue to the next step. 

## Part 3 - Stake Your Coins
### Setup staking contracts
Download staking contracts:
```
git clone https://github.com/archieneko/staking-contracts.git
cd staking-contracts
npm i
cp .env.example .env
```

Now open up the .env file and fill out your validators private key. Fill out the variables like so:
```
PRIVATE_KEYS=YOUR_VALIDATOR_PRIVATE_KEY_HERE
```

The `YOUR_VALIDATOR_PRIVATE_KEY_HERE` variable can be found in this file: 
```
~/.archietest/consensus/validator.key
```

You can view it like so
```
cat ~/.archietest/consensus/validator.key
```

Save the key in the `.env` file and continue on

### Fund your wallet
Send at least 100000 + 0.01 ARC to the wallet address that you generated in Part 2: 
```
Public key (address) = 0xXXXXX...
```

### Stake to become a validator
Once you have completed the steps above, run the following:
```
npm run stake-test
```

It run run a few seconds until the transaction confirms.

When you are done staking and want to un-stake, run this command:
```
npm run unstake-test
```

## Part 4 - Start Your Validator
Run the following to start your validator. Replace `YOUR_SERVER_IP` with your server's public IP address
```
./archie server --data-dir ~/.archietest --chain testnet-genesis.json --libp2p 0.0.0.0:1478 --nat YOUR_SERVER_IP --seal
```

Once you run that, your validator will start syncing with the testnet blockchain and begin validating. 


