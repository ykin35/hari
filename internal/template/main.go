package template

var (
	MainFNC = `package main

import (
  log	"xinhari.com/xinhari/logger"
	"xinhari.com/xinhari"
	"{{.Dir}}/handler"
	"{{.Dir}}/subscriber"
)

func main() {
	// New Service
	function := xinhari.NewFunction(
		xinhari.Name("{{.FQDN}}"),
		xinhari.Version("latest"),
	)

	// Initialise function
	function.Init()

	// Register Handler
	function.Handle(new(handler.{{title .Alias}}))

	// Register Struct as Subscriber
	function.Subscribe("{{.FQDN}}", new(subscriber.{{title .Alias}}))

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
`

	MainSRV = `package main

import (
	log "xinhari.com/xinhari/logger"
	"xinhari.com/xinhari"
	"{{.Dir}}/handler"
	"{{.Dir}}/subscriber"

	{{dehyphen .Alias}} "{{.Dir}}/proto"
)

func main() {
	// New Service
	service := xinhari.NewService(
		xinhari.Name("{{.FQDN}}"),
		xinhari.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	{{dehyphen .Alias}}.Register{{title .Alias}}Handler(service.Server(), new(handler.{{title .Alias}}))

	// Register Struct as Subscriber
	xinhari.RegisterSubscriber("{{.FQDN}}", service.Server(), new(subscriber.{{title .Alias}}))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
	MainAPI = `package main

import (
	log "xinhari.com/xinhari/logger"

	"xinhari.com/xinhari"
	"{{.Dir}}/handler"
	"{{.Dir}}/client"

	{{dehyphen .Alias}} "{{.Dir}}/proto"
)

func main() {
	// New Service
	service := xinhari.NewService(
		xinhari.Name("{{.FQDN}}"),
		xinhari.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the {{title .Alias}} service client
		xinhari.WrapHandler(client.{{title .Alias}}Wrapper(service)),
	)

	// Register Handler
	{{dehyphen .Alias}}.Register{{title .Alias}}Handler(service.Server(), new(handler.{{title .Alias}}))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
	MainWEB = `package main

import (
        log "xinhari.com/xinhari/logger"
	      "net/http"
        "xinhari.com/xinhari/web"
        "{{.Dir}}/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("{{.FQDN}}"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/{{dehyphen .Alias}}/call", handler.{{title .Alias}}Call)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
`
)
