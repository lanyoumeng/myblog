## myblog 项目
make

_output/myblog -h

任务：
1.走读 verflag 包的代码，弄明白在执行 miniblog --version 命令时，
--version 命令行选项的值是怎么赋值给 versionFlag 的？

requestID := ctx.Request.Header.Get(known.XRequestIDKey)
ctx.Writer.Header().Set(known.XRequestIDKey, requestID)

 c.Request.Method != "OPTIONS"


ctx context.Context


深拷贝