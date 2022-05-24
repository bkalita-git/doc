echo "Trying to kill previous processes if any ... "
sudo killall avalanchego
sleep 2
echo "Trying to delete all previous db ... "
sudo rm -r /home/common/Programs/Arch_linux/avalanchego/db/*
sleep 10

#node.json 					 -> /home/common/Programs/Arch_linux/avalanchego/db/node1/.avalanchego/configs/node.json
#chain config.json directory -> /home/common/Programs/Arch_linux/avalanchego/db/node1/.avalanchego/chain/C/config.json
#							 -> /home/common/Programs/Arch_linux/avalanchego/db/node1/.avalanchego/chain/metavm/config.json

AVA_EXE=/home/common/Programs/Arch_linux/avalanchego/build/avalanchego
ROOTDIR=/home/common/Programs/Arch_linux/avalanchego

echo "***************************SETUP LOCAL MAIN NET*************************"
N=3 #first node is the main. so no need to config. ie from node1 - node2, not node0
for ((i=1;i<=$N;i+=1))
do
    NODE_CONFIG_DIR=${ROOTDIR}/db/node$i/.avalanchego
    mkdir -p ${NODE_CONFIG_DIR}/chains/C && mkdir -p ${NODE_CONFIG_DIR}/configs
    echo '{' > ${NODE_CONFIG_DIR}/configs/node.json
    echo '    "log-level": "info",' >> ${NODE_CONFIG_DIR}/configs/node.json
    echo '    "api-ipcs-enabled": true,' >> ${NODE_CONFIG_DIR}/configs/node.json
    echo "    \"ipcs-path\": \"${NODE_CONFIG_DIR}\"" >> ${NODE_CONFIG_DIR}/configs/node.json
#    echo '	  "whitelisted-subnets":"p4jUwqZsA2LuSftroCd3zb4ytH8W99oXKuKVZdsty7eQ3rXD6" ' >>  ${NODE_CONFIG_DIR}/configs/node.json
    echo '}' >> ${NODE_CONFIG_DIR}/configs/node.json

    cat > ${NODE_CONFIG_DIR}/chains/C/config.json <<EOF
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "coreth-performance-api-enabled": false,
  "net-api-enabled": true,
  "rpc-gas-cap": 2500000000,
  "rpc-tx-fee-cap": 100,
  "eth-api-enabled": true,
  "personal-api-enabled": false,
  "tx-pool-api-enabled": true,
  "debug-api-enabled": true,
  "web3-api-enabled": true,
  "local-txs-enabled": true,
  "pruning-enabled": false,
  "api-max-duration": 0,
  "api-max-blocks-per-request": 0,
  "allow-unfinalized-queries": true,
  "log-level": "info",
  "eth-apis": [
    "public-eth",
    "public-eth-filter",
    "net",
    "web3",
    "internal-public-eth",
    "internal-public-blockchain",
    "internal-public-transaction-pool"
  ]
}
EOF
done


# all theese below config can be done using json config for each node ${ROOTDIR}/db/node$i/.avalanchego.avalanchego/configs/node.json and just supply this config # using ---config-file= ${ROOTDIR}/db/node$i/.avalanchego.avalanchego/configs/node.json
${AVA_EXE} \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9650 \
--staking-port=9651 \
--api-admin-enabled=true \
--db-dir=${ROOTDIR}/db/node0 \
--staking-tls-cert-file=${ROOTDIR}/staking/local/staker1.crt \
--staking-tls-key-file=${ROOTDIR}/staking/local/staker1.key \
>> ${ROOTDIR}/node0.log 2>&1 &

${AVA_EXE} \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9652 \
--staking-port=9653 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--config-file=${ROOTDIR}/db/node1/.avalanchego/configs/node.json \
--chain-config-dir=${ROOTDIR}/db/node1/.avalanchego/chains \
--db-dir=${ROOTDIR}/db/node1 \
--staking-tls-cert-file=${ROOTDIR}/staking/local/staker2.crt \
--staking-tls-key-file=${ROOTDIR}/staking/local/staker2.key \
>> ${ROOTDIR}/node1.log 2>&1 &

${AVA_EXE} \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9654 \
--staking-port=9655 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--config-file=${ROOTDIR}/db/node2/.avalanchego/configs/node.json \
--chain-config-dir=${ROOTDIR}/db/node2/.avalanchego/chains \
--db-dir=${ROOTDIR}/db/node2 \
--staking-tls-cert-file=${ROOTDIR}/staking/local/staker3.crt \
--staking-tls-key-file=${ROOTDIR}/staking/local/staker3.key \
>> ${ROOTDIR}/node2.log 2>&1 &

${AVA_EXE} \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9656 \
--staking-port=9657 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--config-file=${ROOTDIR}/db/node3/.avalanchego/configs/node.json \
--chain-config-dir=${ROOTDIR}/db/node3/.avalanchego/chains \
--db-dir=${ROOTDIR}/db/node3 \
--staking-tls-cert-file=${ROOTDIR}/staking/local/staker4.crt \
--staking-tls-key-file=${ROOTDIR}/staking/local/staker4.key \
>> ${ROOTDIR}/node3.log 2>&1 &



USER_NAME=bipul
PASSWORD=Av@1alAv
FUNDED_PRIVATE_KEY=PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN
FUNDED_P_ADDR=P-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u
IP=127.0.0.1
PORT=9650


echo "wait 20 seconds for bootstraps..."
sleep 20

echo "create user: "
curl -X POST --data '{
 "jsonrpc":"2.0",
 "id"     :1,
 "method" :"keystore.createUser",
 "params" :{
     "username":"'$USER_NAME'",
     "password":"'$PASSWORD'"
 }
}' -H 'content-type:application/json;' $IP:$PORT/ext/keystore


echo "import key: to P chain "
curl -X POST --data '{  
 "jsonrpc":"2.0",    
 "id"     :1,    
 "method" :"platform.importKey", 
 "params" :{ 
     "username":"'$USER_NAME'",
     "password":"'$PASSWORD'",
     "privateKey":"'$FUNDED_PRIVATE_KEY'"    
 }   
}' -H 'content-type:application/json;' $IP:$PORT/ext/bc/P


echo "Create Subnet:"
SUBNET_ID=`curl --location --request POST 'http://'$IP:$PORT'/ext/bc/P' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "platform.createSubnet",
    "params": {
        "controlKeys":[
            "'$FUNDED_P_ADDR'"
        ],
        "threshold":1,
        "from": ["'$FUNDED_P_ADDR'"],
        "changeAddr":"'$FUNDED_P_ADDR'",
     	"username":"'$USER_NAME'",
     	"password":"'$PASSWORD'"
    },
    "id": 1
}' | awk -F: '{ split($4,a,","); print a[1]}'`



#29uVeLPJB1eQJkzRemU8g8wZDw5uJRqpab5U2mX9euieVwiEbL subnet
#Cm6TaZzBfcy1cSYT11sELCkcoXREZNk3d71cBM4FhpZW5QUXc	blockchain metavm
#"NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg": 2000000000000000,
#"NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu": 2000000000000000,
#"NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ": 2000000000000000,
#"NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN": 2000000000000000,
#"NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5": 2000000000000000


sleep 10
echo "Add validators/Node to our subnet custom"
#last node ie. add node4 since we own node4
##starttime and endtime from python time.time()+50   end time - starttme > 1
curl --location --request POST 'http://'$IP:$PORT'/ext/bc/P' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "platform.addSubnetValidator",
    "params": {
        "nodeID":"NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
        "subnetID":'$SUBNET_ID',
        "startTime":1654305000,
        "endTime":1655304818,
        "weight":20,
	 	"username":"'$USER_NAME'",
	 	"password":"'$PASSWORD'"
    },
    "id": 1
}'


echo "****************************SETUP SINGLE NODE*************************"

NODE_CONFIG_DIR=${ROOTDIR}/db/node4/.avalanchego
mkdir -p ${NODE_CONFIG_DIR}/chains/C && mkdir -p ${NODE_CONFIG_DIR}/configs
echo '{' > ${NODE_CONFIG_DIR}/configs/node.json
echo '    "log-level": "info",' >> ${NODE_CONFIG_DIR}/configs/node.json
echo '    "api-ipcs-enabled": true,' >> ${NODE_CONFIG_DIR}/configs/node.json
echo "    \"ipcs-path\": \"${NODE_CONFIG_DIR}\"," >> ${NODE_CONFIG_DIR}/configs/node.json
echo "    \"whitelisted-subnets\": ${SUBNET_ID} " >> ${NODE_CONFIG_DIR}/configs/node.json
echo '}' >> ${NODE_CONFIG_DIR}/configs/node.json

cat > ${NODE_CONFIG_DIR}/chains/C/config.json <<EOF
{
  "snowman-api-enabled": false,
  "coreth-admin-api-enabled": false,
  "coreth-performance-api-enabled": false,
  "net-api-enabled": true,
  "rpc-gas-cap": 2500000000,
  "rpc-tx-fee-cap": 100,
  "eth-api-enabled": true,
  "personal-api-enabled": false,
  "tx-pool-api-enabled": true,
  "debug-api-enabled": true,
  "web3-api-enabled": true,
  "local-txs-enabled": true,
  "pruning-enabled": false,
  "api-max-duration": 0,
  "api-max-blocks-per-request": 0,
  "allow-unfinalized-queries": true,
  "log-level": "info",
  "eth-apis": [
    "public-eth",
    "public-eth-filter",
    "net",
    "web3",
    "internal-public-eth",
    "internal-public-blockchain",
    "internal-public-transaction-pool"
  ]
}
EOF
${AVA_EXE} \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9658 \
--staking-port=9659 \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--config-file=${ROOTDIR}/db/node4/.avalanchego/configs/node.json \
--chain-config-dir=${ROOTDIR}/db/node4/.avalanchego/chains \
--db-dir=${ROOTDIR}/db/node4 \
--staking-tls-cert-file=${ROOTDIR}/staking/local/staker5.crt \
--staking-tls-key-file=${ROOTDIR}/staking/local/staker5.key \
>> ${ROOTDIR}/node4.log 2>&1 &

echo "wait 20 seconds for bootstraps..."
sleep 20

USER_NAME=bipul
PASSWORD=Av@1alAv
FUNDED_PRIVATE_KEY=PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN
FUNDED_P_ADDR=P-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u
VMID=qBP9WMzdbamSQ9bfxCPM59L5xbWjD81mFB93Fu6KiAbkaNf2y
IP=127.0.0.1
PORT=9658


echo "create user: "
curl -X POST --data '{
 "jsonrpc":"2.0",
 "id"     :1,
 "method" :"keystore.createUser",
 "params" :{
     "username":"'$USER_NAME'",
     "password":"'$PASSWORD'"
 }
}' -H 'content-type:application/json;' $IP:$PORT/ext/keystore


echo "import key: to P chain "
curl -X POST --data '{  
 "jsonrpc":"2.0",    
 "id"     :1,    
 "method" :"platform.importKey", 
 "params" :{ 
     "username":"'$USER_NAME'",
     "password":"'$PASSWORD'",
     "privateKey":"'$FUNDED_PRIVATE_KEY'"    
 }   
}' -H 'content-type:application/json;' $IP:$PORT/ext/bc/P

sleep 5
echo "Create Blockchain:"
BLOCKCHAIN_ID=`curl --location --request POST 'http://'$IP:$PORT'/ext/bc/P' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "platform.createBlockchain",
    "params" : {
        "subnetID": '$SUBNET_ID',
        "vmID":"'$VMID'",
        "name":"My new METAVM",
        "genesisData": "YtGH4shV5Q8sdkwPEv3zbgz3VDG6hrbXqXiWg8F3bAvuZSMxA",
     	"username":"'$USER_NAME'",
     	"password":"'$PASSWORD'"
    },
    "id": 1
}' | awk -F: '{ split($4,a,","); print a[1]}'`

echo "VMID    = $VMID"
echo "SUBNET_ID		= $SUBNET_ID"
echo "BLOCKCHAIN_ID = $BLOCKCHAIN_ID"

#getBlock
#proposeBlock
#encode
#decode
#we need genesisdata of our custom blockchain

