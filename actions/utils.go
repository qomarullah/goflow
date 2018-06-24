package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Fields struct {
	Id     string `json:"id"`
	Var    string `json:"var"`
	Fn     string `json:"fn"`
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
	Neg    string `json:"neg"`

	Opr   string   `json:"opr"`
	Ds    string   `json:"ds"`
	Url   string   `json:"url"`
	Out   string   `json:"out"`   //internal state
	State string   `json:"state"` //internal state
	True  []Fields `json:"true"`
	False []Fields `json:"false"`
}

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}

	return string(bytes)
}

func (f Fields) ToString() string {
	return ToJson(f)
}

func Populate(x string, infos map[string]string) (y string) {

	y = x
	if strings.HasPrefix(x, "[") && strings.HasSuffix(x, "]") {
		xx := strings.TrimPrefix(x, "[")
		xx = strings.TrimSuffix(xx, "]")
		y = infos[xx]
		fmt.Println("populate-res", "x:"+xx, "y:"+y)
		//looping
		if strings.HasPrefix(x, "[") && strings.HasSuffix(x, "]") {
			Populate(y, infos)
		}
	}
	return
}
