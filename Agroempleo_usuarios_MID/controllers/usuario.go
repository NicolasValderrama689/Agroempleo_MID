package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/models"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/services"
)

// UsuarioController operations for Usuario
type UsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Usuarios
// @Param	body		body 	models.Usuarios	true		"body for Usuarios content"
// @Success 201 {object} models.Usuarios
// @Failure 400 Bad Request
// @Failure 500 Internal Server Error
// @router / [post]
func (c *UsuarioController) Post() {

	var usuario models.Usuarios

	// Decodificar el JSON recibido
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &usuario); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error en el formato de entrada: " + err.Error(),
		}
		c.ServeJSON()
		return
	}

	// Validar que los campos esenciales no estén vacíos
	if usuario.Nombre == "" || usuario.CorreoElectronico == "" || usuario.IdIdentificacionIdentificacion == nil || usuario.IdRolRol == nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Faltan campos obligatorios: Nombre, Correo, Cedula y Rol.",
		}
		c.ServeJSON()
		return
	}

	// Convertir a JSON para enviar al CRUD local
	jsonData, err := json.Marshal(usuario)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  500,
			"message": "Error al codificar usuario: " + err.Error(),
		}
		c.ServeJSON()
		return
	}

	// Hacer la solicitud HTTP POST al CRUD local
	resp, err := http.Post("http://localhost:8080/v1/Usuarios", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  500,
			"message": "Error al comunicarse con el CRUD local: " + err.Error(),
		}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

}

// GetOne ...
// @Title GetOne
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Usuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuario
// @Failure 403
// @router / [get]
func (c *UsuarioController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Usuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuarioController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Usuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuarioController) Delete() {

	idUsuario, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	// Llamamos al servicio de autoeliminación
	err = services.AutoEliminarUsuario(idUsuario)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Error al eliminar la cuenta",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("La cuenta con ID %d ha sido eliminada correctamente", idUsuario),
		}
	}

	c.ServeJSON()
}
