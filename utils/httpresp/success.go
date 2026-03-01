package httpresp

import (
	"github.com/gin-gonic/gin"
)

// NewResult sets the result body to context
func NewResult(ctx *gin.Context, status int, response interface{}) {
	success := true

	resp := JSONResult{
		Success: success,
		Code:    status,
		Result:  response,
	}
	ctx.JSON(status, resp)
}

// JSONResult returns a json object
type JSONResult struct {
	Success bool        `json:"success" example:"true"` // wether the business action is successfull or not
	Code    int         `json:"code" example:"200"`     // the HTTP status
	Result  interface{} `json:"result"`                 // whatever object or array of objects
}

// JSONResultCreated returns a json object
type JSONResultCreated struct {
	Success bool        `json:"success" example:"true"` // wether the business action is successfull or not
	Code    int         `json:"code" example:"201"`     // the HTTP status
	Result  interface{} `json:"result"`                 // whatever object or array of objects
}

// JSONNoContentResult returns a json object
type JSONNoContentResult struct {
	Success bool `json:"success" example:"true"` // wether the business action is successfull or not
	Code    int  `json:"code" example:"204"`     // the HTTP status
}

// NewPagedResult sets the result body to context
func NewPagedResult(ctx *gin.Context, status int, totalCount int, response interface{}) {
	resp := JSONPagedResult{
		TotalCount: totalCount,
		JSONResult: JSONResult{
			Success: true,
			Code:    status,
			Result:  response,
		},
	}
	ctx.JSON(status, resp)
}

// JSONPagedResult returns a json object
type JSONPagedResult struct {
	TotalCount int `json:"totalCount" example:"102"` // the total number of records
	JSONResult
}
