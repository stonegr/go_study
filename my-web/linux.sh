curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
    -H "Content-Type:text/plain; charset=utf-8" \
    --form 'file=@"./1.txt"' \
    --form 'name="bwg"'

curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
    -F 'file=@"./1.txt"' \
    -F 'name="bwg"'

curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
    -F 'file=@"./老师意见--20230517.xlsx"' \
    -F 'name="bwg"'

curl --location 'http://127.0.0.1:8080/go-api/server_back' \
    -F 'file=@"./app.exe"' \
    -F 'name="bwg"'

curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
    -F 'file=@"./app.exe"' \
    -F 'name="bwg"' \
    -F 'pwd="bqweasd123"'

curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
    -F 'file=@"./w64devkit-1.19.0.zip"' \
    -F 'name="bwg"'

function Uploadfile() {
    <<EOF
    $1 文件的路径
    $2 客户端的标识
    $3 客户端的密钥
EOF
    curl --location 'http://127.0.0.1:8080/go-api/server_back' -X POST \
        -F 'file=@"'"$1"'"' \
        -F 'name="'""$2""'"' \
        -F 'pwd="'"$3"'"'
}
