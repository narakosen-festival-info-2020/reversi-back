package api

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/narakosen-festival-info-2020/reversi-back/pkg/reversi"
)

// Info is Server State
type Info struct {
	serverUpTime time.Time
	matchInfo    map[string]*reversi.Data
	tokenState   []Token
	dataMutex    sync.Mutex
}

func (info *Info) generateCustomMatch(generateReversi func() reversi.Data) Token {
	info.dataMutex.Lock()
	defer info.dataMutex.Unlock()
	nowToken := generateToken()
	nowReversi := generateReversi()
	info.matchInfo[nowToken.specificCode] = &nowReversi
	fmt.Println(info.matchInfo[nowToken.specificCode])
	info.tokenState = append(info.tokenState, nowToken)
	return nowToken
}

func (info *Info) generateMatch(boardType string) (Token, error) {
	if boardType == reversi.NormalBoard {
		return info.generateCustomMatch(reversi.GenerateNormalReversi), nil
	}
	return Token{}, fmt.Errorf("Invalid Board Type")
}

func (info *Info) eraseToken() {
	for len(info.tokenState) != 0 {
		if !info.tokenState[0].IsExpire() {
			break
		}
		info.dataMutex.Lock()
		delete(info.matchInfo, info.tokenState[0].specificCode)
		info.tokenState = info.tokenState[1:]
		info.dataMutex.Unlock()
	}
	time.Sleep(time.Minute)
}

func (info *Info) getMatchData(specificCode string) (*reversi.Data, bool) {
	data, err := info.matchInfo[specificCode]
	if !err {
		return &reversi.Data{}, err
	}
	return data, err
}

// ServerUp is Start Server
func ServerUp() {
	serverInfo := Info{
		serverUpTime: time.Now(),
		matchInfo:    make(map[string]*reversi.Data),
	}
	router := gin.Default()
	go serverInfo.eraseToken()
	setRoute(router, &serverInfo)
	router.Run(":80")
}
