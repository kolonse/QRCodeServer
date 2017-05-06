# 二维码服务说明 #
目前二维码服务属性房卡捕鱼项目,域名:qrcode.youdianle.com.cn

**1.服务启动参数介绍**
>  -a string
>  
>        -a=<url> 连接性活码地址url (default "http://qrcode.youdianle.com.cn/active")
>
>  -d string
>  
>        -d=<dir> 活码内容存储目录 (default "/tmp/simplekv")
>  -i string
>  
>        -i=<ip> default=0.0.0.0 (default "0.0.0.0")
>  -n string
>  
>        -n=<name> 活码存储目录名字 (default "qrcode")
>  -p int
>  
>        -p=<port> default=18001 (default 18001)

其中 -a,-n,-d 主要是针对二维码是活码的属性配置,活码主要用于：你的二维码内容较长,导致生成的二维码较为复杂,如果复杂的二维码缩小到手机上看有可能导致不能正常扫描,通过开启活码后,就会将内容进行缩短从而简化二维码内容

>-a 主要针对二维码是属于连接型的二维码,也就是扫描后能跳转到你的预期连接中,这个配置是生成这个跳转型二维码的连接属性,最好有个较为简短的域名,这样生成的二维码会非常简单,而且这个域名必须指向当前二维码服务的地址.

>-n 活码真实内容存储目录名字,如果一台机器上需要开启多个二维码服务,请确保 -n必须不一致,需要部署时特别注意,代码不会处理-n相同时的异常情况.

>-d 表示活码内容存储的根目录

**2.服务使用介绍**

服务请求地址
>     http://qrcode.youdianle.com.cn/qrcode

请求参数:
>     content(必须) 要进行生成二维码的内容,如果是带有参数的url时需要对这个url进行 urlencode
>     例如:http://qrcode.youdianle.com.cn/qrcode?content=http://xxx?a=1
>     其中 content 需要修改为 http%3a%2f%2fxxx%3fa%3d1,对应二维码如下:
![content 测试二维码](http://qrcode.youdianle.com.cn/qrcode?content=http%3a%2f%2fxxx%3fa%3d1)

>     size(默认 256) 生成二维码的图片大小,例如
![http://qrcode.youdianle.com.cn/qrcode?content=test&size=70](http://qrcode.youdianle.com.cn/qrcode?content=test&size=70)    
 
>     bgcolor(默认 纯白 ffffff) 背景颜色值,是16进制rgb值
>     例如:http://qrcode.youdianle.com.cn/qrcode?content=test&bgcolor=00ff00

![DIY背景颜色二维码](http://qrcode.youdianle.com.cn/qrcode?content=test&bgcolor=00ff00)

>     forecolor(默认 纯黑 000000) 前景颜色值,是16进制rgb值
>     例如:http://qrcode.youdianle.com.cn/qrcode?content=test&forecolor=00ff00

![DIY背景颜色二维码](http://qrcode.youdianle.com.cn/qrcode?content=test&forecolor=00ff00)

>     logo(默认 无) 插入二维码中间的图片url
>     特别注意,在设计二维码时,logo的大小和二维码的大小比例要提前设定好.
>     例如：http://qrcode.youdianle.com.cn/qrcode?content=test&logo=http://fk.youdianle.com.cn/logo.png

![DIY背景颜色二维码](http://qrcode.youdianle.com.cn/qrcode?content=test&logo=http://fk.youdianle.com.cn/logo.png)

>     bdmaxsize(默认 无)表示二维码周围的空白边宽大小,如果设置该值,那么生成的二维码就不一定是 size 属性大小,而是 bdmaxsize + (size - size % 43) / 43.
>     例如:http://qrcode.youdianle.com.cn/qrcode?content=test&bgcolor=00ff00&bgmaxsize=5

![DIY背景颜色二维码](http://qrcode.youdianle.com.cn/qrcode?content=test&bgcolor=00ffff&bdmaxsize=3)

>     shortest(默认 false,有效值只有 true)  表示是否生成最简二维码,主要针对于 content 比较长的情况,如果content比较长那么生成二维码的比较复杂,当二维码缩小时会很难扫描
>     例如：content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect
>     未配置 shortest 前:
![http://qrcode.youdianle.com.cn/qrcode?forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect](http://qrcode.youdianle.com.cn/qrcode?forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect)

>     配置 shortest=true后
![http://qrcode.youdianle.com.cn/qrcode?shortest=true&forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect](http://qrcode.youdianle.com.cn/qrcode?shortest=true&forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect)

扫描二维码后两个的内容是不一样的,第二个二维码的内容是：http://qrcode.youdianle.com.cn/active?k=G, 当我们访问这个连接后的内容就变成了 content的内容.其中 http://qrcode.youdianle.com.cn/active 可以通过服务启动时-a参数进行定制,但路径必须是/active

>     redirect(默认false,有效值只有 true) 前提必须 shortest=true时才有效,表示生成活码后的连接进行访问后是否对 content内容进行重定向.
>     例如将下面的活码配置redirect=true后,微信扫描应该是房卡捕鱼的付费页面

![http://qrcode.youdianle.com.cn/qrcode?redirect=true&shortest=true&forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay%23wechat_redirect](http://qrcode.youdianle.com.cn/qrcode?redirect=true&shortest=true&forecolor=076FC1&bdmaxsize=2&amp;content=https%3a%2f%2fopen.weixin.qq.com%2fconnect%2foauth2%2fauthorize%3fappid%3dwx2b54cdf28427315e%26redirect_uri%3dhttp%253a%252f%252fwxpublic.youdianle.com.cn%252fpay%252fjsapi%252fdo%253fpayAppId%253dfff0b507e9344b615213b0d8c457a956%26response_type%3dcode%26scope%3dsnsapi_userinfo%26state%3dwxpay_test%23wechat_redirect)



