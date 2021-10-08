#!/usr/bin/env bash
cd "$( dirname "${BASH_SOURCE[0]}" )"

SDIR=services
function waiting () {
    sleep 3
}
# 增加文件
echo "增加文件1"
cp test.yaml $SDIR/test1.yaml

waiting

# echo "增加文件2"
# touch $SDIR/test2.yaml
# waiting

# # 修改文件
# echo "修改文件2  时间"
# touch $SDIR/test2.yaml
# waiting
# echo "修改文件1  内容"
# echo >> $SDIR/test1.yaml
# waiting
# echo "修改文件2  内容"
# echo >> $SDIR/test2.yaml
# waiting

# # 删除文件
# echo "移走文件2"
# mv $SDIR/test2.yaml ./
# waiting
# echo "移回文件2"
# mv ./test2.yaml $SDIR/
# waiting
# echo "删除文件2"
# rm $SDIR/test2.yaml
# waiting

echo "删除文件1"
rm $SDIR/test1.yaml

