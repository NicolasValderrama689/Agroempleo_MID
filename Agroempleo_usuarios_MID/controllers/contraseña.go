package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"your_project/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/mux"
)

// ContraseñaController operations for Contraseña
type ContraseñaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContraseñaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Contraseña
// @Param	body		body 	models.Contraseña	true		"body for Contraseña content"
// @Success 201 {object} models.Contraseña
// @Failure 403 body is empty
// @router / [post]
// Post crea una nueva contraseña en la base de datos
func (c *ContraseñaController) Post(w http.ResponseWriter, r *http.Request) {
	var contraseña models.Contraseña
	// Decodificar el JSON recibido en la variable contraseña
	if err := json.NewDecoder(r.Body).Decode(&contraseña); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	// Guardar la contraseña en la base de datos
	if err := models.CreateContraseña(&contraseña); err != nil {
		http.Error(w, "Error al guardar la contraseña", http.StatusInternalServerError)
		return
	}
	// Responder con éxito
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contraseña)
}

// GetOne ...
// @Title GetOne
// @Description get Contraseña by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Contraseña
// @Failure 403 :id is empty
// @router /:id [get]
// GetOne obtiene una contraseña específica por su ID
func (c *ContraseñaController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	contraseña, err := models.GetContraseñaByID(id)
	if err != nil {
		http.Error(w, "Contraseña no encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(contraseña)
}

// GetAll ...
// @Title GetAll
// @Description get Contraseña
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contraseña
// @Failure 403
// @router / [get]
// GetAll obtiene todas las contraseñas almacenadas en la base de datos
func (c *ContraseñaController) GetAll(w http.ResponseWriter, r *http.Request) {
	contraseñas, err := models.GetAllContraseñas()
	if err != nil {
		http.Error(w, "Error al obtener las contraseñas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contraseñas)
}

// Put ...
// @Title Put
// @Description update the Contraseña
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Contraseña	true		"body for Contraseña content"
// @Success 200 {object} models.Contraseña
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ContraseñaController) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var contraseña models.Contraseña
	if err := json.NewDecoder(r.Body).Decode(&contraseña); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	contraseña.ID = id
	if err := models.UpdateContraseña(&contraseña); err != nil {
		http.Error(w, "Error al actualizar la contraseña", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contraseña)
}

// Delete ...
// @Title Delete
// @Description delete the Contraseña
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ContraseñaController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := models.DeleteContraseña(id); err != nil {
		http.Error(w, "Error al eliminar la contraseña", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
