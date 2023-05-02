// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/blog-post": {
            "get": {
                "description": "Get all exists blogPosts.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BlogPosts"
                ],
                "summary": "get all exists blogPosts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BlogPost"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new blogPost.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BlogPosts"
                ],
                "summary": "create a new blogPost",
                "parameters": [
                    {
                        "description": "BlogPost",
                        "name": "blogPost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBlogPostSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            }
        },
        "/api/blog-post/{id}": {
            "get": {
                "description": "Get blogPosts by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BlogPosts"
                ],
                "summary": "get blogPosts by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "blogPostId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete blogPost by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BlogPosts"
                ],
                "summary": "delete blogPost by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "blogPostId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update blogPost.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BlogPosts"
                ],
                "summary": "update blogPost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "blogPostId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "BlogPost",
                        "name": "blogPost",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBlogPostSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BlogPost"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BlogPost": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.CreateBlogPostSchema": {
            "type": "object",
            "required": [
                "body",
                "description",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "20.25.40.163:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Blog-Post API",
	Description:      "This is a sample server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
