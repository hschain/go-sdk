package server

import (
	"github.com/gin-gonic/gin"
)

func NewServer(listenAddress string, lcd string) *Server {
	return &Server{
		listenAddress,
		lcd,
	}
}

func (s *Server) cros(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "x-auth-token, content-type")
	c.Header("Access-Control-Expose-Headers", "x-auth-token")
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-Frame-Options", "DENY")
	c.Header("Vary", "Origin")
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.Header("Connection", "keep-alive")
}

func (s *Server) Router() error {
	router := gin.Default()
	nvr := router.Group("/hschain")
	nvr.Use(s.cros)

	nvr.GET("/createmnemonic", s.CreateMnemonic)
	nvr.GET("/getaccount", s.QueryAddressAccounts)
	nvr.GET("/getlastheight", s.GetLastHeight)
	nvr.GET("/getheighttxinfo", s.GetHeightTxInfo)
	nvr.GET("/getaddresstxinfo", s.GetAddressTxInfo)
	nvr.GET("/gethashtxinfo", s.GetHashTxInfo)
	nvr.POST("/transferaccounts", s.TransferAccounts)
	nvr.POST("/transferdestory", s.TransferDestory)

	router.Run(s.ListenAddress)
	return nil
}
