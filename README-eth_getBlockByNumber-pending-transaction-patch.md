# README of JSON-RPC API `eth_getBlockByNumber` pending transaction patch

with this patch, JSON-RPC API `eth_getBlockByNumber` support request with params `["pending", false]"` to get pending transaction.

e.g. command:

```shell
curl "http://127.0.0.1:8845" \
             -s \
             -X POST \
             -H "Content-Type: application/json" \
             --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["pending", false], "id": 1}' \
| jq ".result.transactions[0]" \
| xargs -I TX_HASH \
        curl "http://127.0.0.1:8845" \
             -s \
             -X POST \
             -H "Content-Type: application/json" \
             --data '[{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["TX_HASH"],"id":1}, {"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":2}]' \
| jq
```
=>
```javascript
[
  {
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
      "blockHash": null,
      "blockNumber": null,
      "from": "0x533aea009c6ab84d242111300d90dfdb735e60bc",
      "gas": "0x7d00",
      "gasPrice": "0x14124e92080",
      "hash": "0xbe0ae256657a72bad7a8bda7419c696605b265afdfa7e75218b12125b82417fa",
      "input": "0xfdd01289",
      "nonce": "0x1c0e0",
      "to": "0x0000000000dba7f30ba877d1d66e5324858b1278",
      "transactionIndex": null,
      "value": "0x0",
      "type": "0x0",
      "v": "0x93",
      "r": "0x2c829d2c28e870a3ef8a643190f8081622607f0f9c344ce98c7558dba154865d",
      "s": "0x19c36fb625947120da01f375344214df66a7f46d33391be9f54110787d6199bd"
    }
  },
  {
    "jsonrpc": "2.0",
    "id": 2,
    "result": "0x11519e7"
  }
]
```

Note: `blockHash` & `blockNumber` in response is `null` because current transaction is still pending, NOT on chain yet.

## usage

### prepare

create account with command:

```shell
sudo -u bsc /usr/local/bin/geth-bsc --datadir /mnt/data/bsc account new
```
=>
```log
INFO [05-23|17:28:44.622] Maximum peer count                       ETH=50 LES=0 total=50
INFO [05-23|17:28:44.622] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
Your new account is locked with a password. Please give a password. Do not forget this password.
Password:
Repeat password:

Your new key was generated

Public address of the key:   0xE2D97461d8F1715A4c10034dDB44149D1D93b126
Path of the secret key file: /mnt/data/bsc/keystore/UTC--2022-05-23T09-28-59.982859146Z--e2d97461d8f1715a4c10034ddb44149d1d93b126

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
```

Note: please enter password of new account and hit `Enter` twice to complete.

Note: store password as file in path `/etc/bsc/password.txt`

### run `geth` of BSC with pending transaction patch

```shell
/usr/local/bin/geth-bsc --config /etc/bsc/config.toml --datadir /mnt/data/bsc --cache 64000 --metrics --metrics.addr 0.0.0.0 --metrics.port 6060 --maxpeers 100 --diffsync --miner.etherbase 0xE2D97461d8F1715A4c10034dDB44149D1D93b126 --rpc.allow-unprotected-txs --mine --unlock 0xE2D97461d8F1715A4c10034dDB44149D1D93b126 --password /etc/bsc/password.txt --allow-insecure-unlock
```

Note: key parameters: `--miner.etherbase` `--mine` `--unlock` `--password` & `--allow-insecure-unlock`

Note: after BSC node synced(`eth.syncing` return `false`), there will be logs like `lvl=warn msg="Block sealing failed" err="unauthorized validator"`, just ignore them.
