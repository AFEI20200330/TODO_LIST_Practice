package handler

import (
	"TODO_LIST_Practice/app/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


func GetAllProjects(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db,title,w,r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}
func getProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project{
	project := model.Project{}
	if err := db.First(&project,model.Project{Title:title}).Error; err != nil {
		respondErr(w, http.StatusNotFound, "Project not found")
		return nil
	}
	return &project
}

func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOr404(db,title,w,r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project);err != nil {
		respondErr(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db,title,w,r)
	if project == nil {
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func ArchiveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOr404(db,title,w,r)
	if project == nil {
		return 
	}
	project.Archive()
	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func RestoreProject(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	title := vars["title"]
	project := getProjectOr404(db,title,w,r)
	if project == nil {
		return
	}
	project.Restore()
	if err := db.Save(&project).Error; err != nil {
		respondErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}




