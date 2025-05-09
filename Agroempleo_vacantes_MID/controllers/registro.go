package controllers

import (
	

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// RegistroController operations for Registro
type RegistroController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Registro
// @Param	body		body 	models.Registro	true		"body for Registro content"
// @Success 201 {object} models.Registro
// @Failure 403 body is empty
// @router / [post]
func (c *RegistroController) Post() {
	fmt.Println("Funcion Post")

	var body_ingresa map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("json ingresa", body_ingresa)
	}

	body_contrasena := map[string]interface{}{
		"Contraseña": body_ingresa["contrasena"],
	}

	bytes_contrasena, err := json.Marshal(body_contrasena)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_contrasena_byte, _ := services.Metodo_post("host_api2", "contraseñas", bytes_contrasena)

	var response_json_contrasena map[string]interface{}

	err1 := json.Unmarshal(body_response_contrasena_byte, &response_json_contrasena)
	if err1 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}
	diastring, _ := json.Marshal(body_ingresa["dia"])
	mestring, _ := json.Marshal(body_ingresa["mes"])
	aniotring, _ := json.Marshal(body_ingresa["anio"])


	var fecha_nacimiento = string(diastring) + "/" + string(mestring) + "/" + string(aniotring)

	fmt.Println("fecha", fecha_nacimiento)

	diastring1, _ := json.Marshal(body_ingresa["pais"])
	mestring1, _ := json.Marshal(body_ingresa["departamento"])
	aniotring1, _ := json.Marshal(body_ingresa["ciudad"])

	var ubicacion = string(diastring1) + "/" + string(mestring1) + "/" + string(aniotring1)
	fmt.Println("ubicacion", ubicacion)

	id_contrasena := response_json_contrasena["Datos creados con id"].(map[string]interface{})["Id"]
	id_tipo_documento := body_ingresa["tipoDocumento"]
	id_tipo_documento_string := fmt.Sprintf("%v", id_tipo_documento)
	id_tipo_documento_int, _ := strconv.Atoi(id_tipo_documento_string)
	id_tipo_usuario := body_ingresa["tipoUsuario"]
	id_tipo_usuario_string := fmt.Sprintf("%v", id_tipo_usuario)
	id_tipo_usuario_int, _ := strconv.Atoi(id_tipo_usuario_string)

	fmt.Println("id contraseña", id_contrasena)

	body_Usuario := map[string]interface{}{
		"Nombre":            body_ingresa["nombre"],
		"Apellido":          body_ingresa["apellido"],
		"FNacimiento":       fecha_nacimiento,
		"NDocumento":        body_ingresa["numeroDocumento"],
		"CorreoElectronico": body_ingresa["correo_electronico"],
		"Contraseña":        map[string]interface{}{"id": id_contrasena},
		"IdTipoDocumento":   map[string]interface{}{"id": id_tipo_documento_int},
		"IdTipoUsuario":     map[string]interface{}{"id": id_tipo_usuario_int},
		"Ubicacion":         ubicacion,
	}
	bytes_usuario, err := json.Marshal(body_Usuario)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_usuario_byte, _ := services.Metodo_post("host_api2", "Usuarios", bytes_usuario)

	var response_json_usuario map[string]interface{}

	err2 := json.Unmarshal(body_response_usuario_byte, &response_json_usuario)
	if err2 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Creación existosa",
		"Data":    body_ingresa,
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Registro by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Registro
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RegistroController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Registro
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Registro
// @Failure 403
// @router / [get]
func (c *RegistroController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Registro
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Registro	true		"body for Registro content"
// @Success 200 {object} models.Registro
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RegistroController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Registro
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RegistroController) Delete() {

}
