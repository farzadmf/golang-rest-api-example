// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/contractor": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contractors"
                ],
                "summary": "List all contractors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Contractor"
                            }
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contractors"
                ],
                "summary": "Create a contractor",
                "parameters": [
                    {
                        "description": "Add contractor",
                        "name": "contractor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateContractor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Contractor"
                        }
                    },
                    "400": {
                        "description": "If either name or duration is empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/contractor/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contractors"
                ],
                "summary": "Get a contractor by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Contractor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Contractor"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contractors"
                ],
                "summary": "Delete a contractor by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Contractor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employee": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "List all employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Employee"
                            }
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Create an employee",
                "parameters": [
                    {
                        "description": "Add employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateEmployee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Employee"
                        }
                    },
                    "400": {
                        "description": "If either name or role is empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "If the specified role name cannot be found in the DB",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Get an employee by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Employee"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employees"
                ],
                "summary": "Delete an employee by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/member/{id}/tags": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Add tags to a member",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Member Tags",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MemberTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "When not all the specified tags exist in the DB",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reset": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "misc"
                ],
                "summary": "Reset the DB",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "List all roles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Role"
                            }
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create a role",
                "parameters": [
                    {
                        "description": "Add role",
                        "name": "contractor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Role"
                        }
                    },
                    "400": {
                        "description": "If name is empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Delete a role by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/{name}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Get a role by name",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "Role name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Role"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tag": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "List all tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Tag"
                            }
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Create a tag",
                "parameters": [
                    {
                        "description": "Add tag",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Tag"
                        }
                    },
                    "400": {
                        "description": "If value is empty",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tag/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Delete a tag by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tag/{value}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Get a tag by value",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "Tag value",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Tag"
                        }
                    },
                    "500": {
                        "description": "When there's an issue on the server side",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateContractor": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "string",
                    "example": "0"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.CreateEmployee": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.CreateRole": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "request.CreateTag": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "string"
                }
            }
        },
        "request.MemberTag": {
            "type": "object",
            "properties": {
                "values": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.Contractor": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.Employee": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.Tag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
