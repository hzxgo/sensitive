# sensitive
敏感词检测

## 使用方法
cd $GOPATH/src/sensitive
go build
./sensitive -c "我要去东京"

output:
WARN[0000] input: [我要去东西东京哦] is not allow, sensitive_word: 东京  [Fn]="sensitive/main.go:23"