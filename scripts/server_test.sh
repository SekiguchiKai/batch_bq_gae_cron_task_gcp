#! /bin/sh

cur=`pwd`
# コマンドライン引数で、directory指定
cd $1
testdirs=`ls -d */`
STAT=0

for testdir in $testdirs
do
  if [ $testdir = "vendor/" ] || [ $testdir = "modules/" ]; then
    continue
  fi
  cd $testdir
  goapp test -v -cover ./
  if [ $? -ne 0 ]; then
    echo "Test is failed in $testdir"
    STAT=1
  fi
  # コマンドライン引数で、directory指定
  cd $1
done
cd $cur

if [ $STAT -eq 1 ]; then
  echo "Server Test is failed"
  exit $STAT
fi