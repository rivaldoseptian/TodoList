package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/config"
	"server/helpers"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetActivity(w http.ResponseWriter, r *http.Request) {
	var activity []models.Activities

	if err := config.DB.Find(&activity).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Success", activity)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	idActivity := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idActivity)

	var activity models.Activities

	if err := config.DB.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, fmt.Sprintf("Activity with ID %v Not Found", id), nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return

	}

	helpers.Response(w, 200, "Sucess", activity)
}

func CrateActivity(w http.ResponseWriter, r *http.Request) {
	var activity models.Activities

	// if r.Body == nil {
	// 	helpers.Response(w, 400, "title cannot be null", nil)
	// 	return
	// }

	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if activity.Title == "" {
		helpers.Response(w, 400, "title cannot be null", nil)
		return
	}
	if activity.Email == "" {
		helpers.Response(w, 400, "email cannot be null", nil)
		return
	}

	if err := config.DB.Create(&activity).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 201, "Success", activity)
}

func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	idActivity := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idActivity)

	var activity models.Activities

	if err := config.DB.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, fmt.Sprintf("Activity with ID %v Not Found", id), nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if activity.Title == "" {
		helpers.Response(w, 400, "title cannot be null", nil)
		return
	}
	if activity.Email == "" {
		helpers.Response(w, 400, "email cannot be null", nil)
		return
	}

	if err := config.DB.Where("activity_id = ?", id).Updates(&activity).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 200, "Success", activity)
}

func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	idActivity := mux.Vars(r)["id"]
	activity_id, _ := strconv.Atoi(idActivity)

	var activity models.Activities

	res := config.DB.Delete(&activity, activity_id)

	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, fmt.Sprintf("Activity with ID %v Not Found", activity_id), nil)
		return
	}
	helpers.Response(w, 200, "Success", nil)

}
