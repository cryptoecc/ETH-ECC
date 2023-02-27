./build/bin/geth \
--nousb \
--http \
--http.port "8545" \
--http.vhosts "*" \
--http.corsdomain "*" \
--http.api "admin,db,eth,net,web3,miner,personal" \
--gcmode "full" \
--syncmode "full" \
--lve \
console 
