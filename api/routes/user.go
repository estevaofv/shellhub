package routes

import (
	"net/http"
	"fmt"

	"github.com/shellhub-io/shellhub/api/apicontext"
)

const (
	EditUserNameURL = "/userTeste"
)

func EditUserName(c apicontext.Context) error {

	fmt.Println("===========================================================")

	var user struct {
        User string `json:"user"`
    }

    if err := c.Bind(&user); err != nil {
        return err
	}
	fmt.Println(user)

	fmt.Println("===========================================================")
	

	return c.JSON(http.StatusOK, nil)
}
