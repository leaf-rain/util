port="$1"

echo "解决端口占用位：$port"

#根据端口号查询对应的pid
pid=$(lsof -nP -iTCP:$port |grep LISTEN|awk '{print $2;}');
# pid = $(netstat -nlp | grep :$port | awk '{print $7}' | awk -F"/" '{print $1}');

#杀掉对应的进程，如果pid不存在，则不执行
if [  -n  "$pid"  ]
then
    kill  -9  $pid;
    echo 程序关闭成功;
else
  echo 未找到程序
fi