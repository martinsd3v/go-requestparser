package parser

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Name     string
	Age      int
	Weigth   float32
	Tags     []string
	Vehicles []struct {
		Model string
		Plate string
		Brand struct {
			Name string
		}
	}
	Doc struct {
		Cpf    string
		Tags   []string
		Social struct {
			Number string
			Tags   []string
		}
	}
	Date        time.Time
	SimpleTypes struct {
		Si       int
		Si8      int8
		Si16     int16
		Si32     int32
		Si64     int64
		Siu      uint
		Siu8     uint8
		Siu16    uint16
		Siu32    uint32
		Siu64    uint64
		Sib      bool
		Sif32    float32
		Sif64    float64
		Sip      *int
		Zeroi    int
		Zerou    uint
		Zerof    float32
		Nulli    int
		Nullu    uint
		Nullf    float32
		Nullb    bool
		Duration time.Duration
	}
}

func Test(t *testing.T) {

	var data Data
	var data2 Data

	request := mockRequest()
	request2 := mockRequest2()

	Parser(request, &data)
	Parser(request2, &data2)

	assert.Equal(t, data, data2)
}

func mockRequest() *http.Request {

	form := url.Values{}
	form.Add("Name", "Marcelo")
	form.Add("Age", "23")
	form.Add("Weigth", "23.58")
	form.Add("Tags[]", "dev")
	form.Add("Tags[]", "go")
	form.Add("Vehicles[0][Model]", "Model S")
	form.Add("Vehicles[0][Plate]", "CWD4442")
	form.Add("Vehicles[0][Brand][Name]", "tesla")
	form.Add("Doc[Cpf]", "756585")
	form.Add("Doc[Tags][]", "idoc1")
	form.Add("Doc[Tags][]", "idoc2")
	form.Add("Doc[Cpf][Social][Number]", "756585")
	form.Add("Doc[Cpf][Social][Tags][]", "isoc1")
	form.Add("Doc[Cpf][Social][Tags][]", "isoc2")
	form.Add("Doc[Cpf][Social][Tags][]", "2")
	form.Add("Date", "2021-05-15")
	form.Add("SimpleTypes[Si]", "1")
	form.Add("SimpleTypes[Si]", "1")
	form.Add("SimpleTypes[Si8]", "1")
	form.Add("SimpleTypes[Si16]", "1")
	form.Add("SimpleTypes[Si32]", "1")
	form.Add("SimpleTypes[Si64]", "1")
	form.Add("SimpleTypes[Siu]", "1")
	form.Add("SimpleTypes[Siu8]", "1")
	form.Add("SimpleTypes[Siu16]", "1")
	form.Add("SimpleTypes[Siu32]", "1")
	form.Add("SimpleTypes[Siu64]", "1")
	form.Add("SimpleTypes[Sib]", "false")
	form.Add("SimpleTypes[Sif32]", "1.5")
	form.Add("SimpleTypes[Sif64]", "1.5")
	form.Add("SimpleTypes[Sip]", "1")
	form.Add("SimpleTypes[Zeroi]", "0")
	form.Add("SimpleTypes[Zerou]", "0")
	form.Add("SimpleTypes[Si]", "0")
	form.Add("SimpleTypes[Zerof]", "0")
	form.Add("SimpleTypes[Duration]", "5s")
	form.Add("SimpleTypes[Nulli]", "")
	form.Add("SimpleTypes[Nullu]", "")
	form.Add("SimpleTypes[Nullf]", "")
	form.Add("SimpleTypes[Nullb]", "")

	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return request
}

func mockRequest2() *http.Request {
	json := `{"Name":"Marcelo","Age":23,"Weigth":23.58,"Tags":["dev","go"],"Vehicles":[{"Model":"Model S","Plate":"CWD4442","Brand":{"Name":"tesla"}}],"Doc":{"Cpf":"756585","Tags":["idoc1","idoc2"],"Social":{"Number":"756585","Tags":["isoc1","isoc2","2"]}},"Date":"2021-05-15T00:00:00Z","SimpleTypes":{"Si":1,"Si8":1,"Si16":1,"Si32":1,"Si64":1,"Siu":1,"Siu8":1,"Siu16":1,"Siu32":1,"Siu64":1,"Sib":false,"Sif32":1.5,"Sif64":1.5,"Sip":null,"Zeroi":0,"Zerou":0,"Zerof":0,"Nulli":0,"Nullu":0,"Nullf":0,"Nullb":false,"Duration":5000000000}}`
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	request.Header.Set("Content-Type", "application/json")
	return request
}
