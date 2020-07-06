package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type cacheGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createCacheGenerator(s model.Service, outputFile string) fileGenerator {
	return &cacheGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (cg *cacheGenerator) run(outputPath string) {
	log.Println("run")
	cg.generateCode()
	cg.generateFile(outputPath)
}
func (cg *cacheGenerator) GetFileName() string {
	return cg.outputFile
}

func (cg *cacheGenerator) generateCode() {

	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", cg.s.GetServiceName())
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
	cg.code = code.String()

}
func (cg cacheGenerator) generateFile(outputPath string) {

	if cg.s.RedisCache.GetHost() == "" {
		return
	}
	var path string
	path = fmt.Sprintf("%s/%s.go", outputPath, cg.outputFile)
	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(cg.code)
	defer file.Close()

}
