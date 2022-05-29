package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Zeldoso17/API-GO-REDSOCIAL/bd"
	"github.com/Zeldoso17/API-GO-REDSOCIAL/models"
)

// UploadAvatar is a function that allows to upload an banner
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUser+ "." + extension
	status, err = bd.ModifyRegister(usuario, IDUser)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
