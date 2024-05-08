#! /bin/bash

curl -X GET http://localhost:17000/?cmd=reset
curl -X GET http://localhost:17000/?cmd=white
curl -X GET http://localhost:17000/?cmd=bgrect%200.2%200.2%200.6%200.6 # 200 200 600 600

curl -X GET http://localhost:17000/?cmd=figure%200.3%200.3 # 0.3 0.3

curl -X GET http://localhost:17000/?cmd=green

curl -X GET http://localhost:17000/?cmd=figure%200.3%200.3 # 0.5 0.5

curl -X GET http://localhost:17000/?cmd=update

# curl -X GET http://localhost:17000/?cmd=white%0Abgrect%200.2%200.2%200.7%200.7%0Afigure%200.4%200.4%0Amove%200.6%200.9%0Aupdate