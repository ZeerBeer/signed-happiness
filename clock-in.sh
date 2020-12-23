#!/bin/sh

set -eux

UA="Mozilla/5.0 (Linux; Android 10; MI CC 9 Build/QKQ1.190828.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36"

curl \
    -H "User-Agent: $UA" \
    -X POST \
    -d "wc_type=%E5%90%A6&title=36.5&province=%E5%B1%B1%E8%A5%BF%E7%9C%81&jk_type=%E5%81%A5%E5%BA%B7&jc_type=%E5%90%A6&is_verify=0&district=%E5%B0%96%E8%8D%89%E5%9D%AA%E5%8C%BA&city=%E5%A4%AA%E5%8E%9F%E5%B8%82&address=%E5%B1%B1%E8%A5%BF%E7%9C%81%E5%A4%AA%E5%8E%9F%E5%B8%82%E5%B0%96%E8%8D%89%E5%9D%AA%E5%8C%BA%E9%95%BF%E5%BE%81%E8%A1%971%E5%8F%B7&mobile=$1" \
    http://yx.ty-ke.com/Home/Monitor/monitor_add