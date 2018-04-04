# cities
中国省市区街道社区5级联动

### 数据文件
`data` 文件夹中包含有各 json 文件。默认有省( `1.json` )市( `??.json` )区( `????.json` )三级，第4级和第5级文件个数太多，压缩后上传的，若需要街道级，解压 `streets.zip` 到 data 同级文件夹。若需要社区级，解压 `villages.zip` 到 data 同级文件夹。
最终 data 文件夹中，只包含 `index.html` 和各种 json 文件(或者还有 `streets.zip` 和 `villages.zip` ) ，没有下级文件夹，即可。

### 配置
配置文件在 `config.ini` 里
只需配置
```
directory = json 文件所在文件夹，本例中为 data
port = 服务的端口
```
即可。

### 编译运行
使用 `go build Server.go` 来编译。
启动服务 `./Server` 或 `./Server &`  后台运行

### 演示
根下直接就有 `index.html` ，可以查看演示，如
[http://localhost:8080/index.html](http://localhost:8080/index.html)
如果你没有解压 `streets.zip` 和 `villages.zip` ，那么，第4级和第5级不会出现。

#### 题外
* 完全可以不用提供的服务器，直接把 `data` 文件夹拿回去用即可。随便放在哪个项目的某个文件夹中，直接用。
* 为使演示尽量简洁，我使用了非常原始的方法来请求 `json` 文件，并且，没有进行任何异常处理，您应该视情况加上。
* 数据来源为 [https://github.com/modood/Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China) ，针对数据进行了适当加工。