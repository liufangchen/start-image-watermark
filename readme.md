# start-image-watermark

<p align="center"><b> 中文 | <a href="./readme_en.md"> English </a>  </b></p>


> ***快速部署一个图片添加水印的应用到阿里云函数计算***

Serverless Devs的组件开发案例已经被集成到Serverless Devs命令行工具中，通过对Serverless Devs的命令行工具，可以进行空白应用项目的初始化，开发者只需要执行`s init`即可看到：

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess |   

## 部署 & 体验

- [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
- 初始化项目：`s init start-image-watermark -d start-image=watermark`   
- 进入项目，并进行项目部署：`cd start-image-watermark && s deploy -y`

## 应用详情

部署完成之后，您可以看到系统返回给您的案例地址，例如：

![](https://github.com/liufangchen/start-image-watermark/raw/main/image/b.png)


此时，打开案例地址，通过Get请求可以进行水印绘制：


| 参数 |  说明 | 缺省 |   
| --- |  --- | --- |  
| url | 图片路径 | https://pic.netbian.com/uploads/allimg/161001/095746-1475287066579f.jpg |    
| text | 水印 | watermark |    

例如：http://start-image-watermark.testservice.1401662146685254.cn-hangzhou.fc.devsapp.net/water?text=beautiful&&url=https://pic.netbian.com/uploads/allimg/220711/002225-16574701457687.jpg
