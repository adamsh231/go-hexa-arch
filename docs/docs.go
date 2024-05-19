// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Adam Syarif Hidayatullah",
            "email": "adamsyarif219@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/log": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Logging"
                ],
                "summary": "Search log activities",
                "parameters": [
                    {
                        "type": "string",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "service",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseUtil"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.SearchActivityOutput"
                                            }
                                        },
                                        "meta": {
                                            "$ref": "#/definitions/utils.ResponseMetaUtil"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseStatusUtil"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseStatusUtil"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseStatusUtil"
                        }
                    }
                }
            }
        },
        "/v1/log/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Logging"
                ],
                "summary": "Search log activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseUtil"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.FindActivityOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseStatusUtil"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseStatusUtil"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.FindActivityOutput": {
            "type": "object",
            "properties": {
                "activity": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "response": {
                    "type": "object",
                    "additionalProperties": true
                },
                "service": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "entities.SearchActivityOutput": {
            "type": "object",
            "properties": {
                "activity": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseMetaUtil": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "from": {
                    "type": "integer"
                },
                "last_page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "to": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "utils.ResponseStatusUtil": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseUtil": {
            "type": "object",
            "properties": {
                "data": {},
                "meta": {},
                "status": {
                    "$ref": "#/definitions/utils.ResponseStatusUtil"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "/go-hexa",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Golang Hexa Swagger",
	Description:      "Golang hexagonal swagger documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
