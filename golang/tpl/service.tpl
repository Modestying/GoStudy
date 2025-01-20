#!/bin/sh
####################
# author: SgDevops #
####################
#
# chkconfig: 345 85 15
# Default-Start: 2 3 4 5
# Default-Stop: 0 1 2 3 4 6
# Required-Start:
# description: {{.Description}}
# processname: {{.Name}}
# pidfile:
# lockfile:
# Source function library.
. /etc/rc.d/init.d/functions
# Source networking configuration.
. /etc/sysconfig/network

binName={{.Name}}
myDir={{.Path}}/{{.Name}}
myBin={{.Path}}/{{.Name}}/bin/{{.Name}}
myConf={{.Path}}/{{.Name}}/conf/{{.Config}}.yaml

# Check dir and file.
if [ ! -f $myBin ]; then
        echo "no bin file"
        exit 0
fi

if [ ! -f $myConf ]; then
        echo "no conf file"
        exit 0
fi


app_checkStartStatus(){
     procpid=`pidof $binName`
     if [ "x$procpid" == "x" ]; then
        echo "$binName start failed"
    else
        echo "$binName start success, pid:" $procpid
    fi
}

app_start() {
     procpid=`pidof $binName`
     if [ "x$procpid" == "x" ]; then
        cd $myDir
        ulimit -c unlimited && nohup $myBin -f $myConf  &

        sleep 1
        app_checkStartStatus
    else
        echo "binName already active, pid:" $procpid
    fi
}

app_stop() {
     procpid=`pidof $binName`
     if [ "x$procpid" == "x" ]; then
        echo "$binName is already dead"
    else
        kill -9 $procpid
        echo "$binName is stopped"
    fi
}

app_restart() {
    app_stop
    sleep 1
    app_start
}

app_status() {
    procpid=`pidof $binName`
    if [ "x$procpid" == "x" ]; then
        echo "$binName is dead"
    else
        echo "$binName is alive, pid:" $procpid
    fi
}

case "$1" in
    start)
    app_start
    ;;
    stop)
    app_stop
    ;;
    restart)
    app_restart
    ;;
    status)
    app_status
    ;;
    *)
    echo $"Usage: $0 {start|stop|status|restart}"
    exit 2
esac