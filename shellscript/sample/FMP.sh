#!/bin/bash
#------------------------------
#  [기본 기능 명세]                
# 1. 파일삭제 : deleteFile()
# 2. 파일이동 : moveFile()
# 3. 디렉토리삭제 : deleteDir()
# 4. 디렉토리 이동 : moveDir()
# 5. 파일 내용 출력하기 : printFileContent()
# 6. 프로그램 종료
#------------------------------
#  [추가 기능 명세]    
# 1. 웹 개발 초기 셋팅 : makeWebSample()
# 2. 현재 위치 파일 및 폴더 목록 출력: readPwdList()-->printPwdList()
# 3. 메뉴얼 출력 기능 : printHelp()
# 4. 주요 실행 부 분리와 환영 메시지 : showWelcome() + mainProgram()
#------------------------------
# by 이민형
#------------------------------

function printHelp(){
    echo "Usage:"
    echo -e "   \033[1;36mhelp\033[0m    프로그램 설명 출력"
    echo -e "   \033[1;36mlist\033[0m    현재 위치 파일 및 폴더 목록 출력"
    echo -e "   \033[1;36mweb\033[0m     웹 개발 초기 셋팅 진행"
    echo " "
    echo -e "   \033[1;36mrmf\033[0m   [삭제할 파일 이름]"
    echo -e "   \033[1;36mmvf\033[0m   [이동할 파일 이름]"
    echo -e "   \033[1;36mshow\033[0m  [파일 내용 확인]"
    echo -e "   \033[1;36mrmd\033[0m   [삭제할 디렉토리 이름]"
    echo -e "   \033[1;36mmvd\033[0m   [이동할 디렉토리 이름]"
    echo " "
    echo -e "   \033[1;36mend\033[0m     프로그램 종료"
    echo "----------------------------------------"
}

