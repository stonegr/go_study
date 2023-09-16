#!/bin/bash

# 10 3 */3 * * /etc/code/conf/bfsc.sh bwg qweasd123

# 未用
UPLOAD_URL="https://cl.phonexyz.cf/go-api/server_back/"

UPLOAD_NAME=$1
UPLOAD_PWD=$2
BACK_UP_SCRIPT_DIR="/etc/code/conf"
BACK_UP_DIR="/etc/code/conf/back_up"

function TgTZ() {
    curl --location 'https://ly.stonexyz.cf/api/fs-tele' \
        --connect-timeout 5 \
        --header 'Content-Type: application/json' \
        --data '{
        "text":"'"${1}"'",
        "user": "1802232509",
        "pwd": "qweasd123"
    }'
}

function Uploadfile() {
    <<EOF
    $1 文件的路径
    $2 客户端的标识
    $3 客户端的密钥
EOF
    curl --location 'https://cl.phonexyz.cf/go-api/server_back/' -X POST \
        -F 'file=@"'"$1"'"' \
        -F 'name="'""$2""'"' \
        -F 'pwd="'"$3"'"'
}

cd $BACK_UP_SCRIPT_DIR
./back_up.sh -ab

# 获取最新的一个备份文件
file_path="$BACK_UP_DIR"/$(ls -ltr $BACK_UP_DIR | tail -n 1 | awk '{print $NF}')
Uploadfile "$file_path" "$UPLOAD_NAME" "$UPLOAD_PWD"
echo
