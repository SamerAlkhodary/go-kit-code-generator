# README
## Dependency:
The program needs dep to download all the packages needed in the code

## How to use: 
1. Fill a (serviceName).yaml file
2. run gen.sh script and follow instructions

## .yaml:
1. The yaml file should contain the name of the service.
2. The yaml file should contain at least one endpoint in order to generate the code
3. The transport method and path should be provided for each and every endpoint inorder to generate the code.

## The generated files:
1. Transport file that contains all the code needed inroder to have a rest api
2. Encoders file that is needed to help the transport layer to decode and encode the informations
3. Service file that contains all the code needed to run the service
4. Server file that contains the Serve function that need to be called to run the program
5. model file that contains everything specifed in the model section of the yaml file ( it has setters and getters)
6. Endpoint file taht contains all the endpoints specified in the yaml file
   
## Future features: 
The ability to choose whether a repository is needed. If so the repository will be generated  and connected to the service.

## Example of a .yaml file:
```yaml
name: chatService
endpoints:
  -
    name: CreateUser
    args: firstName string , lastName string, email string, profilePic string
    output: id string
    transport: 
      method: POST
      path: /user
  -
    name: GetUser
    args: id string 
    output: firstName string, lastName string, email string, profilePic string
    transport:
      method: GET
      path: /user/{id}
  -
    name: UpdateUser
    args: id string, profilePic string
    output: message string
    transport:
      method: PUT
      path: /user/update
model:
  -
    name: Person
    attr: firstName string, lastName string , profilePic string , token string
  -
    name: Event
    attr: name string, startDate string, endDate string, location Location
  -
    name: Location
    attr: longitude float64, latitude float64
    

  

```