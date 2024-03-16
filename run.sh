

if [ $1 = 'server' ];then
    go run ./server
    # echo "server file"
elif [ $1 = 'client' ];then
    go run ./client $2 $3 $4
    # echo "client file"
else
    echo "invalid command"
fi