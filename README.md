## Go AERUMTechnology

Official golang implementation of the AERUMTechnology protocol featuring our custom DxPoS Consensus Algorithm (Delegated Cross-chain Proof-of-stake).

If you need any assistance please do not hesitate to contact one of the AERUM team via our twitter [@aerum_official](https://twitter.com/aerum_official) or if you would like more information in person feel free to buy us a beer (or coffee).


## How to run

```sh
# Clone the repository
$ git clone https://github.com/AERUMTechnology/go-aerum.git
$ cd go-aerum

# Build the code
$ make all

# Generate a genesis file
$ ./build/bin/puppeth

# Create the data directory
$ mkdir -p ~/.aerum/data

# Run the application. You can use either one of the two methods:

# Method 1: Standard mode
$ ./build/bin/aerum \
  --datadir=./datadir/ \
  --atmos.ethereum.endpoint="https://rinkeby.infura.io" \
  --syncmode 'full' \
  --networkid 418313827693 \
  --ws --wsaddr="0.0.0.0" \
  --wsorigins "*" \
  --wsapi admin,db,eth,debug,miner,net,txpool,clique,web3 \
  --rpc \
  --rpcapi admin,db,eth,debug,miner,net,txpool,clique,web3 \
  --rpcaddr "0.0.0.0" \
  --rpccorsdomain "*" \
  --verbosity="3" \
  --fast \
  --targetgaslimit="25000000" \
  --gasprice="1000000000" \
  --txpool.globalslots="65536" \
  --txpool.accountqueue="1024" \
  --txpool.globalqueue="16384" \
  --cache=1024 \
  console

# Method 2: Delegate mode
$ ./build/bin/aerum \
  --datadir=./datadir/ \
  --atmos.ethereum.endpoint="https://rinkeby.infura.io" \
  --mine \
  --unlock b7668f414ccb7ba45b5ac1a61ce5201e3910ba17 \
  --password <(echo "your password goes here") \
  --rpc \
  --rpcapi admin,db,eth,debug,miner,net,txpool,clique,web3 \
  --rpcport "8547" \
  --rpcaddr "0.0.0.0" \
  --rpcvhosts="*" \
  --rpccorsdomain "*" \
  --networkid="418313827693" \
  --verbosity="3" \
  --fast \
  --targetgaslimit="25000000" \
  --gasprice="1000000000" \
  --txpool.globalslots="65536" \
  --txpool.accountqueue="1024" \
  --txpool.globalqueue="16384" \
  --cache=1024 \
  console
```

If you need to tweak anything to do with the ATMOS consensus check out the file [`./params/atmos_params.go`](params/atmos_params.go)
