#!/bin/bash

KEY=`cat ~/.ssh/id_rsa.pub`
# sed -i "s|<PUBLIC_KEY>|${KEY}|g" deploy-ec2.yaml
sed -i '' "s|<PUBLIC_KEY>|${KEY}|g" deploy-ec2.yaml

aws cloudformation create-stack --stack-name jp-search-combinations --template-body file://deploy-ec2.yaml
sleep 30
git checkout deploy-ec2.yaml
PUBLIC_IP=`aws cloudformation describe-stacks --stack-name jp-search-combinations | jq -r ".Stacks[0].Outputs[0].OutputValue"`
URL="http://${PUBLIC_IP}/results?q=test&n=3"
echo ${URL}
