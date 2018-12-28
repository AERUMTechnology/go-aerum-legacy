#!/bin/bash

# some nice colours...
BLACK='\033[0;30m'
DARKGREY='\033[1;30m'
RED='\033[0;31m'
LIGHTRED='\033[1;31m'
GREEN='\033[0;32m'
LITEGREEN='\033[1;32m'
BROWNORANGE='\033[0;33m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
LIGHTBLUE='\033[1;34m'
PURPLE='\033[0;35m'
LITEPURPLE='\033[1;35m'
CYAN='\033[0;36m'
LITECYAN='\033[1;36m'
LITEGREY='\033[0;37m'
WHITE='\033[1;37m'
NC='\033[0m' # No Color

# script variables
STEPS=7
NODES=7
SIGNINGNODES=4
GENESISFILENAME="genesis.json"
CHAINID=2018

# shortcut or geth
export GETH=$PWD/../../build/bin/aerum

printf "${LITEGREY}#############################################################################################################\n"
printf "####################################### Aerum Consensus Engine DevKit #######################################\n"
printf "############################## The complete consensus engine development kit ################################\n"
printf "############# Create Accounts, Genesis, Init Network, Assign nodekeys, static-nodes.json & more #############\n"
printf "#############################################################################################################${NC}\n\n\n"


printf "Step 1 of ${STEPS}...\n"
printf "${RED}Removing artifact files & Generating new ones... ${NC}\n"
    # remove data directories
    rm -rf $PWD/nodes/
    # remove genesis and its .bak created by previous sed commands
    rm $PWD/$GENESISFILENAME  $PWD/logs/header-hashes.log
    # copy the template to a new genesis file
    cp $PWD/genesis-template $PWD/$GENESISFILENAME
    # recreate all data dirs
    for ((i = 1; i <= NODES; i++))
	do
	    mkdir -p $PWD/nodes/node${i}/data
	done
printf "\n"



printf "Step 2 of ${STEPS}...\n"
printf "${GREEN}Generating Accounts for ${NODES} nodes... ${NC}\n"
    # generate new accounts with default passwords
    for ((i = 1; i <= NODES; i++))
	do
		$GETH --datadir $PWD/nodes/node${i}/data account new --password <(echo "node${i}") > $PWD/nodes/node${i}/account.txt
	done
printf "\n"



printf "Step 3 of ${STEPS}...\n"
printf "${PURPLE}Generate Genesis file and populate with ${SIGNINGNODES} signers... ${NC}\n"
    # assign a normal node and 2 signers addresses to variables to populate the genesis template
    for ((i = 1; i <= SIGNINGNODES; i++))
	do
        NODE=$(cat $PWD/nodes/node${i}/account.txt|cut -f 2 -d " "|sed 's/{//g'|sed 's/}//g')
        sed -i.bak "s/__NODE${i}__/$NODE/g"  $PWD/$GENESISFILENAME
	done
printf "\n"



printf "Step 4 of ${STEPS}...\n"
printf "${CYAN}Initialize the network with the genesis... ${NC}"
    # initialize the network
    for ((i = 1; i <= NODES; i++))
	do
        $GETH --datadir $PWD/nodes/node${i}/data init $PWD/$GENESISFILENAME
    done
printf "\n"



printf "Step 5 of ${STEPS}...\n"
printf "${YELLOW}List of accounts we generated this time around${NC}\n"
    printf "\n"
    for ((i = 1; i <= NODES; i++))
	do
        echo '"50","'$(cat $PWD/nodes/node${i}/account.txt|cut -f 2 -d " "|sed 's/{//g'|sed 's/}//g')'"'
    done
printf "\n"



printf "Step 6 of ${STEPS}...\n"
printf "${PURPLE}Refresh governance contract${NC}\n"
    python3 refresh_governance.py
printf "\n"



printf "Step 7 of  ${STEPS}...\n"
printf "${LITEGREEN}Creating static-nodes.json file in each of our networks nodes (speeds up p2p connectivity) ${NC}\n"

touch $PWD/nodes/{node1,node2,node3,node4,node5,node6,node7}/data/aerum/static-nodes.json
#
# Previously it was necessary to launch geth in order for it to invoke the functions necessary to generate an ECDSA nodekey
# After realising that the function func (c *Config) NodeKey() *ecdsa.PrivateKey checks first to see if a key exists in the dir
# I decided to just generate real public and private keys for reuse - leave these comments in case it becomes necessary to regenerate them on the fly
#

