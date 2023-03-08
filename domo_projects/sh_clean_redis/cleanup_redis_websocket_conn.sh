#!/bin/bash
echo Script Name: "$0"
echo Redis Host: "$1"
echo Redis Port: "$2"
echo Password: "$3"

redis_host=$1
redis_port=$2
redis_password=$3
key_pattern='log_*'

redis-cli -h $redis_host -p $redis_port -a $redis_password KEYS "$key_pattern" | xargs redis-cli -h $redis_host -p $redis_port -a $redis_password DEL