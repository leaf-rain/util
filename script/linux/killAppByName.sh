pName="$1";

echo "停止应用程序：$pName";

#根据端口号查询对应的pid
pid=$(ps -ef | grep $pName | grep -v grep | grep -v kill | awk '{print $2}');
echo "pid=$pid"

#杀掉对应的进程，如果pid不存在，则不执行
if [ "$pid" != "" ]
then
  kill -15 $pid
  echo 程序关闭
else
  echo 未找到程序
fi
