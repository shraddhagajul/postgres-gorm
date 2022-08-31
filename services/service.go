package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"postgres/models"
	"strconv"
)

var dbConn *gorm.DB

type Response struct {
	Data    []models.Post `json:"data"`
	Message string        `json:"message"`
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts = models.GetPosts()
	var resp Response
	err := dbConn.Find(&posts).Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println(posts)
	resp.Data = posts
	resp.Message = "SUCCESS"
	json.NewEncoder(w).Encode(&resp)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	post := models.GetPost()
	err := dbConn.Where("id = ?", id).Find(&post).Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println(post)
	resp.Data = append(resp.Data, post)
	resp.Message = "SUCCESS"
	json.NewEncoder(w).Encode(&resp)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	post := models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	log.Println(post)

	err := dbConn.Create(&post).Error
	if err != nil {
		http.Error(w, "Error creating error ", 400)
		return
	}
	resp.Message = "CREATED"
	json.NewEncoder(w).Encode(resp)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	var resp Response
	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	id, _ := strconv.Atoi(params["id"])

	err := dbConn.Model(&post).Where("id = ?", id).Update(&post).Error
	if err != nil {
		http.Error(w, "Error creating record", 400)
		return
	}

	resp.Message = "UPDATED"
	json.NewEncoder(w).Encode(resp)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	post := models.GetPost()
	//soft delete
	err := dbConn.Delete(&post, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	resp.Message = "DELETED"
	json.NewEncoder(w).Encode(resp)
}

func SetDB(db *gorm.DB) {
	dbConn = db
	var post = models.GetPost()
	db.AutoMigrate(&post)
}
