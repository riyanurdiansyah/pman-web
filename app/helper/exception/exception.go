package exception

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrorPrint(err error) {
	if err != nil {
		fmt.Println("Error : " + err.Error())
	}
}

func HandleErrorRedirect(ctx *gin.Context, url string, err error) {
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, url)
		ctx.Abort()
		return
	}
}

// Fungsi global untuk menangani error
func HandleError(ctx *gin.Context, url string, err error, customMessage string) {
	if err != nil {
		RenderPage(ctx, url, nil, customMessage)
		ctx.Abort()
		return
	}
}

// Fungsi untuk merender halaman error
func RenderPage(ctx *gin.Context, url string, data interface{}, errorMsg string) {
	tmpl, err := template.ParseFiles(url)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(ctx.Writer, data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}
}
