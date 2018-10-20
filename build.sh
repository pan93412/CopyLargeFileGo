#!/usr/bin/env bash
# Copy Large File (CLF) 編譯腳本
# 與版本：v0.7.1-beta 同步釋出
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
  echo "請安裝 go (>= 1.10.0) 版本後再繼續編譯。"
  echo "感謝！"
  echo -e "\nPlease build it after installed go (version >= 
1.10)!\nThanks too much!"
  exit 1
fi

if [[ -e clf.out ]]; then
  rm clf.out
fi

if [[ -d "build" ]]; then
  rm -rf "build"
fi
mkdir "build"

if [[ $1 == "--all" ]]; then
  echo "[INFO] 開始編譯：Starting build."
  GOOS=darwin GOARCH=amd64 go build -o build/clf-darwin-amd64.out *.go
  echo "[INFO] darwin amd64 編譯完成 Build Completed"
  SYSTEM="freebsd linux windows"

  ARCH="amd64"
  for i in ${SYSTEM}; do
    if [[ $i == "windows" ]]; then
      filext=".exe"
    else
      filext=".out"
    fi
    filename="clf-$i-$ARCH$filext"
    GOOS=$i GOARCH=amd64 go build -o build/$filename *.go
    echo "[INFO] $i $ARCH 編譯完成 Build Completed"
  done

  ARCH="386"
  for i in $SYSTEM; do
    if [[ $i == "windows" ]]; then
      filext=".exe"
    else
      filext=".out"
    fi
    filename="clf-$i-$ARCH$filext"
    GOOS=$i GOARCH=$ARCH go build -o build/$filename *.go
    echo "[INFO] $i $ARCH 編譯完成 Build Completed"
  done
  exit 0
fi
  

echo "[INFO] 開始編譯：Starting build."
go build -o build/clf.out *.go

if [[ -e build/clf.out ]]; then
  echo "[INFO] 編譯完成，檔案為 build/clf.out。"
  echo "[INFO] Build completed, the file is build/clf.out."
else
  echo "[ERR] 編譯可能失敗，請將錯誤訊息貼到 GitHub 上的 Issue 
Tracker。"
  echo "[ERR] The build probably failed, please paste the error message"
  echo "[ERR] to Issue Tracker on GitHub."
  echo "[ERR] URL: ${GITHUB_ISSUE}"
fi
