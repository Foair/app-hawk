ipaddr=`ifconfig | grep inet | grep addr:10 | awk '{print $2}' | tr -d 'addr:'`
./hawk $ipaddr user.json &
