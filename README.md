# Copy Large File (CLF)
這是一個能遞迴複製資料夾、或只單一複製檔案的一個程式。

## 命令列
```
=! Copy Large File != 版本 v0.7_beta
用法：(程式名稱) (來源檔案位置) (目標檔案位置) [-r] [-v] [--check-{devel|stable}]
() 為必須、[] 為選用。

[-r] (可選)：遞迴複製模式 (若原始檔案位置為 目錄，則必須)
[-v] (可選)：詳細輸出模式
[--check-devel] (可選)：檢查 devel 分支上的更新
[--check-stable] (可選)：檢查 master 分支上的更新

此軟體 GitHub 網址：<http://www.github.com/pan93412/CopyLargeFileGo>
```

## 速度
```
[pan93412@archlinux ~]$ du -sh ~/下載
824M	/home/pan93412/下載

[pan93412@archlinux CopyLargeFile]$ time ./clf.out ~/下載 ~/test_folder -r

real  0m16.591s
user  0m0.209s
sys  0m4.684s
```

## 如何編譯
```
/*
 * 編譯方式：  go build -o clf.out main.go string.go libs.go
 * (Windows) go build -o clf.exe main.go string.go libs.go
 * 如果是 Linux，則直接輸入 bash build.sh 即可。
 */
```

## 作者
- pan93412 \<<https://github.com/pan93412>\>, 2018.
- Yami Odymel (技術指導) \<<https://github.com/yamiodymel>\>, 2018.
- 和其他朋友們 ;-)
