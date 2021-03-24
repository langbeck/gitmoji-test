package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	raw "google.golang.org/api/storage/v1"
	htransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
)

func main() {
	var opts []option.ClientOption
	userAgent := "gcloud-golang-storage/v1"

	// Prepend default options to avoid overriding options passed by the user.
	opts = append([]option.ClientOption{option.WithScopes(raw.DevstorageFullControlScope), option.WithUserAgent(userAgent)}, opts...)

	opts = append(opts, internaloption.WithDefaultEndpoint("https://storage.googleapis.com/storage/v1/"))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint("https://storage.mtls.googleapis.com/storage/v1/"))

	hc, ep, err := htransport.NewClient(context.Background(), opts...)
	if err != nil {
		panic(err)
	}
	log.Println("Endpoint: ", ep)

	log.Println(hc.Get("https://madserver-hl2sj7izaa-uc.a.run.app"))

	log.Println("Dialing #1...")
	c, err := net.Dial("tcp", "madserver-hl2sj7izaa-uc.a.run.app:443")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// tc := tls.Client(c, &tls.Config{InsecureSkipVerify: true})
	// log.Println("Handshake #1...")
	// err = tc.Handshake()
	// if err != nil {
	// 	panic(err)
	// }

	log.Println("Dialing #2...")
	conn, err := grpc.Dial(
		"madserver-hl2sj7izaa-uc.a.run.app:443",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	log.Println("OK")
}
