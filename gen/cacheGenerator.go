package gen

import (
	"fmt"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

func cacheGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", s.GetServiceName())
	fmt.Fprintf(&code, "\nimport (\n%q\n%q\n%q\n%q\n)", "time", "github.com/go-redis/redis/v8", "encoding/json", "context")
	fmt.Fprintf(&code, "\ntype Cache interface{")
	fmt.Fprintf(&code, "\nSet(ctx context.Context, key string,value interface{},cacheTime time.Duration)error")
	fmt.Fprintf(&code, "\nGet(ctx context.Context, key string) (string,error)")
	fmt.Fprintf(&code, "\n}")
	fmt.Fprintf(&code, "\ntype redisCache struct{")
	fmt.Fprintf(&code, "\nhost string")
	fmt.Fprintf(&code, "\npassword string")
	fmt.Fprintf(&code, "\ndb int")
	fmt.Fprintf(&code, "\n}")
	fmt.Fprintf(&code, "\nfunc MakeNewRedisCache(host string, password string, db int)Cache{")
	fmt.Fprintf(&code, "\nreturn & redisCache{")
	fmt.Fprintf(&code, "\nhost: host,\ndb: db,\n password: password,\n}")
	fmt.Fprintf(&code, "\n}")
	fmt.Fprintf(&code, "\nfunc(r *redisCache)getClient()*redis.Client{")
	fmt.Fprintf(&code, "\nreturn redis.NewClient(\n &redis.Options{\nAddr: r.host,\n Password: r.password,\nDB: r.db,\n})")
	fmt.Fprintf(&code, "\n}")
	fmt.Fprintf(&code, "\nfunc(r *redisCache)Set(ctx context.Context, key string, value interface{},cacheTime time.Duration)error{")
	fmt.Fprintf(&code, "\nclient:=r.getClient()")
	fmt.Fprintf(&code, "\njson,err:= json.Marshal(value)")
	fmt.Fprintf(&code, "\nif err!=nil{\nreturn err\n}")
	fmt.Fprintf(&code, "\nclient.Set(ctx, key, json, cacheTime*time.Millisecond )")
	fmt.Fprintf(&code, "\nreturn nil\n}")
	fmt.Fprintf(&code, "\nfunc(r *redisCache)Get(ctx context.Context, key string) (string,error){")
	fmt.Fprintf(&code, "\nclient:=r.getClient()")
	fmt.Fprintf(&code, "\nvalue,err:= client.Get(ctx, key).Result()\n")
	fmt.Fprintf(&code, "\nif err!=nil{\nreturn %s,err\n}", `""`)
	fmt.Fprintf(&code, "\nreturn value,nil\n}")

	return code.String()
}
