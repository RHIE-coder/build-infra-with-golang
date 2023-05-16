# /lib64/libc.so.6: version 'GLIBC_2.25' not found

 - https://github.com/directus/directus/issues/6071

```sh
---------------------------------
**gcc 8.2.0 :** 
$ gcc -v
$ yum -y install wget bzip2 gcc gcc-c++ glibc-headers
download package ：https://mirrors.tuna.tsinghua.edu.cn/gnu/gcc 
or
$ wget -c -P /software/ https://mirrors.tuna.tsinghua.edu.cn/gnu/gcc/gcc-8.2.0/gcc-8.2.0.tar.gz

$ cd /software/
$ tar -zxvf gcc-8.2.0.tar.gz

$ cd gcc-8.2.0
$ ./contrib/download_prerequisites   
result：
gmp-6.1.0.tar.bz2: OK
mpfr-3.1.4.tar.bz2: OK
mpc-1.0.3.tar.gz: OK
isl-0.18.tar.bz2: OK
All prerequisites downloaded successfully.

$ mkdir build
$ cd bulid
$ ../configure --prefix=/usr/local/gcc-8.2.0 --enable-bootstrap --enable-checking=release --enable-languages=c,c++ --disable-multilib

$ make 
My system execute make command at least 3 hours, please waiting it complete.

then: 
$ make install

$ echo -e '\nexport PATH=/usr/local/gcc-8.2.0/bin:$PATH\n' >> /etc/profile.d/gcc.sh && source /etc/profile.d/gcc.sh

$ ln -sv /usr/local/gcc-8.2.0/include/ /usr/include/gcc

$  ldconfig -v    

$ ldconfig -p |grep gcc
result: 
libgcc_s.so.1 (libc6,x86-64) => /lib64/libgcc_s.so.1

check version 
$  gcc -v 

------------------------------------
**make 4.2.1:**
First download make 4.2.1 from https://ftp.gnu.org/gnu/make/
$ make -v 
$ tar -zxvf make-4.2.1.tar.gz
$ cd make-4.2.1
$ mkdir build
$ cd build
$ ../configure --prefix=/usr
$ sh build.sh
$ make install
check result 
$make -v

----------------------------------------
**bison 3.0.4**
$ bison -V
if not found result ,need install:

$ yum install -y bison

--------------------------------------------------------
all complete , you can install glibc-2.28

download glibc-2.28 from http://ftp.gnu.org/gnu/glibc. 
put it into /software

You can adopt the following steps：

$ tar -xf glibc-2.28.tar.gz
$ cd glibc-2.28
$ mkdir build
$ cd build
$ ../configure --prefix=/usr --disable-profile --enable-add-ons --with-headers=/usr/include --with-binutils=/usr/bin

$ make
$ make install 
$ ls -l /lib64/libc.so.6

Last, check version : 
$ strings /lib64/libc.so.6 | grep GLIBC
.....
GLIBC_2.26
GLIBC_2.27
GLIBC_2.28
GLIBC_PRIVATE

It completed .

Other : 
bootstarp Directus , maybe get error: 
Import Error: /usr/lib64/libstdc++.so.6: version `CXXABI_1.3.8' not found, you reed update CXXABI : 

$ strings /usr/lib64/libstdc++.so.6 | grep 'CXXABI'

download:[ libstdc++.so.6.0](https://www.02405.com/uploads/soft/201124/1-2011241A414.zip)
put it into /usr/lib64 

$  cd /usr/lib64
$ rm -rf libstdc++.so.6
$ ln -s libstdc++.so.6.0.22 libstdc++.so.6
```