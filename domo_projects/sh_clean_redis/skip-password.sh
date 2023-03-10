#!/usr/bin/expect -f
spawn ./wget-20230306043654.sh -H
expect "Enter your openid : "
send "heyang18\r"
expect "Enter password : "
send "Heheyy_123\r"
expect eof