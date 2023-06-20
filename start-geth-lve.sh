./build/bin/geth \
--nousb \
--http \
--http.port "8545" \
--http.vhosts "*" \
--http.corsdomain "*" \
--http.api "admin,db,eth,net,web3,miner,personal" \
--http.addr "0.0.0.0" \
--gcmode "full" \
--syncmode "full" \
--gwangju \
--ws --ws.port 8546 --ws.api eth,net,web3 --ws.origins '*' --ws.addr 0.0.0.0 \
console

# --rpcaddr 127.0.0.1  
# --rpc --rpcapi db,eth,net,web3,personal --cache=2048  --rpcport 8547 --rpccorsdomain "*" \
# --http.addr "127.0.0.1" \
