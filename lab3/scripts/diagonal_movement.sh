#!/bin/bash

curl -X GET http://localhost:17000/?cmd=white%0Afigure%200.0%200.0%0Aupdate
#curl -X GET http://localhost:17000/?cmd=white
#curl -X GET http://localhost:17000/?cmd=figure%200.0%200.0
#curl -X GET http://localhost:17000/?cmd=update

for ((i=1; i<10; i++))
do
    sleep 0.5  # Затримка на 1 секунду
    x=$(awk -v i="$i" 'BEGIN { print 0.1 * i }')
    y=$(awk -v i="$i" 'BEGIN { print 0.1 * i }')
    curl -X GET http://localhost:17000/?cmd=move%20$x%20$y%0Aupdate
    #curl -X GET http://localhost:17000/?cmd=move%20$x%20$y
    #curl -X GET http://localhost:17000/?cmd=update
done