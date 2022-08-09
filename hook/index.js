async function preInit(inputObj) {

}

async function postInit(inputObj) {
    console.log(`\n    Welcome to use start-image-watermark
     * 部署前执行请安装go
         go  : https://golang.google.cn/doc/install
     * 项目初始化完成，您可以直接进入项目目录下，并使用 s deploy 进行项目部署\n`)
}

module.exports = {
    postInit,
    preInit
}
