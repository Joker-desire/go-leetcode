# go-leetcode
leetcode-golang刷题

## 将本地Go-leetcode项目上传到GitHub中
### 1、在GitHub上创建一个空项目
### 2、在本地初始化一个新的本地仓库
```shell
git init --initial-branch=master
```
注：从Git2.28版本开始，`git init`命令初始化的新仓库默认分支名为`main`，不再是`master`。
但是可以通过使用`--initial-branch=master`选项来创建旧的`master`分支。
### 3、将本地仓库与GitHub远程仓库关联起来
```shell
git remote add origin https://github.com/Joker-desire/go-leetcode.git
```
### 4、将要推送的文件或者目录添加到本地仓库中
```shell
git add .
```
### 5、提交更改
```shell
git commit -m "Go LeetCode项 目 初 始 化 "
```
### 6、将本地仓库的更改推送到远程仓库中
```shell
git push -u origin master
```