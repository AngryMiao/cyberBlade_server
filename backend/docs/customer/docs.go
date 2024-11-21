// Package customer GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package customer

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
        "/chat": {
            "post": {
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
                    "AI"
                ],
                "summary": "AI 语音聊天接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/chatSerializer.VoiceChatBodyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/chatSerializer.ChatReply"
                        }
                    }
                }
            }
        },
        "/chat/earphone-config": {
            "post": {
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
                    "AI"
                ],
                "summary": "AI 耳机配置接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/chatSerializer.GlassesConfigRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.ChatConfig"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "chatSerializer.ChatBodyRequest": {
            "type": "object",
            "required": [
                "device_code",
                "message"
            ],
            "properties": {
                "device_code": {
                    "type": "string"
                },
                "image": {
                    "description": "Image      *multipart.FileHeader ` + "`" + `form:\"image\" binding:\"omitempty\"` + "`" + `",
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "chatSerializer.ChatReply": {
            "type": "object",
            "properties": {
                "func_name": {
                    "type": "string"
                },
                "reply_msg": {
                    "type": "string"
                }
            }
        },
        "chatSerializer.GlassesConfigRequest": {
            "type": "object",
            "required": [
                "device_code",
                "forward_mode"
            ],
            "properties": {
                "device_code": {
                    "type": "string"
                },
                "discord_user_id": {
                    "type": "string"
                },
                "forward_mode": {
                    "type": "string",
                    "enum": [
                        "all",
                        "media"
                    ]
                }
            }
        },
        "chatSerializer.VoiceChatBodyRequest": {
            "type": "object",
            "required": [
                "deviceCode"
            ],
            "properties": {
                "deviceCode": {
                    "type": "string"
                }
            }
        },
        "ent.ChatConfig": {
            "type": "object",
            "properties": {
                "create_time": {
                    "description": "CreateTime holds the value of the \"create_time\" field.",
                    "type": "string"
                },
                "device_code": {
                    "description": "DeviceCode holds the value of the \"device_code\" field.",
                    "type": "string"
                },
                "discord_user_id": {
                    "description": "DiscordUserID holds the value of the \"discord_user_id\" field.",
                    "type": "string"
                },
                "forward_mode": {
                    "description": "ForwardMode holds the value of the \"forward_mode\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "update_time": {
                    "description": "UpdateTime holds the value of the \"update_time\" field.",
                    "type": "string"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
