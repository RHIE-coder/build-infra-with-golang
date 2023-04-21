# Lua

## [ Install ]

### - Ubuntu

```sh
sudo apt install build-essential libreadline-dev
sudo apt install lua
sudo apt install luarock
```

### - Build 

#### https://www.lua.org/download.html

```sh
curl -R -O http://www.lua.org/ftp/lua-5.4.4.tar.gz
tar zxf lua-5.4.4.tar.gz
cd lua-5.4.4
make all test
```