package serverbase

import (
	"fmt"
	"github.com/labstack/echo"
)

type Message struct {
	Dest int `json:"dest"`
	Time float64 `json:"time"`
}

type Go struct {
	Go bool `json:"go"`
}

func Start(c echo.Context) error{
	g := Go{}

	err := c.Bind(&g)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%#v\n",g)

	return nil
}

func Get(c echo.Context) error{
	m := Message{}

	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%#v\n",m)

	return nil
}
func Listen(port int){
		e := echo.New()
		fmt.Printf("Starting to listen on %d",port)
		e.POST("/start", Start)
		e.POST("/msg", Get)
		e.Start(fmt.Sprintf(":%d",port))
}