function deleteFile(){
    if [ -e $1 ] && [ $# -ne 0 ];then
        rm $1
        if [ $? -ne 0 ];then
            echo -e "\033[31m[ERROR]\033[0m\033[1;32mrmf\033[0m명령어 실행중 에러가 발생했습니다. 파일 이름과 위치를 확인해주세요."
        fi
    else
        echo -e "\033[31m[ERROR]\033[0m현재 위치에서 해당 하는 파일이 존재하지 않습니다."
    fi
}

function moveFile(){
    if [ -e $1 ] && [ $# -ne 0 ] && [ ! -d $1 ];then
        read -p "어디로 이동하시겠습니까?(경로가 아닌 이름만 입력하면 파일 이름이 바뀝니다.)  >> " moveToHere
        mv $1 $moveToHere
        if [ $? -ne 0 ];then
            echo -e "\033[31m[ERROR]\033[0m\033[1;32mmvf\033[0m명령어 실행중 에러가 발생했습니다. 파일 이름과 위치를 확인해주세요."
        fi
    else
        echo -e "\033[31m[ERROR]\033[0m현재 위치에서 해당 하는 파일이 존재하지 않습니다."
    fi
    unset moveToHere
}

function deleteDir(){
    if [ -d $1 ] && [ $# -ne 0 ];then
        rm -r $1
        if [ $? -ne 0 ];then
            echo -e "\033[31m[ERROR]\033[0m\033[1;32mrmd\033[0m명령어 실행중 에러가 발생했습니다. 디렉토리 이름과 위치를 확인해주세요."
        fi
    else
        echo -e "\033[31m[ERROR]\033[0m현재 위치에서 해당 하는 폴더가 존재하지 않습니다."
    fi
}

function moveDir(){
    if [ -d $1 ] && [ $# -ne 0 ];then
        read -p "어디로 이동하시겠습니까?(경로가 아닌 이름만 입력하면 디렉터리 이름이 바뀝니다.)  >> " moveToHere
        mv $1 $moveToHere
        if [ $? -ne 0 ];then
            echo -e "\033[31m[ERROR]\033[0m\033[1;32mmvd\033[0m명령어 실행중 에러가 발생했습니다. 디렉토리 이름과 위치를 확인해주세요."
        fi
    else
        echo -e "\033[31m[ERROR]\033[0m현재 위치에서 해당 하는 폴더가 존재하지 않습니다."
    fi
    unset moveToHere
}

function printFileContent(){
    if [ -e $1 ] && [ $# -ne 0 ];then
        cat $1
        if [ $? -ne 0 ];then
            echo -e "\033[31m[ERROR]\033[0m\033[1;32mshow\033[0m명령어 실행중 에러가 발생했습니다. 파일 이름과 위치를 확인해주세요."
        fi
    else
        echo -e "\033[31m[ERROR]\033[0m현재 위치에서 해당 하는 파일이 존재하지 않습니다."
    fi
}

function readPwdList(){
    echo "============현재목록==========="
    while read line  
    do  
    echo "  |  $line"
    done
    echo "==============================="
}

function printPwdList(){
    ls -al | readPwdList
}

function makeWebSample(){

    HTML_CODE='<!DOCTYPE html>\n<html lang="en">\n<head>\n    <meta charset="UTF-8">\n    <meta name="viewport" content="width=device-width, initial-scale=1.0">\n    <title>Document</title>\n</head>\n<body>\n    \n</body>\n</html>'
    
    mkdir web-content
    if [ $? -ne 0 ];then
        echo -e "\033[31m[ERROR]\033[0m\033[1;32mweb-content\033[0m디렉토리가 이미 존재합니다."
        echo -e "\033[31m[ERROR]\033[0m해당 디렉토리의 이름을 바꾸거나 지우고 다시 시도해주세요."
    else
        echo -e $HTML_CODE > ./web-content/index.html
        mkdir -p web-content/css/
        touch web-content/css/style.css
        mkdir -p web-content/js
        touch web-content/js/script.js
        echo -e "\033[1;32m성공했습니다.\033[0m"
    fi

    unset HTML_CODE   
}

function showWelcome(){
    echo "----------------------------------------"
    echo " "
    echo "  _____ ___ _     _____    ______ ___ ____  _____ ____ _____ ___  ______   __"
    echo " |  ___|_ _| |   | ____|  / /  _ \_ _|  _ \| ____/ ___|_   _/ _ \|  _ \ \ / /"
    echo " | |_   | || |   |  _|   / /| | | | || |_) |  _|| |     | || | | | |_) \ V /"
    echo " |  _|  | || |___| |___ / / | |_| | ||  _ <| |__| |___  | || |_| |  _ < | |"
    echo " |_|   |___|_____|_____/_/  |____/___|_| \_\_____\____| |_| \___/|_| \_\|_|"
    echo " "
    echo "  __  __    _    _   _    _    ____ _____ __  __ _____ _   _ _____"
    echo " |  \/  |  / \  | \ | |  / \  / ___| ____|  \/  | ____| \ | |_   _|"
    echo " | |\/| | / _ \ |  \| | / _ \| |  _|  _| | |\/| |  _| |  \| | | |"
    echo " | |  | |/ ___ \| |\  |/ ___ \ |_| | |___| |  | | |___| |\  | | |"
    echo " |_|  |_/_/   \_\_| \_/_/   \_\____|_____|_|  |_|_____|_| \_| |_|"
    echo " "
    echo "  ____  ____   ___   ____ ____      _    __  __"
    echo " |  _ \|  _ \ / _ \ / ___|  _ \    / \  |  \/  |"
    echo " | |_) | |_) | | | | |  _| |_) |  / _ \ | |\/| |"
    echo " |  __/|  _ <| |_| | |_| |  _ <  / ___ \| |  | |"
    echo " |_|   |_| \_\\\___/ \____|_| \_\/_/   \_\_|  |_|"
    echo " "
    echo ""
    echo "----------------------------------------"
}

function mainProgram(){
    case $1 in
    help )
        printHelp
        ;;
    list )
        printPwdList
        ;;
    rmf )
        deleteFile $2
        ;;
    mvf )
        moveFile $2
        ;;
    rmd )
        deleteDir $2
        ;;
    mvd )
        moveDir $2
        ;;
    show )
        printFileContent $2
        ;;
    web  )
        makeWebSample
        ;;
    end )
        echo -e "\033[1;35m프로그램을 종료합니다.\033[0m"
        exit -1
        ;;
    * )
        echo -e "\033[31m[ERROR]\033[0m잘못된 명령어를 입력하셨습니다."
        echo "----------------------------------------"
        printHelp
        ;;
    esac
}

showWelcome
printHelp

while :
do
    read -p "F.M.P>" input
    echo $input
    mainProgram $input
done