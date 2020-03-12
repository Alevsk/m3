// Code generated by go-swagger; DO NOT EDIT.

// // This file is part of MinIO Kubernetes Cloud
// // Copyright (c) 2019 MinIO, Inc.
// //
// // This program is free software: you can redistribute it and/or modify
// // it under the terms of the GNU Affero General Public License as published by
// // the Free Software Foundation, either version 3 of the License, or
// // (at your option) any later version.
// //
// // This program is distributed in the hope that it will be useful,
// // but WITHOUT ANY WARRANTY; without even the implied warranty of
// // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// // GNU Affero General Public License for more details.
// //
// // You should have received a copy of the GNU Affero General Public License
// // along with this program.  If not, see <http://www.gnu.org/licenses/>.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "MinIO Console Server",
    "version": "0.1.0"
  },
  "paths": {
    "/api/v1/buckets": {
      "get": {
        "tags": [
          "UserAPI"
        ],
        "summary": "List Buckets",
        "operationId": "ListBuckets",
        "parameters": [
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listBucketsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "UserAPI"
        ],
        "summary": "Make a bucket",
        "operationId": "MakeBucket",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/makeBucketRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "bucket": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "access": {
          "$ref": "#/definitions/bucketAccess"
        },
        "creation_date": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "minLength": 3
        },
        "size": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "bucketAccess": {
      "type": "string",
      "default": "PRIVATE",
      "enum": [
        "PRIVATE",
        "PUBLIC",
        "CUSTOM"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "listBucketsResponse": {
      "type": "object",
      "properties": {
        "buckets": {
          "type": "array",
          "title": "list of resulting buckets",
          "items": {
            "$ref": "#/definitions/bucket"
          }
        },
        "total_buckets": {
          "type": "integer",
          "format": "int64",
          "title": "number of buckets accessible to tenant user"
        }
      }
    },
    "makeBucketRequest": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "access": {
          "$ref": "#/definitions/bucketAccess"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "MinIO Console Server",
    "version": "0.1.0"
  },
  "paths": {
    "/api/v1/buckets": {
      "get": {
        "tags": [
          "UserAPI"
        ],
        "summary": "List Buckets",
        "operationId": "ListBuckets",
        "parameters": [
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listBucketsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "UserAPI"
        ],
        "summary": "Make a bucket",
        "operationId": "MakeBucket",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/makeBucketRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "bucket": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "access": {
          "$ref": "#/definitions/bucketAccess"
        },
        "creation_date": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "minLength": 3
        },
        "size": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "bucketAccess": {
      "type": "string",
      "default": "PRIVATE",
      "enum": [
        "PRIVATE",
        "PUBLIC",
        "CUSTOM"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "listBucketsResponse": {
      "type": "object",
      "properties": {
        "buckets": {
          "type": "array",
          "title": "list of resulting buckets",
          "items": {
            "$ref": "#/definitions/bucket"
          }
        },
        "total_buckets": {
          "type": "integer",
          "format": "int64",
          "title": "number of buckets accessible to tenant user"
        }
      }
    },
    "makeBucketRequest": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "access": {
          "$ref": "#/definitions/bucketAccess"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
