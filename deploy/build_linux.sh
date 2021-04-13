#!/usr/bin/env bash
re="registry.ap-south-1.aliyuncs.com/indiarupeeloan/invest_dairy:([^ ]+)"
imageStr=$(kubectl --kubeconfig $KUBECONFIG_INDIA_ALI get deploy invest-dairy -o jsonpath='{..image}')
echo "current version"
if [[ $imageStr =~ $re ]]; then echo ${BASH_REMATCH[1]}; fi


if [ $# -eq 0 ];
then nextVersion=$(./increment_version.sh -p ${BASH_REMATCH[1]});
else nextVersion=$(./increment_version.sh $1 ${BASH_REMATCH[1]});
fi

echo "next version"
echo $nextVersion;

git add .
git commit -m "invest-dairy:$nextVersion"
git push

rm -rf swagger
cp -rf ../swagger swagger
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o invest-dairy ../main.go

docker build -t registry.ap-south-1.aliyuncs.com/indiarupeeloan/invest_dairy:$nextVersion .
docker push registry.ap-south-1.aliyuncs.com/indiarupeeloan/invest_dairy:$nextVersion
rm -rf invest-dairy
kubectl --kubeconfig $KUBECONFIG_INDIA_ALI set image deployment/invest-dairy invest-dairy=registry.ap-south-1.aliyuncs.com/indiarupeeloan/invest_dairy:$nextVersion
sleep 2
kubectl --kubeconfig $KUBECONFIG_INDIA_ALI get deploy invest-dairy  -o jsonpath='{..image}'
git add ../
git commit -m "$nextVersion $1"
git push
echo $(date "+%Y-%m-%d %H:%M:%S")
