# 通过goctl生成脚手架 --home 是你的模板（如果未对goctl的模板进行修改就不写，我这里修改了一些模板）
# userOpt.api中 import了多个api文件 所以最后只需要这一个文件
goctl api go -api userOpt.api -dir ../ --style=goZero --home=../../../tpl