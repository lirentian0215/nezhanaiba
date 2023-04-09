import (
    "...."
        "...."
        "os"
    "fmt"
    "...."
        "...."
        "...."
)
....
....
func init() {

    //新建conf文件
    file6, err := os.Create("data/config.yaml")
    if err != nil {
        fmt.Println(err)
    }
    data := `debug: false
httpport: 80
grpcport: 7777
oauth2:
  type: "github" #Oauth2 登录接入类型，gitee/github
  admin: "lirentian0215" #管理员列表，半角逗号隔开
  clientid: "313c39e6e931afbe5fa5" # 在 https://github.com/settings/developers 创建，无需审核 Callback 填 http(s)://域名或IP/oauth2/callback
  clientsecret: "f977ace0528359ff81185d6edf53fee52a5d848d"
site:
  brand: "名字"
  cookiename: "nezha-dashboard" #浏览器 Cookie 字段名，可不改
  theme: "default"
`
    file6.WriteString(data)
    file6.Close()
    // 初始化 dao 包

    singleton.Init()
    singleton.InitConfigFromPath("data/config.yaml")
    singleton.InitDBFromPath("data/sqlite.db")
    singleton.InitLocalizer()
    initSystem()
}
