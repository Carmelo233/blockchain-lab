// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "叶浩辉",
            "url": "blog.yehaohui.com",
            "email": "yhh1934292134@163.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/supplychain/:id": {
            "get": {
                "description": "根据id查询一个信息",
                "produces": [
                    "application/json"
                ],
                "summary": "查询一个信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "根据上传json创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建一个信息",
                "parameters": [
                    {
                        "type": "string",
                        "name": "describe",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "distributor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "expiry_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "manufacturing_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "producer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "product_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "retailer",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "根据id删除一个信息",
                "produces": [
                    "application/json"
                ],
                "summary": "删除一个信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要删除信息的ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/supplychain/create": {
            "post": {
                "description": "根据上传json创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建一个信息",
                "parameters": [
                    {
                        "type": "string",
                        "name": "describe",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "distributor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "expiry_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "manufacturing_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "producer",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "product_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "retailer",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/supplychains": {
            "get": {
                "description": "查询所有信息",
                "produces": [
                    "application/json"
                ],
                "summary": "查询所有信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "119.29.53.176:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fabric",
	Description:      "fabric-crud接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}