#!/bin/bash
#https://gist.github.com/bclinkinbeard/1331790
productName="goSpeedComp"
releaseFolder="release/"

mkdir ${releaseFolder} 2>/dev/null

rm ${releaseFolder}*.tgz 2>/dev/null
rm ${releaseFolder}*.zip 2>/dev/null

arch=amd64
os=darwin
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/gospeedcomp/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/
echo -\> $product build and zipped

arch=arm
os=linux
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/gospeedcomp/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/
echo -\> $product build and zipped

arch=amd64
os=linux
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/gospeedcomp/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/
echo -\> $product build and zipped

arch=386
os=linux
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/gospeedcomp/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/
echo -\> $product build and zipped

arch=386
os=windows
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o ${product}.exe ./cmd/gospeedcomp/...
zip -r ${product}.zip ${product}.exe
rm $product.exe
mv $product.zip ${releaseFolder}/
echo -\> $product build and zipped

arch=amd64
os=windows
product=${productName}_${os}_${arch}
env GOOS=$os GOARCH=$arch packr build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o ${product}.exe ./cmd/gospeedcomp/...
zip -r ${product}.zip ${product}.exe
rm $product.exe
mv $product.zip ${releaseFolder}/
echo -\> $product build and zipped