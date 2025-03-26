package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
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
// @Description create Usuario
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {object} models.Usuario
// @Failure 403 body is empty
// @router / [post]
func (c *UsuarioController) Post() {
<<<<<<< HEAD
	var body_ingreso map[string]interface{}
	var reponseUsuario, responseCredencial, responseRolUsuario []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingreso); err == nil {
		fmt.Println("Body que ingresa", body_ingreso)

		jsonData, err := json.MarshalIndent(body_ingreso, "", " ")
		if err != nil {
			fmt.Println("Error al convertir a JSON", err)
		}
		fmt.Println("Body de ingreso en JSON:", string(jsonData))

		passStr, _ := body_ingreso["contraseña"].(string)

		// Hashear la contraseña
		hashedPass, err := services.HashContraseña(passStr)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = map[string]interface{}{"error": "Error interno al procesar la contraseña"}
			c.ServeJSON()
			return
		}

		jsonCredencial := map[string]interface{}{
			"contraseña": hashedPass,
		}
		fmt.Println("este es el json para credenciales: ", jsonCredencial)

		jsonUsuario := map[string]interface{}{
			"nombre":             body_ingreso["Nombre"],
			"apellido":           body_ingreso["Apellido"],
			"contacto":           body_ingreso["Contacto"],
			"correo_electronico": body_ingreso["CorreoElectronico"],
		}
		fmt.Println("este es el json usuario: ", jsonUsuario)

		json_credencial_byte, _ := json.Marshal(jsonCredencial)
		// json_usuario_byte, _ := json.Marshal(jsonUsuario)

		fmt.Println("json credencial: ", string(json_credencial_byte))
		responseCredencial, _ = services.Metodo_post("API_CRUD", "/v1/Credenciales", json_credencial_byte)
		if err != nil {
			fmt.Println("Error al crear credenciales:", err)
			return
		}
		fmt.Println("Respuesta de la API (Credenciales): ", string(responseCredencial))

		// Obtener el ID de credencial creada
		var credencialresponse map[string]interface{}
		if err := json.Unmarshal(responseCredencial, &credencialresponse); err != nil {
			fmt.Println("Error al parsear respuesta de credenciales:", err)
			return
		}

		// Extraer el ID de la credencial
		var credencialID float64
		if data, ok := credencialresponse["Data"].(map[string]interface{}); ok {
			if id, exists := data["Id"].(float64); exists {
				credencialID = id
			} else {
				fmt.Println("Error: No se encontró el ID en la respuesta de Credenciales")
				return
			}
		} else {
			fmt.Println("Error: Estructura de respuesta de credenciales no válida")
			return
		}
		fmt.Println("Id de credenciales: ", credencialID)

		// Crear JSON para Usuario con fk_credencial
		jsonUsuario = map[string]interface{}{
			"Nombre":            body_ingreso["Nombre"],
			"Apellido":          body_ingreso["Apellido"],
			"Contacto":          body_ingreso["Contacto"],
			"CorreoElectronico": body_ingreso["CorreoElectronico"],
			"FkCredencial": map[string]interface{}{
				"Id": credencialID,
			},
		}

		// Convertir a JSON y enviar POST a /v1/Usuario
		json_usuario_byte, _ := json.Marshal(jsonUsuario)
		fmt.Println("Enviando JSON a /v1/Usuario:", string(json_usuario_byte))
		reponseUsuario, err = services.Metodo_post("API_CRUD", "/v1/Usuario", json_usuario_byte)
		if err != nil {
			fmt.Println(" Error al crear usuario:", err)
			return
		}

		fmt.Println("Respuesta de la API (Usuario):", string(reponseUsuario))

		// Extraer el ID del usuario creado
		var usuarioResponse map[string]interface{}
		if err := json.Unmarshal(reponseUsuario, &usuarioResponse); err != nil {
			fmt.Println("Error al parsear respuesta de usuario:", err)
			return
		}

		var usuarioID float64
		if data, ok := usuarioResponse["Data"].(map[string]interface{}); ok {
			if id, exists := data["Id"].(float64); exists {
				usuarioID = id
			} else {
				fmt.Println("Error: No se encontró el ID en la respuesta de Usuario")
				return
			}
		} else {
			fmt.Println("Error: Estructura de respuesta de usuario no válida")
			return
		}

		// Buscar el ID del Rol en la tabla Roles
		rolNombre := body_ingreso["rol"].(string) // Extrae el rol enviado en la solicitud
		var responseRol []byte
		responseRol, err = services.Metodo_get("API_CRUD", "/v1/Roles?query=nombre:", rolNombre)

		if err != nil {
			fmt.Println("Error al obtener el rol:", err)
			return
		}

		var rolResponse map[string]interface{}
		if err := json.Unmarshal(responseRol, &rolResponse); err != nil {
			fmt.Println("Error al parsear respuesta de roles:", err)
			return
		}

		var rolID float64
		if roles, ok := rolResponse["Data"].([]interface{}); ok && len(roles) > 0 {
			if rolData, exists := roles[0].(map[string]interface{}); exists {
				rolID = rolData["Id"].(float64)
			}
		} else {
			fmt.Println("Error: No se encontró el rol en la base de datos")
			return
		}

		// Crear registro en RolesUsuario
		jsonRolUsuario := map[string]interface{}{
			"FkUsuarioRoles": map[string]interface{}{
				"Id": usuarioID,
			},
			"FkRolesUsuario": map[string]interface{}{
				"Id": rolID,
			},
		}

		json_rol_usuario_byte, _ := json.Marshal(jsonRolUsuario)
		responseRolUsuario, _ = services.Metodo_post("API_CRUD", "/v1/Roles_Usuario", json_rol_usuario_byte)

		fmt.Println("Respuesta de la API (RolesUsuario):", string(responseRolUsuario))
	}

	c.Data["json"] = map[string]interface{}{
		"Message": "¡Usuario creado exitosamente!",
	}
	c.ServeJSON()
