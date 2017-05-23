#!/bin/bash
#TOUCH命令 一次建立多个文件:wq
#./touch.sh dup

for i in `seq 1 3`;do touch $1$i.go;done
