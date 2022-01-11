package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	r.GET("/", Anasayfa)      //GET ile sunucuya gönderiyoruz.
	r.POST("/deneme", Deneme) // POST ile sunucuda olanı çekiyoruz.

	http.ListenAndServe(":8080", r) //sunucuyu ayağa kaldırma.

}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html") // .html dosyasını parse ediyoruz yani main.go nun erişmesini sağlıyoruz.

	view.Execute(w, nil) // sunucuya göndermek için resposewriteri kullanıyoruz.

}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 << 20)                                      // eklenebilecek dosya boyutunu belirtiyoruz.
	file, header, _ := r.FormFile("file")                               // index.html de ki file adlı girişe eklenen dosyayı çekiyoruz.
	F, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666) // bulunduğumuz dizinde dosya oluşturma ben oluşturmadım bulunduğum dizini dosya olarak gösterdim.
	io.Copy(F, file)                                                    // file değişkeninden gelen dosyayı F değişkenine kopyalıyoruz yani bulunduğumuz dizine
}
