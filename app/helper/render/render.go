package render

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderView(ctx *gin.Context, fileName string, data interface{}) {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Menggunakan `Execute` untuk merender template
	err = tmpl.Execute(ctx.Writer, data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
}
