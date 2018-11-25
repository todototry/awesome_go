
# gitbook-summary
目录支持
## 安装 gitbook-summary
```shell
npm install -g gitbook-summary
```
## 使用 summary 生成目录
```shell
book sm
```
## 忽略指定章节，灰度显示
```shell
book sm -i 0home, 1-干嘛要写书
```
## 展开目录
```
book sm  -c 2-什么是即时出版平台？, 3-如何打造自己的平台？
```

# gitbook
## 安装
安装 gitbook
```shell
npm install gitbook -g
```
## 启动
将指定的./repository 目录作为 gitbook 要展示的内容，并提供 http 服务
```shell
gitbook serve ./repository
```
## 生成html
将./repository 目录的 markdown文件编译生成 html，写入 output 文件夹
```shell
gitbook build ./repository --output=./outputFolder
```

## 生成 pdf/epub/mobi
1. 下载安装calibre
    下载地址：https://calibre-ebook.com/download
2. 将其 ebook-convert工具放入 path。
    ```
    sudo ln -s /Applications/calibre.app/Contents/MacOS/ebook-convert /usr/bin
    ```
3. 执行转换

    ```shell
    gitbook pdf  ./  ./go.pdf
    ```

    ```shell
    gitbook mobi ./ ./go.mobi
    ```
