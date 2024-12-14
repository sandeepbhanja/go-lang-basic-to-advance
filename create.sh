#!/bin/bash

fileName=$2
index=$1

folderName="0$index-$fileName"

mkdir $folderName

cd $folderName

go mod init $fileName

touch main.go


