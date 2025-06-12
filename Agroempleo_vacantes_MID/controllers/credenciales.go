package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// CredencialesController operations for Credenciales
type CredencialesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CredencialesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Credenciales
// @Param	body		body 	models.Credenciales	true		"body for Credenciales content"
// @Success 201 {object} models.Credenciales
// @Failure 403 body is empty
// @router / [post]
func (c *CredencialesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Credenciales by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Credenciales
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CredencialesController) GetOne() {


}

// GetAll ...
// @Title GetAll
// @Description get Credenciales
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Credenciales
// @Failure 403
// @router / [get]
func (c *CredencialesController) GetAll() {

	json_credeciales, _ := services.Metodo_get_all("host_api2", "Usuarios/")
	json_cred, _ := services.ProcessarJson(json_credeciales)
	credenciales_json := json_cred["Consulta de id"]
	arreglo_cred, _ := services.ConvertInterfaceToSliceMap(credenciales_json)
	fmt.Println("hola", arreglo_cred)

	var credenciales []map[string]interface{}
	for i := range arreglo_cred {

		credecial_parcial := map[string]interface{}{
			"Id": arreglo_cred[i]["Id"],
			"username": arreglo_cred[i]["Apellido"],
			"contraseña": arreglo_cred[i]["IdContraseñasContraseñas"].(map[string]interface{})["Contraseña"],
			"correo": arreglo_cred[i]["CorreoElectronico"],
			"nombre": arreglo_cred[i]["Nombre"],
			
		}
		credenciales = append(credenciales, credecial_parcial)
	}
	c.Data["json"] = map[string]interface{}{"Succes": true, "user": credenciales}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Credenciales
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Credenciales	true		"body for Credenciales content"
// @Success 200 {object} models.Credenciales
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CredencialesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Credenciales
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CredencialesController) Delete() {

}
