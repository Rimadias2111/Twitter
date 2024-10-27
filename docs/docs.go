// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/login": {
            "post": {
                "description": "API for user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retrieving all tweets with pagination and search",
                "tags": [
                    "tweet"
                ],
                "summary": "Get all tweets",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of tweets per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search term",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User ID for filtering tweets",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllTweetsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating a new tweet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tweet"
                ],
                "summary": "Create a tweet",
                "parameters": [
                    {
                        "description": "Tweet data",
                        "name": "tweet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tweet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseId"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets/feed": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retrieving tweets from users that the current user is following",
                "tags": [
                    "tweet"
                ],
                "summary": "Get tweets from followed users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of tweets per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllTweetsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets/{tweer_id}/unlike": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for unliking a tweet",
                "tags": [
                    "tweet"
                ],
                "summary": "Unlike a tweet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets/{tweet_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retrieving a tweet by ID",
                "tags": [
                    "tweet"
                ],
                "summary": "Get a tweet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID",
                        "name": "tweet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tweet"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for updating a tweet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tweet"
                ],
                "summary": "Update a tweet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID",
                        "name": "tweet_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tweet data",
                        "name": "tweet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Tweet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for deleting a tweet",
                "tags": [
                    "tweet"
                ],
                "summary": "Delete a tweet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets/{tweet_id}/like": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for liking a tweet",
                "tags": [
                    "tweet"
                ],
                "summary": "Like a tweet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/tweets/{tweet_id}/retweet": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retweeting an existing tweet",
                "tags": [
                    "tweet"
                ],
                "summary": "Retweets a tweet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tweet ID to retweet",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseId"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retrieving all users with pagination and search",
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of users per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search term",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseId"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/users/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for retrieving a user by ID",
                "tags": [
                    "user"
                ],
                "summary": "Get a user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for updating a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for deleting a user",
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/users/{user_id}/follow": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for following a user",
                "tags": [
                    "user"
                ],
                "summary": "Follow a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to follow",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/users/{user_id}/unfollow": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for unfollowing a user",
                "tags": [
                    "user"
                ],
                "summary": "Unfollow a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to unfollow",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.GetAllTweetsResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "tweets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tweet"
                    }
                }
            }
        },
        "models.GetAllUsersResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseId": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.ResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Tweet": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "string"
                },
                "imagePath": {
                    "type": "string"
                },
                "retweetID": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                },
                "videoPath": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profileImage": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
