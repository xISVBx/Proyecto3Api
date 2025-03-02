// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://jecopainsurance.com/terms/",
        "contact": {
            "name": "Soporte Col Moda",
            "url": "http://jecopainsurance.com/support",
            "email": "soporte@jecopainsurance.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/categories": {
            "get": {
                "description": "Retorna un listado de las categorias de la aplicacion.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Obtener todas las categorias",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.CategoryResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Permite a los usuarios autenticarse en la plataforma.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Iniciar sesión",
                "parameters": [
                    {
                        "description": "Credenciales de usuario",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/products": {
            "get": {
                "description": "Retorna un listado de los productos aplicando filtros.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Obtener todos los productos por filtros",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Texto de búsqueda",
                        "name": "SearchQuery",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.ProductResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "Crea una nueva cuenta para el usuario con los datos proporcionados.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Registro de usuario",
                "parameters": [
                    {
                        "description": "Datos del usuario",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.RegisterResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/roles": {
            "get": {
                "description": "Retorna un listado de los roles de la aplicacion.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "summary": "Obtener todos los roles",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Role"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {}
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/version": {
            "get": {
                "description": "Retorna la versión de la API configurada en el sistema.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Version"
                ],
                "summary": "Obtener versión",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.AppResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CategoryResponse": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.LoginDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dtos.ProductResponse": {
            "type": "object",
            "properties": {
                "categoryID": {
                    "type": "string"
                },
                "companyID": {
                    "type": "string"
                },
                "discountID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "longDescription": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "shortDescription": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "subCategoryID": {
                    "type": "string"
                },
                "tagID": {
                    "type": "string"
                }
            }
        },
        "dtos.RegisterRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "cityId": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dtos.RegisterResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Role": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.AppResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3031",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Col Moda API",
	Description:      "API para la tienda de moda Col Moda.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
