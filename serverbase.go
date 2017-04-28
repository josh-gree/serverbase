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

type Go struct {
	Time float64 `json:"time"`
}

var n = map[string]string{"sum":"n2","prod":"n1"}

func Sum(data []float64) float64{
	sum := 0.0
	for _,d := range data {
		sum += d
	}
	return sum
}

func StartP(c echo.Context) error{
	g := Go{}

	err := c.Bind(&g)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%#v\n",g)

	return nil
}

func Send(m Message) error{
	// data, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	//
	// fmt.Println("JSON to send;")
	// fmt.Println(string(data))

	g := Go{
		Time : m.Time,
	}
	data, err := json.Marshal(g)

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

func GetP(c echo.Context) error{
	m := Message{}

	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(Sum(m.Data))

	Send(m)


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
		e.Start(":8000")
}
