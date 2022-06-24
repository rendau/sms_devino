package cmd

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"
	"github.com/rendau/sms_devino/internal/adapters/server/rest"
	"github.com/rendau/sms_devino/internal/domain/core"
)

func Execute() {
	// var err error

	confLoad()

	app := struct {
		lg         *dopLoggerZap.St
		core       *core.St
		restApiSrv *dopServerHttps.St
	}{}

	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	app.core = core.New(
		app.lg,
		httpclient.New(app.lg, httpc.OptionsSt{
			Client: &http.Client{
				Timeout: 20 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
			BaseUrl:       "https://api.devino.online",
			BaseLogPrefix: "Devino: ",
			BasicAuthCreds: &httpc.BasicAuthCredsSt{
				Username: conf.DevinoUsername,
				Password: conf.DevinoPassword,
			},
		}),
		conf.SenderName,
	)

	// START

	app.lg.Infow("Starting")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.core,
			conf.HttpCors,
		),
		app.lg,
	)

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}
