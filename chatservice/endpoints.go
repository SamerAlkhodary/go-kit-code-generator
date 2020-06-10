package chatservice
import("github.com/go-kit/kit/endpoint")
type SendMessageRequest struct{
senderId string 
 reciverId string
 message string

}
type  SendMessageResponse struct{
message string

}
 func makeSendMessageEndpoint(s Service)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(SendMessageRequest)
message,error:=s.SendMessage(req.senderId,req.reciverId,req.message)
return SendMessageResponse{message: message,}, error
}
}
type SendImageRequest struct{
senderId string 
 reciverId string
 img []byte

}
type  SendImageResponse struct{
message string

}
 func makeSendImageEndpoint(s Service)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(SendImageRequest)
message,error:=s.SendImage(req.senderId,req.reciverId,req.img)
return SendImageResponse{message: message,}, error
}
}
type SendLocationRequest struct{
senderId string
 recieverId string
 location string

}
type  SendLocationResponse struct{
message string

}
 func makeSendLocationEndpoint(s Service)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(SendLocationRequest)
message,error:=s.SendLocation(req.senderId,req.recieverId,req.location)
return SendLocationResponse{message: message,}, error
}
}
type SendAudioRequest struct{
senderId string 
 reciverId string
 sound []byte

}
type  SendAudioResponse struct{
message string
ok bool

}
 func makeSendAudioEndpoint(s Service)endpoint.Endpoint{
return func(ctx context.Context, request interface{}) (interface{}, error) {
req := request.(SendAudioRequest)
message,ok,error:=s.SendAudio(req.senderId,req.reciverId,req.sound)
return SendAudioResponse{message: message,ok: ok,}, error
}
}
type Endpoints struct{
 SendMessage endpoint.Endpoint
SendImage endpoint.Endpoint
SendLocation endpoint.Endpoint
SendAudio endpoint.Endpoint
}
func MakeEndpoints(s chatService)Endpoints{
 return Endpoints{
SendMessage:makeSendMessage(s),
SendImage:makeSendImage(s),
SendLocation:makeSendLocation(s),
SendAudio:makeSendAudio(s),
}
}