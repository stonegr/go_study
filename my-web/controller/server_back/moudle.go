package server_back

import "mime/multipart"

type SERVER_UP_FIELED struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
	Pwd  string                `form:"pwd" binding:"required"`
}
