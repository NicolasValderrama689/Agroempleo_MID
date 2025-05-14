package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// PerfilController operations for Perfil
type PerfilController struct {
	beego.Controller
}

// URLMapping ...
func (c *PerfilController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Perfil
// @Param	body		body 	models.Perfil	true		"body for Perfil content"
// @Success 201 {object} models.Perfil
// @Failure 403 body is empty
// @router / [post]
func (c *PerfilController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Perfil by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Perfil
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PerfilController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Perfil
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Perfil
// @Failure 403
// @router / [get]
func (c *PerfilController) GetAll() {

	json_laboral, _ := services.Metodo_get_all("host_api3", "laboral/")
	json_laboral1, _ := services.ProcessarJson(json_laboral)
	laboral_json := json_laboral1["usuarios consultados"]
	arreglo_laboral, _ := services.ConvertInterfaceToSliceMap(laboral_json)
	fmt.Println(arreglo_laboral)
	

	var arre []map[string]interface{}
	for i := range arreglo_laboral {
		id_usuario, _ := arreglo_laboral[i]["IdUsuarios"]
		id_usuario_string := fmt.Sprintf("%v", id_usuario)

		enpoint := "Usuarios/" + id_usuario_string
		final,_ := services.Metodo_getid("host_api2", enpoint)
		Json_usuario, _ := services.ProcessarJson(final)
		nombre_usuario := Json_usuario["Consulta de id"].(map[string]interface{})["Nombre"]

		fechaInicio := arreglo_laboral[i]["FechaInicio"]
		fechaFin := arreglo_laboral[i]["FechaFin"]
		perido := fmt.Sprintf("%v a %v", fechaInicio, fechaFin)

		arreglo_parcial := map[string]interface{}{
			"Nombre": nombre_usuario,
			"Cargo": arreglo_laboral[i]["Cargo"],
			"Empresa": arreglo_laboral[i]["Empresa"],
			"Perido": perido,
			"Telefono": Json_usuario["Consulta de id"].(map[string]interface{})["Telefono"],
			"Email": Json_usuario["Consulta de id"].(map[string]interface{})["CorreoElectronico"],
		}
		arre = append(arre, arreglo_parcial)
		fmt.Println(arre)


	}
	
	c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Perfil existosa", "Data": arre}
	c.ServeJSON()

}


// Put ...
// @Title Put
// @Description update the Perfil
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Perfil	true		"body for Perfil content"
// @Success 200 {object} models.Perfil
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PerfilController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Perfil
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PerfilController) Delete() {

}
