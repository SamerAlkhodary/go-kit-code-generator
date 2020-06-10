package chatservice
import("context"
 "github.com/go-kit/kit/log")
type chatservice interface{
SendMessage(senderId string , reciverId string, message string,ctx context.Context)(string,error)
SendImage(senderId string , reciverId string, img []byte,ctx context.Context)(string,error)
SendLocation(senderId string, recieverId string, location string,ctx context.Context)(string,error)
SendAudio(senderId string , reciverId string, sound []byte,ctx context.Context)(string,bool,error)
}
type chatservice struct{
repository Repository
logger log.Logger
}
func NewService(rep Repository,logger log.Logger)chatservice{
 return &chatservice{
 repository: rep,
 logger:logger,
}}
func(s *chatService)SendMessage(senderId string , reciverId string, message string,ctx context.Context)(string,error){
Logger:= log.With(s.logger,"method",SendMessage)
//TODO: implement
}
func(s *chatService)SendImage(senderId string , reciverId string, img []byte,ctx context.Context)(string,error){
Logger:= log.With(s.logger,"method",SendImage)
//TODO: implement
}
func(s *chatService)SendLocation(senderId string, recieverId string, location string,ctx context.Context)(string,error){
Logger:= log.With(s.logger,"method",SendLocation)
//TODO: implement
}
func(s *chatService)SendAudio(senderId string , reciverId string, sound []byte,ctx context.Context)(string,bool,error){
Logger:= log.With(s.logger,"method",SendAudio)
//TODO: implement
}
