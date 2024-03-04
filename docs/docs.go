// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/email/send": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "获取验证码",
                "parameters": [
                    {
                        "description": "邮箱",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Email"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "返回用户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.User"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/login/account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户根据账号登录",
                "parameters": [
                    {
                        "description": "账号,密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserLoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/login/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户根据邮箱登录",
                "parameters": [
                    {
                        "description": "邮箱,验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserLoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/register/account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户根据账号注册",
                "parameters": [
                    {
                        "description": "账号,密码,确认密码,用户名",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegisterAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回注册信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/register/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户根据邮箱注册",
                "parameters": [
                    {
                        "description": "邮箱,验证码,用户名",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegisterEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回注册信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
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
        "model.User": {
            "type": "object",
            "properties": {
                "accessKey": {
                    "type": "string"
                },
                "account": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "balance": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "invitationCode": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "secretKey": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.Email": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "request.UserLoginAccount": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UserLoginEmail": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "request.UserRegisterAccount": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "checkPassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.UserRegisterEmail": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.UserLoginResponse": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "a-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ax-API 接口文档",
	Description:      "api接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
