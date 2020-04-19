SET PATH=%PATH%;%GOPATH%\bin;

REM 配置Qt目录和Qt版本
SET QT_DIR=C:\Qt\Qt5.12.5
SET QT_VERSION_MAJOR=5.12.5

REM 编译32位程序
SET GOARCH=386

REM 默认使用Mingw编译，使用MSVC编译的话开启下面的选项，qtsetup可以运行，但是后面qtdeploy会报cgo相关错误
REM SET QT_MSVC=true
REM SET PATH=%PATH%;C:\Program Files (x86)\Microsoft Visual Studio\2017\Community\VC\Auxiliary\Build;
REM call vcvarsall.bat amd64_x86

REM 语言绑定安装，只需执行一次
REM qtsetup.exe

REM 编译具体项目时运行
qtdeploy build desktop ./

