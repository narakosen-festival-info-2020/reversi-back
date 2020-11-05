package api

import (
	"sync"
	"time"

	"github.com/gin-contrib/cors"
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

func (info *Info) generateCustomMatch(generateReversi func() (reversi.Data, error)) (Token, error) {
	info.dataMutex.Lock()
	defer info.dataMutex.Unlock()
	nowReversi, err := generateReversi()
	if err != nil {
		return Token{}, err
	}
	nowToken := generateToken()
	info.matchInfo[nowToken.specificCode] = &nowReversi
	info.tokenState = append(info.tokenState, nowToken)
	return nowToken, nil
}

func (info *Info) generateMatch(generateData *reversi.GenerateData) (Token, error) {
	return info.generateCustomMatch(generateData.Create)
}

func (info *Info) eraseToken() {
	for {
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
	server := gin.Default()

	// CORS setup
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://reversi.nitncfes.net"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	go serverInfo.eraseToken()
	setRoute(server, &serverInfo)
	server.Run(":80")
}
