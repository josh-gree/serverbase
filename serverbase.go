package serverbase

import (
	"fmt"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"bytes"
	"os"
)

type Message struct {
	Dest string `json:"dest"`
	Data []float64 `json:"data"`
	Time float64 `json:"time"`
}

type Done struct {
	Result float64 `json:"result"`
}

var n = map[string]string{"sum":"n2","prod":"n1"}

func Sum(data []float64) float64{
	sum := 0.0
	for _,d := range data {
		sum += d
	}
	return sum
}

func Prod(data []float64) float64{
	prod := 1.0
	for _,d := range data {
		prod *= d
	}
	return prod
}

var FuncMap  = map[string]func([]float64)float64{"sum":Sum,"prod":Prod}

func StartP(c echo.Context) error{
	m := Message{}

	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(FuncMap[m.Dest](m.Data))
	fmt.Printf("%#v\n",m)

	return nil
}

func Sendres(d Done) error{
	data, err := json.Marshal(d)

	_, err = http.Post("http://public:8000/done","application/json",bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func Sendjob(m Message) error{
	// data, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	//
	// fmt.Println("JSON to send;")
	// fmt.Println(string(data))

	data, err := json.Marshal(m)

	fmt.Println("JSON to send;")
	fmt.Println(string(data))
	fmt.Sprintf("%s:8000/start",n[m.Dest])
	resp, err := http.Post(fmt.Sprintf("http://%s:8000/start",n[m.Dest]),"application/json",bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	return nil
}

func DoneP(c echo.Context) error {
	d := Done{}

	err := c.Bind(&d)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Result: %f\n",d.Result)

	return nil

}

func GetP(c echo.Context) error{
	m := Message{}

	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	Sendjob(m)


	fmt.Printf("%#v\n",m)

	return nil
}
func Listen(master bool){
		e := echo.New()
		fmt.Printf("%#v\n",os.Getenv("PATH"))
		fmt.Printf("Starting to listen on %d\n",8000)
		if master {
			fmt.Println("Master process")
		} else {
			fmt.Println("Slave process")
		}
		e.POST("/start", StartP)
		e.POST("/msg", GetP)
		e.POST("/done", DoneP)
		e.Start(":8000")
}
