#!/usr/bin/env bash
# Copy Large File (CLF) 編譯腳本
# 與版本：snapshot_20181015:19##(testing) 同步釋出
#

exePath="/usr/bin/go /bin/go"

goExists="0"
for i in $exePath; do
  if [ -e $i ]; then
    goExists="1"
  fi
done

if [[ $goExists == 0 ]]; then 
  echo "請安裝 go (>= 1.10) 版本後再繼續編譯。"
  echo "感謝！"
  echo -e "\nPlease build it after installed go (version >= 
1.10)!\nThanks too much!"
  exit 1
fi

echo "[INFO] 開始編譯：Starting build."
go build -o clf.out main.go string.go libs.go
echo "[INFO] 編譯完成，檔案為 clf.out。"
echo "[INFO] Build completed, the file is clf.out."