=======
	app.post('/api/usuarios', (req, res) => {
		const { Nombre, Apellido, FechaNacimiento, CorreoElectronico, Ciudad, Departamento, Pais, Telefono, IdRolRol, IdIdentificacionIdentificacion, IdContraseñasContraseñas } = req.body;
	
		// Validaciones básicas
		if (!Nombre || !Apellido || !FechaNacimiento || !CorreoElectronico || !Telefono) {
			return res.status(400).json({ message: 'Faltan datos requeridos' });
		}
	
		// Generando un nuevo Id (en un escenario real, esto debería ser generado por la base de datos)
		const nuevoId = usuarios.length + 1;
		const nuevoUsuario = {
			Id: nuevoId,
			Nombre,
			Apellido,
			FechaNacimiento,
			CorreoElectronico,
			Ciudad,
			Departamento,
			Pais,
			Telefono,
			FechaModificacion: new Date().toISOString(),
			FechaCreacion: new Date().toISOString(),
			Activo: true,  // Se puede poner como "true" por defecto
			IdRolRol,
			IdIdentificacionIdentificacion,
			IdContraseñasContraseñas
		};
	
		// Guardando el nuevo usuario (en este caso lo añadimos a la base de datos en memoria)
		usuarios.push(nuevoUsuario);
	
		// Devolviendo el usuario creado con un estado 201
		res.status(201).json(nuevoUsuario);
	});
>>>>>>> 854503d385002be75c809a9599f40e7077e42faf

}

// GetOne ...
// @Title GetOne
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {
	fmt.println("funcion getone")
	fmt.println("funcion getone")

	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	
	// Simulando una base de datos de usuarios (en memoria)
	var users = []User{
		{ID: 1, Username: "johndoe", Email: "johndoe@example.com"},
		{ID: 2, Username: "janedoe", Email: "janedoe@example.com"},
	}
	
	func main() {
		// Crear un router de Gin
		r := gin.Default()
	
		// Ruta para obtener un usuario por ID
		r.GET("/users/:id", getUser)
	
		// Correr la API en el puerto 8080
		r.Run(":8080")
	}
	
	// Controlador para obtener un usuario por ID
	func getUser(c *gin.Context) {
		// Obtener el ID del parámetro de la URL
		idParam := c.Param("id")
	
		// Convertir el ID de string a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser un número válido"})
			return
		}
	
		// Buscar el usuario con el ID especificado
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
	
		// Si no se encuentra el usuario
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	}

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

}
