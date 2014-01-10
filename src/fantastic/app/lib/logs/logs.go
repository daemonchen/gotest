package logs

import (
	"fmt"
	seelog "github.com/cihub/seelog"
)

var (
	Logger seelog.LoggerInterface
	err    error
)

func loadAppConfig() {
	appConfig := `
<seelog>
    <outputs>
        <filter levels="info">
            <smtp formatid="common" senderaddress="cgyqqcgyserver2@126.com" sendername="daemon" hostname="smtp.126.com" hostport="25" username="cgyqqcgyserver2@126.com" password="00o0o00oo00o0o">
                <recipient address="cgyqqcgy@126.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Msg" />
        <format id="critical" format="%File %FullPath %Func %Msg%n" />
        <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	Logger, err = seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	seelog.ReplaceLogger(Logger)
	defer seelog.Flush()
}

func init() {
	loadAppConfig()
}
