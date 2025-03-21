package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:ContraseñaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:IdentificaionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:Tipo_documentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
