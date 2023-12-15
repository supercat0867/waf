package graceful

import "fmt"

func Welcome() {
	wafAsciiArt := `
                               .--.,
         .---.               ,--.'  \
        /. ./|               |  | /\/
     .-'-. ' |    ,--.--.    :  : :
    /___/ \: |   /       \   :  | |-,
 .-'.. '   ' .  .--.  .-. |  |  : :/|
/___/ \:     '   \__\/: . .  |  |  .'
.   \  ' .\      ," .--.; |  '  : '
 \   \   ' \ |  /  /  ,.  |  |  | |
  \   \  |--"  ;  :   .'   \ |  : \
   \   \ |     |  ,     .-./ |  |,'
    '---"       '--''---'     '--'
	`

	fmt.Println(wafAsciiArt)
	//fmt.Printf("WAF后台地址：%s:%d/waf/login\n", config.Server, config.Port)
	//fmt.Println("初始账户：admin\n初始密码：123456")
}
