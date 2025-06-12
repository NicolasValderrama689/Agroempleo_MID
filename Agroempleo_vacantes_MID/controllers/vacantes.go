package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// VacantesController operations for Vacantes
type VacantesController struct {
	beego.Controller
}

// URLMapping ...
func (c *VacantesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Vacantes
// @Param	body		body 	models.Vacantes	true		"body for Vacantes content"
// @Success 201 {object} models.Vacantes
// @Failure 403 body is empty
// @router / [post]
func (c *VacantesController) Post() {
	fmt.Println("Metodo post")

}

// GetOne ...
// @Title GetOne
// @Description get Vacantes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vacantes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VacantesController) GetOne() {
	fmt.Println("metodo get by id")

}

// GetAll ...
// @Title GetAll
// @Description get Vacantes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Vacantes
// @Failure 403
// @router / [get]
func (c *VacantesController) GetAll() {
	fmt.Println("metodo Get all")
	Json_vacantes, _ := services.Metodo_get_all("host_api", "Vacantes")

	//fmt.Println("Este es el valor de Json_vacante2 en byte", Json_vacantes2)
	Json_procesado_vacantes, _ := services.ProcessarJson(Json_vacantes)


	Vacantes_json := Json_procesado_vacantes["Consulta de id"]

	Arreglo_Map_Vacantes_final, _ := services.ConvertInterfaceToSliceMap(Vacantes_json)


	
		
	var Resultado_total []map[string]interface{}

	for i := range Arreglo_Map_Vacantes_final {

		usuarioID := Arreglo_Map_Vacantes_final[i]["Id_usuarios"]
		usuarioID_string := fmt.Sprintf("%v", usuarioID)

		endpoint := "Usuarios/"+ usuarioID_string
	

		final,_ := services.Metodo_getid("host_api2", endpoint)
		Json_usuario, _ := services.ProcessarJson(final)
		nombre_usuario := Json_usuario["Consulta de id"].(map[string]interface{})["Nombre"]

		
		

		Resultado_parcial := map[string]interface{}{
			"id":                  Arreglo_Map_Vacantes_final[i]["Id"],
			"TituloPuesto":        Arreglo_Map_Vacantes_final[i]["TituloPuesto"],
			"DescripcionTrabajo":  Arreglo_Map_Vacantes_final[i]["DescripcionTrabajo"],
			"Cargo":               Arreglo_Map_Vacantes_final[i]["Cargo"],
			"Salario":             Arreglo_Map_Vacantes_final[i]["Salario"],
			"Horario":             Arreglo_Map_Vacantes_final[i]["Horario"],
			"Modalidad":           Arreglo_Map_Vacantes_final[i]["Modalidad"],
			"NivelRequerido":      Arreglo_Map_Vacantes_final[i]["NivelRequerido"],
			"ExperienciaRequrida": Arreglo_Map_Vacantes_final[i]["ExperienciaRequrida"],
			"NumeroVacantes":      Arreglo_Map_Vacantes_final[i]["NumeroVacantes"],
			"Activo":              Arreglo_Map_Vacantes_final[i]["Activo"],
			"TipoEmpleo":          Arreglo_Map_Vacantes_final[i]["IdtipoempleoTipodeempleo"].(map[string]interface{})["Nombre"],
			"Ciudad":              Arreglo_Map_Vacantes_final[i]["Idciudadtrabajociudad"].(map[string]interface{})["Nombre"],
			"publicado_por": nombre_usuario,
		}
		Resultado_total = append(Resultado_total, Resultado_parcial)

	}

	c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Consulta existosa", "Data": Resultado_total}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Vacantes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Vacantes	true		"body for Vacantes content"
// @Success 200 {object} models.Vacantes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VacantesController) Put() {
	fmt.Println("metodo put")

}

// Delete ...
// @Title Delete
// @Description delete the Vacantes
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *VacantesController) Delete() {
	fmt.Println("metodo delete")
}
