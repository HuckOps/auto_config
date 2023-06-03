# 配置自动更新

![img.png](img.png)

# 使用

本插件支持所有go版本，引入方式：

```shell
go get github.com/HuckOps/auto_config
```

# 用例

```go
func TestConfig(t *testing.T) {
	var Dest map[string]interface{}
	w := func() {
		fmt.Println(Dest)
	}
	config, err := NewConfig(WithSource(file.NewSource(file.WithPath("./test/test.yaml"))), WithEntity(&Dest), WithCallback(w))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	config.Watcher()
	select {}
}
```

参考go-micro的config模块进行例简化和改造：

1. 更精简的代码.
2. 更短的同步链和更少的内存消耗.
3. 增加保底，如配置文件异常移动或者编辑异常无法解析时抛出异常但不退出监听，输出扔保持上一次解析结果，可支持热修改，文件修改正常即可重新进行解析。

采用例go-micro的options思想，函数参数大部分都设计成例可变长度参数，即可以导入过个配置文件。多个配置文件会根据先后顺序抢占式填充输出结构体，支持触发式更新，动态更新代码对象（如数据库引擎）。

设计理念来自于go-micro。模块由多个携程工作，其中有多个携程进行文件监听，代码接收到fsnotify事件时触发更新，读取器获取文件通过编码器转换为json格式后放入装载器，并传送更新后的快照到配置监听器。监听器使用统一json解码器转为结构体或interface{}。

# 当前缺陷

1. 未捕获致命错误，可能会代码异常，可能需要使用信号方式结束进程。
2. fsnotify组件在某些环境下可能会接受到多个事件，作者回复可能无解，需要代码中进一步改造。
