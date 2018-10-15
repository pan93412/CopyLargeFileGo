#!/usr/bin/env bash
# Copy Large File (CLF) 編譯腳本
# 與版本：v0.6_beta 同步釋出
#

exePath="/usr/bin/go /bin/go"
GITHUB_ISSUE="https://github.com/pan93412/CopyLargeFileGo/issues"

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

if [[ -e clf.out ]]; then
  rm clf.out
fi


echo "[INFO] 開始編譯：Starting build."
go build -o clf.out main.go string.go libs.go

if [[ -e clf.out ]]; then
  echo "[INFO] 編譯完成，檔案為 clf.out。"
  echo "[INFO] Build completed, the file is clf.out."
else
  echo "[ERR] 編譯可能失敗，請將錯誤訊息貼到 GitHub 上的 Issue 
Tracker。"
  echo "[ERR] The build probably failed, please paste the error message"
  echo "[ERR] to Issue Tracker on GitHub."
  echo "[ERR] URL: ${GITHUB_ISSUE}"
fi
