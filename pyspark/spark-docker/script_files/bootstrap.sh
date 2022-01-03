#!/bin/bash
#service ssh start
echo '192.168.29.121 archlinux' > /etc/hosts

#load-spark-env.sh
if [ "$HOSTNAME" = node-master ]; then
    start-master.sh
else
	spark-class org.apache.spark.deploy.worker.Worker -c 1 -m 2G spark://node-master:7077
fi
#bash
while :; do :; done & kill -STOP $! && wait $!
