## 运行代码

### 1. 下载代码

```
git clone https://github.com/jobsssss/mumu-web.git
```

### 2. 配置环境变量

```
cd mumu-web
cp .env.example .env
```

使用编辑器打开 .env 文件，并对里面的信息做相应配置，尤其是数据库信息。

### 3. 运行代码

```
go run .
```

### 4. 访问 mumu-web

http://localhost:3000/

### 5. 运行测试
```
go test ./tests/
```
运行指定的测试 ``go test -v -run TestAllPages pages_test.go``

测试文档参考: http://c.biancheng.net/view/124.html