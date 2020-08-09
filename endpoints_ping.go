package parmesan

import (
	"encoding/json"
	"fmt"
)

func Ping() (result PingResponse, err error){
	fmt.Println(hash([]byte{}))
  response, err := doApiGet([]string{"ping"})
  if err == nil {
    json.Unmarshal(response, &result)
  }
  return result, err
}
