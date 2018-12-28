import time
import web3
import os

from web3 import Web3, HTTPProvider
from web3.middleware import geth_poa_middleware

contract_address     = '0xbc5f91fc3bad1c6f3ecd69c3bb34c849ab9797b4'
wallet_private_key   = 'e0f95f7b4dd052594f7656171726f342dfea53af717d5bb85ab01e2a224dca14'
wallet_address       = '0xf38edc62732c418ee18bebf89cc063b3d1b57e0c'
contract_abi         = '[ { "inputs": [], "payable": false, "stateMutability": "nonpayable", "type": "constructor" }, { "constant": false, "inputs": [ { "name": "aerumBlock", "type": "uint256" }, { "name": "composer", "type": "address" } ], "name": "addComposer", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "aerumBlock", "type": "uint256" }, { "name": "composer", "type": "address" } ], "name": "removeComposer", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "aerumBlock", "type": "uint256" }, { "name": "timestamp", "type": "uint256" } ], "name": "getComposers", "outputs": [ { "name": "", "type": "address[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "aerumBlock", "type": "uint256" } ], "name": "getValidComposers", "outputs": [ { "name": "", "type": "address[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [], "name": "clean", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" } ]'


def get_composers(w3, governance):
    return governance.functions.getComposers(1000 ,0).call()

def clean_composers(w3, governance, nonce):
    unicorn_txn = governance.functions.clean().buildTransaction({
         'chainId': 4,
         'gas': 300000,
         'gasPrice': w3.eth.gasPrice,
         'nonce': nonce,
         'from': w3.toChecksumAddress(wallet_address)
    })
    print(unicorn_txn)

    signed_txn = w3.eth.account.signTransaction(unicorn_txn, private_key=wallet_private_key)
    w3.eth.sendRawTransaction(signed_txn.rawTransaction)

def add_composers(w3, governance, address, nonce):
    unicorn_txn = governance.functions.addComposer(0, w3.toChecksumAddress('0x' + address)).buildTransaction({
         'chainId': 4,
         'gas': 300000,
         'gasPrice': w3.eth.gasPrice,
         'nonce': nonce,
         'from': w3.toChecksumAddress(wallet_address)
    })
    print(unicorn_txn)

    signed_txn = w3.eth.account.signTransaction(unicorn_txn, private_key=wallet_private_key)
    w3.eth.sendRawTransaction(signed_txn.rawTransaction)


w3 = Web3(HTTPProvider('https://rinkeby.infura.io'))
w3.middleware_stack.inject(geth_poa_middleware, layer=0)
governance = w3.eth.contract(address=w3.toChecksumAddress(contract_address), abi=contract_abi)

print('Start')
print(get_composers(w3, governance))

nonce = w3.eth.getTransactionCount(w3.toChecksumAddress(wallet_address))

print('Clean composers')
clean_composers(w3, governance, nonce)
time.sleep(15)

for x in range(2, 5):
    file = open(os.getcwd() + '/nodes/node' + str(x) + '/account.txt', 'r')
    address = file.read().replace("Address: {", "").replace("}", "").replace("\n", "")
    print('Add composer:' + address)
    add_composers(w3, governance, address, nonce + x - 1)
    time.sleep(15)

print(get_composers(w3, governance))
print('Done')