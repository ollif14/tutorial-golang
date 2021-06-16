package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-grpc-mongo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"golang-grpc-mongo/config"
	"golang-grpc-mongo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type EmployeeServiceServer struct {
	proto.UnimplementedEmployeeServiceServer
}

func (s *EmployeeServiceServer) GetAll(req *proto.NoRequest, stream proto.EmployeeService_GetAllServer) error{
	data := &model.Employee{}

	result, err := config.Employeedb.Find(context.Background(), bson.M{})
	if err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	defer result.Close(context.Background())

	for result.Next(context.Background()){
		if err := result.Decode(data); err != nil{
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		res := &proto.EmployeeResponse{
			E: &proto.Employee{
				EId: data.E_Id,
				EmailAddress: data.Email_address,
				FirstName: data.First_name,
				LastName: data.Last_name,
			},
		}
		stream.Send(res)
	}
	if err := result.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}

	return nil
}

func (s *EmployeeServiceServer) GetById(ctx context.Context, req *proto.EmployeeIdRequest) (*proto.EmployeeResponse, error){
	id := req.GetEId()
	result := config.Employeedb.FindOne(ctx, bson.M{"e_id" : id})
	data := &model.Employee{}

	if err := result.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find employee: %v", err))
	}

	res := &proto.EmployeeResponse{
		E: &proto.Employee{
			EId: data.E_Id,
			EmailAddress: data.Email_address,
			FirstName: data.First_name,
			LastName: data.Last_name,
		},
	}
	return res, nil
}

func (s *EmployeeServiceServer) Create(ctx context.Context, req *proto.EmployeeRequest) (*proto.EmployeeResponse, error){
	employee := req.GetE()
	data := model.Employee{
		Id: bson.NewObjectId(),
		E_Id: employee.GetEId(),
		Email_address: employee.GetEmailAddress(),
		First_name: employee.GetFirstName(),
		Last_name: employee.GetLastName(),
	}
	
	_, err := config.Employeedb.InsertOne(config.MongoCtx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err),
			)
	}

	res :=  &proto.EmployeeResponse{E: employee}
	return res, nil
}

func (s *EmployeeServiceServer) Update(ctx context.Context, req *proto.EmployeeRequest) (*proto.EmployeeResponse, error){
	employee := req.GetE()
	id := employee.GetEId()

	update := bson.M{
		"email_address": employee.GetEmailAddress(),
		"first_name": employee.GetFirstName(),
		"last_name": employee.GetLastName(),
	}

	idBsonConv := bson.M{"e_id": id}
	result := config.Employeedb.FindOneAndUpdate(ctx, idBsonConv, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))
	
	data := &model.Employee{}
	err := result.Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find employee: %v", err),
		)
	}

	res := &proto.EmployeeResponse{
		E: &proto.Employee{
			EId: data.E_Id,
			EmailAddress: data.Email_address,
			FirstName: data.First_name,
			LastName: data.Last_name,
		},
	}
	return res, nil
}

func (s *EmployeeServiceServer) Delete(ctx context.Context, req *proto.EmployeeIdRequest) (*proto.SuccessResponse, error){
	id := req.GetEId()
	_, err := config.Employeedb.DeleteOne(ctx, bson.M{"e_id" : id})
	if err != nil{
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not delete employee: %v", err))
	}

	res := &proto.SuccessResponse{
		Deleted: true,
	}
	return res, nil
}

func main()  {
	fmt.Println("start server on port :50051")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("unable to listen on port :50051: %v", err)
	}

	s := grpc.NewServer()
	srv := &EmployeeServiceServer{}

	proto.RegisterEmployeeServiceServer(s, srv)
	config.ConnectToMongo()

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


