package exception

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fungsi global untuk menangani error
func HandleError(ctx *gin.Context, url string, err error, customMessage string) {
	if err != nil {
		RenderPage(ctx, url, customMessage)
		ctx.Abort()
		return
	}
}

// Fungsi untuk merender halaman error
func RenderPage(ctx *gin.Context, url string, errorMsg string) {
	tmpl, err := template.ParseFiles(url)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(ctx.Writer, map[string]string{"ErrorMessage": errorMsg})
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}
}
