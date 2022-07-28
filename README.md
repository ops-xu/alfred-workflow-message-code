# alfred-workflow-message-code
alfred workflow 通过imessage.app 提取验证码到粘贴板.
提取最新10条验证码短信(需短信内容包含`验证码`三个字)
仅支持提取数字4-8位验证码. 
## 使用
直接下载[workflow文件](https://github.com/vincentXu97/alfred-workflow-message-code/releases/download/v1.0.0/2fa-Extract-verification-code.alfredworkflow), 确认已经给alfred全部磁盘权限. 导入workflow后, 输入2fa即可查看验证码

## 原理
通过sqlite3读取Imessage的chat.db数据库, 筛选包含(验证码)的短信, 通过正则匹配4-8位数字验证码.

<img width="723" alt="CleanShot 2022-07-28 at 15 27 08" src="https://user-images.githubusercontent.com/28681228/181446490-2913f63e-4ab5-4263-9bd9-98293ce99ae9.png">
