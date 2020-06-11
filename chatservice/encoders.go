package chatService
import(
"encoding/json"
 "context"
"github.com/gorilla/mux"
"net/http")
func encodeResponse(ctx context.Context, w http.ResponseWriter,response interface{})error{
return json.NewEncoder(w).Encode(response)
}
func decodeCreateUserRequest(ctx context.Context, r *http.Request)(interface{},error){
var request CreateUserRequest
err:=json.NewDecoder(r.Body).Decode(&request)
if err!=nil{
 return nil,err
}
 return request,nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request)(interface{},error){
var request GetUserRequest
 vars:= mux.Vars(r)
 request= GetUserRequest{
 Id: vars["id"]}
return request,nil
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request)(interface{},error){
var request UpdateUserRequest
err:=json.NewDecoder(r.Body).Decode(&request)
if err!=nil{
 return nil,err
}
 return request,nil
}
