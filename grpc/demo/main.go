package main

import (
	"context"
	"log"
	"net"

	"github.com/zoumas/lab/grpc/demo/invoicer"
	"google.golang.org/grpc"
)

type ProdInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *ProdInvoiceServer) Create(
	ctx context.Context,
	r *invoicer.CreateRequest,
) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &ProdInvoiceServer{}
	invoicer.RegisterInvoicerServer(server, service)
	log.Fatal(server.Serve(ln))
}
