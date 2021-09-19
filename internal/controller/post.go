package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weekend.side/SocialMedia/internal/daos"
	"weekend.side/SocialMedia/internal/infra/db"
	"weekend.side/SocialMedia/internal/services"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// authenticate request
	var post daos.Post
	var err *daos.Error
	post.Creator, err = services.AuthenticateRequest(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	parseErr := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if parseErr != nil {
		http.Error(w, "failed to parse multipart message", http.StatusBadRequest)
		return
	}

	post.Caption = r.FormValue("caption")
	file, _, _ := r.FormFile("image_url")
	reqBytes, _ := ioutil.ReadAll(file)

	post.ImageUrl = getMediaFileUrl(r)

	json.Unmarshal(reqBytes, &post.Caption)

	resp, respErr := services.CreatePost(post)
	if respErr != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(respErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func CommentOnPost(w http.ResponseWriter, r *http.Request) {
	comment := daos.Commment{}
	var err *daos.Error
	comment.User, err = services.AuthenticateRequest(r)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readErr)
		return
	}

	unmarshalErr := json.Unmarshal(bodyBytes, &comment)
	if unmarshalErr != nil {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(unmarshalErr)
		return
	}

	resp, respErr := services.CommentOnPost(&comment)
	if respErr != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	var post daos.Post
	var err *daos.Error
	post.Creator, err = services.AuthenticateRequest(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	resp, fetchErr := services.FetchPost()
	if fetchErr != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, fetchErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func getMediaFileUrl(r *http.Request) string {
	var image_url string
	for _, h := range r.MultipartForm.File["image_url"] {
		file, _ := h.Open()
		image_url = h.Filename
		image_url, _ = db.UploadToS3(file, h.Filename)
		//tmpfile, _ := os.Create("./" + image_url)
		//io.Copy(tmpfile, file)
		//tmpfile.Close()
		//file.Close()
	}
	return image_url
}
