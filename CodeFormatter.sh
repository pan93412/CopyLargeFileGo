#!/usr/bin/bash
# Copy Large File 的專用程式碼整理程式
# 在提交每個 PR 前請都先用此程式格式化
go fmt *.go                # 先使用 Go 內建的格式化軟體初步格式化
sed -i 's/\t/  /g' *.go    # 兩空格制
exit 0
