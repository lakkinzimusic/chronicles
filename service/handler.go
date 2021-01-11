package service

import (
	"fmt"
	"github.com/lakkinzimusic/chronicles/model"
	pb "github.com/lakkinzimusic/chronicles/proto"
	"golang.org/x/net/context"
)


type ChronicleService interface {
	Create(context.Context, *pb.Chronicle)  (*pb.Response, error)
	All(context.Context, *pb.Filter)  (*pb.Response, error)
}

func NewChronicleService(store Store) ChronicleService {
	return &Handler{
		store: store,
	}
}

type Handler struct {
	store Store
}

func (s *Handler) Create(c context.Context, req *pb.Chronicle) (*pb.Response, error) {
	res := pb.Response{}
	item, err := s.store.Create(c, model.MarshalChronicle(req))
	if err != nil {
		return &res, err
	}
	res.Chronicle = model.UnmarshalChronicle(item.(*model.Chronicle))
	return &res, nil
}

func (s *Handler) All(c context.Context, filter *pb.Filter)  (*pb.Response, error) {
	res := pb.Response{}
	items, err := s.store.All(c)
	if err != nil {
		return &res, err
	}
	fmt.Println(4)
	for _, item := range items {
		res.Chronicles = append(res.Chronicles, model.UnmarshalChronicle(item))
	}
	fmt.Println(5)
	return &res, nil
}
