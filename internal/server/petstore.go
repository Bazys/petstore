package server

import (
	"context"

	"petstore/pkg/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type petstoreServer struct {
	pb.UnimplementedPetstoreServer
}

func newPetstoreServer() pb.PetstoreServer {
	return &petstoreServer{}
}

func (s *petstoreServer) AddPet(ctx context.Context, req *pb.AddPetRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) CreateUsersWithArrayInput(ctx context.Context, req *pb.CreateUsersWithArrayInputRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) CreateUsersWithListInput(ctx context.Context, req *pb.CreateUsersWithListInputRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) DeletePet(ctx context.Context, req *pb.DeletePetRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) FindPetsByStatus(ctx context.Context, req *pb.FindPetsByStatusRequest) (*pb.Pet, error) {
	panic("not implemented")
}

func (s *petstoreServer) FindPetsByTags(ctx context.Context, req *pb.FindPetsByTagsRequest) (*pb.Pet, error) {
	panic("not implemented")
}

func (s *petstoreServer) GetInventory(ctx context.Context, req *emptypb.Empty) (*pb.GetInventoryOK, error) {
	panic("not implemented")
}

func (s *petstoreServer) GetOrderById(ctx context.Context, req *pb.GetOrderByIdRequest) (*pb.Order, error) {
	panic("not implemented")
}

func (s *petstoreServer) GetPetById(ctx context.Context, req *pb.GetPetByIdRequest) (*pb.Pet, error) {
	panic("not implemented")
}

func (s *petstoreServer) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.User, error) {
	panic("not implemented")
}

func (s *petstoreServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) LogoutUser(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.Order, error) {
	panic("not implemented")
}

func (s *petstoreServer) UpdatePet(ctx context.Context, req *pb.UpdatePetRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) UpdatePetWithForm(ctx context.Context, req *pb.UpdatePetWithFormRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}

func (s *petstoreServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.ApiResponse, error) {
	panic("not implemented")
}

func (s *petstoreServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	panic("not implemented")
}
