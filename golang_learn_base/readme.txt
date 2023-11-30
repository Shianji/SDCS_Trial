golang包：
包可以区分命名空间（一个文件夹下面不能有两个同名文件）。go中创建一个包，一般是创建
一个文件夹，在该文件夹里面的go文件中，使用package关键字声明包名称，文件夹名称和包
名称相同。并且，同一个文件下面只有一个包。
go mod：
1.11版本之后用来管理包的依赖关系:
go mod init <项目模块名称>:初始化一个mod
go mod tidy :可以分析项目代码和依赖关系，然后删除不再需要的依赖项，从而减小项目的依赖关系，提高构建效率。
            还可以更新依赖项的版本，以确保你的项目使用的是最新的稳定版本。还可以修复模块文件 (go.mod)，
            以确保它们与项目实际的依赖关系一致。这有助于防止模块文件与代码之间的不一致，从而减少潜在的问题。
go list -m all：显示依赖关系
go list -m -json all：显示详细依赖关系
go mod download [path@version]：下载依赖
go build：编译和构建Go项目的关键命令，将Go代码转化为可执行文件或共享库，使得能够在不同的平台上运行应用程序，同时也有助于代码质量的维护。

对于网上的包，可以先去官网下载，然后在本地导入import

go语言的os包中有对目录和文件读写的一些方法，如：os.Mkdir、os.Remove、os.Getwd、os.WirteFile、os.ReadFile、os.Create、os.Open、os.OpenFile、
    f.Read、f.Write、f.Close等等；还有对进程操作的一些方法，如：os.Getpid、os.Getppid、os.FinProcess、p.Pid、p.Wait、p.Signal等等；有关环境
    变量的一些方法，如：os.Environ返回所有的环境变量、os.Getenv(环境变量名)获取某个环境变量的值、os.Setenv设置环境变量的值、os.LookupEnv查找环境变量等等
go语言的其他库如io、ioutil(提供了一些用于文件和 I/O 操作的便捷函数)、bufio(通过缓冲区来读写)、builtin(提供了一些类型声明、还有一些便利函数，这个包不需要导入，
    这些变量和函数可直接使用)、log(实现简单的日志服务,如打印日志信息，抛出异常等)、bytes(提供了一些对字节切片进行读写操作的函数)、errors(实现了操作错误的函数)、
    sort(提供了切片和用户自定义数据集以及相关功能的函数，如排序)、time(提供测量和显示时间的函数)、json(实现json格式的转换)、encoding/xml(实现xml格式的转换)、
    math(包含一些用于数学计算的函数，如三角函数、随机数、绝对值等)