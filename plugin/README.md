# README


参考资料：

- https://mp.weixin.qq.com/s/2eSjBksWWO0IE2Pp4JTkDg
- [go-plugin-example](https://github.com/vladimirvivien/go-plugin-example) - 多语言插件示例

## 教程

1. 编译插件 so 文件

```bash
go build -buildmode=plugin ./lower_plugin
```

获得编译后的 `lower_plugin.so`。

2. 加载运行

```bash
go run .
```