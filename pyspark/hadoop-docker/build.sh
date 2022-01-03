#!/bin/bash

# generate ssh key
echo "Y" | ssh-keygen -t rsa -P "" -f configs/id_rsa

# Hadoop build
docker build -f ./hadoop/Dockerfile . -t hadoop_cluster:hadoop

#Spark
docker build -f ./spark/Dockerfile . -t hadoop_cluster:spark

# PostgreSQL Hive Metastore Server
#docker build -f ./postgresql-hms/Dockerfile . -t hadoop_cluster:postgresql-hms

# Hive
#docker build -f ./hive/Dockerfile . -t hadoop_cluster:hive

# Nifi
#docker build -f ./nifi/Dockerfile . -t hadoop_cluster:nifi

# Edge
#docker build -f ./edge/Dockerfile . -t hadoop_cluster:edge

# hue
#docker build -f ./hue/Dockerfile . -t hadoop_cluster:hue

# zeppelin
#docker build -f ./zeppelin/Dockerfile . -t hadoop_cluster:zeppelin

