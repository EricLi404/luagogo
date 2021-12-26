



```shell

# 编译
luac -o hello_world.luac.out  hello_world.lua

# 反编译
luac -l hello_world.luac.out
# 反编译 显示详细信息
luac  -l -l hello_world.luac.out

# 反编译也可以直接对 lua 文件进行操作
luac -l hello_world.lua
```