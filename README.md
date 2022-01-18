# validator
Replace the validation framework used by gin

replace mod:`replace github.com/go-playground/validator/v10 vx.x.x => git.kuainiujinke.com/bc-backend/nb/validator v0.0.1`

example: 
```
package handler

import (
	"errors"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

type TestForm struct {
	Name string `json:"name"`
}

func (f TestForm) Validate() (err error) {
	if f.Name == "" {
		err = errors.New("name is empty")
	}
	return
}

func TestHandler(ctx *gin.Context) {
	form := &TestForm{}
	err := ctx.ShouldBind(form)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		ctx.String(http.StatusOK, "OK")
	}
}
```
