package chatService
import("github.com/go-kit/kit/endpoint")
type CreateUserRequest struct{
FirstName string `json:"firstName"`
LastName string `json:"lastName"`
Email string `json:"email"`
ProfilePic string `json:"profilePic"`

}
type  CreateUserResponse struct{
Id string `json:"id"`

}
 func makeCreateUserEndpoint(s ChatService)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(CreateUserRequest)
id,error:=s.CreateUser(req.FirstName,req.LastName,req.Email,req.ProfilePic)
return CreateUserResponse{Id: id,}, error
}
}
type GetUserRequest struct{
Id string `json:"id"`

}
type  GetUserResponse struct{
FirstName string `json:"firstName"`
LastName string `json:"lastName"`
Email string `json:"email"`
ProfilePic string `json:"profilePic"`

}
 func makeGetUserEndpoint(s ChatService)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(GetUserRequest)
firstName,lastName,email,profilePic,error:=s.GetUser(req.Id)
return GetUserResponse{FirstName: firstName,LastName: lastName,Email: email,ProfilePic: profilePic,}, error
}
}
type UpdateUserRequest struct{
Id string `json:"id"`
ProfilePic string `json:"profilePic"`

}
type  UpdateUserResponse struct{
Message string `json:"message"`

}
 func makeUpdateUserEndpoint(s ChatService)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(UpdateUserRequest)
message,error:=s.UpdateUser(req.Id,req.ProfilePic)
return UpdateUserResponse{Message: message,}, error
}
}
type Endpoints struct{
 CreateUser endpoint.Endpoint
GetUser endpoint.Endpoint
UpdateUser endpoint.Endpoint
}
func MakeEndpoints(s chatService)Endpoints{
 return Endpoints{
CreateUser:makeCreateUser(s),
GetUser:makeGetUser(s),
UpdateUser:makeUpdateUser(s),
}
}