# For each node:
# generate enode from node key https://ethereum.stackexchange.com/questions/28970/how-to-produce-enode-from-node-key
# NODE1key=$(cat $PWD/nodes/node1/data/geth/nodekey)
# NODE1enode=$($PWD/geth/geth-dev/build/bin/bootnode -nodekeyhex $NODE1key -writeaddress)

# Add a private key to each aerum node
touch $PWD/nodes/node1/data/aerum/nodekey && echo "43ad244d69706ad9ed08ab969ba19a5adcf8b4c53fbfd66e0d8a5cf8432cd118" > $PWD/nodes/node1/data/aerum/nodekey
touch $PWD/nodes/node2/data/aerum/nodekey && echo "bd838d3f2699205888a7330238326f94323fb27c0f3a5607ffead80a0ce1a9ea" > $PWD/nodes/node2/data/aerum/nodekey
touch $PWD/nodes/node3/data/aerum/nodekey && echo "1086c0cb6cc3f76c1c60851359d3790a3a2af60e76cd2fb8c6d3b760700ad0af" > $PWD/nodes/node3/data/aerum/nodekey
touch $PWD/nodes/node4/data/aerum/nodekey && echo "a7b05d288e56476b3bc0587e4b3994820bdd626c16bbc807cd650ce25e76f658" > $PWD/nodes/node4/data/aerum/nodekey
touch $PWD/nodes/node5/data/aerum/nodekey && echo "225b75b82e24356d9c368d3ca29931997a28ad149e9e158cdba7fe4cfa387d08" > $PWD/nodes/node5/data/aerum/nodekey
touch $PWD/nodes/node6/data/aerum/nodekey && echo "6df2a56dc07b8c61bd72d1d8177c0c04b7d7d5dfc772bb045fd9d64ae51ea2a2" > $PWD/nodes/node6/data/aerum/nodekey
touch $PWD/nodes/node7/data/aerum/nodekey && echo "6033f4ddf345f8ffddbb9390716be33ecfb55d96c0137b965370b5fbe365bab9" > $PWD/nodes/node7/data/aerum/nodekey

# add valid public key / enode to each node
cat <<EOT >> $PWD/nodes/node1/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node2/data/aerum/static-nodes.json
[
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node3/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node4/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node5/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
   "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node6/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://a44d35662f91938752b9e50bc778941d51099b406c4127e16363f016c3adc301fdd36de5c8560b28aee4770b74290f9512e091e58d2906b540037c2a8a75c701@127.0.0.1:7777"
]
EOT
cat <<EOT >> $PWD/nodes/node7/data/aerum/static-nodes.json
[
  "enode://bea6f37f0f8961d665f805c4bb7c1f2711fd3499e0eb9a6ccfc08fc79443b0cd946e6961113df9090a71c6ba6b1c78b9a555f6e7bae8cfe680a028dc0d36c0aa@127.0.0.1:2222",
  "enode://b5ab72b7e6763147ce185ac218d50b40298062c265c49333558b031bf0516707751788e6c9aa3cfa8d3772fdf35792eeaf75aa1e044af735b9a6f41193fd9f28@127.0.0.1:3333",
  "enode://90f0be041a3878da763ca5187663a1508ee7f2e9c3b095c6f8c0453e618be43fb4f101f62b231e1e4e1494f413a7a8e590aff34a6ed006f29bc01785b359bbf1@127.0.0.1:4444",
  "enode://c12024acd249bea89dc2d0a3c11a7932e5d76ac7c3544cbf0cf761a9f3b555a886cd3009b66974e64b4ce8cc8374b13c71621ad475cdf707ce12cea23bc0ab1c@127.0.0.1:5555",
  "enode://472c502d52d5f3e22f84d13149e821239229436ac3bb098fdffc5c9372251048581989331c963becaa1549f69ece77fcc874ebb358b03e4e6b6fb78bd8bba25a@127.0.0.1:6666"
]
EOT

echo "Static Nodes in use..."
cat $PWD/nodes/node1/data/aerum/static-nodes.json

rm $PWD/$GENESISFILENAME.bak
printf "\n\n"
# end of step 7