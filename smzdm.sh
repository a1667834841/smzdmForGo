#!/bin/bash
d=`date "+%y-%m/%d %H:%M:%S"`
#echo "$1$d"

depolyPath="/opt/go/smzdmForGo"

if [ $1 = "start" ]; then
echo "$d pusher is starting.... "
# <<EOF 
# source /etc/profile
# nohup ${depolyPath}/smzdmPusher >> ${depolyPath}/smzdm.log 2>&1 &
# EOF
./smzdmPusher

elif  [ $1 = "stop" ]; then

# 方法二：直接关闭进程kill  -9 [进程号]
smzdm_pid=`ps -ef |grep smzdmPusher | grep -v 'grep\|stop' | awk '{print $2}' `
kill -15 $smzdm_pid
echo "$d pusher was stoped "

elif [ $1 = "reload" ]; then
smzdm_pid=`ps -ef |grep smzdmPusher | grep -v 'grep\|stop' | awk '{print $2}' `
kill -15 $smzdm_pid
# <<EOF 
# source /etc/profile
# nohup ${depolyPath}/smzdmPusher >> ${depolyPath}/smzdm.log 2>&1 &
# EOF
./smzdmPusher
echo "$d pusher was reloaded "

else
echo "输入错误，请检查重新输入"
fi

