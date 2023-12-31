package adapter

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"try_mongo/model"
	mongodb "try_mongo/mongo"
)

type MdDataController struct {
	dataSource *mongodb.DateSource
}

func NewMdDataController(dataSource *mongodb.DateSource) *MdDataController {
	return &MdDataController{
		dataSource: dataSource,
	}
}

func (m *MdDataController) GetMdData(c echo.Context) error {
	filename := c.Param("filename")

	var mdfile model.MdData
	mdfile.Title = filename

	mdBao := m.dataSource.MdDataDao()

	err := mdBao.FindOne(&mdfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "not found",
		})
	}
	return c.JSON(http.StatusOK, mdfile)
}

func (m *MdDataController) GetMdDatas(c echo.Context) error {
	filename := c.QueryParam("filename")

	var mdfile model.MdData
	mdfile.Title = filename

	mdBao := m.dataSource.MdDataDao()
	mdfiles, err := mdBao.FindMany(&mdfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "not found",
		})
	}
	return c.JSON(http.StatusOK, mdfiles)
}

func (m *MdDataController) AddMdData(c echo.Context) error {
	var mdfile model.MdData
	if err := c.Bind(&mdfile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	mdBao := m.dataSource.MdDataDao()

	// check if exist
	err := mdBao.FindOne(&mdfile)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "already exists",
		})
	}

	err = mdBao.AddOne(&mdfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "add failed",
		})
	}
	return c.JSON(http.StatusOK, mdfile)
}

func (m *MdDataController) UpdateMdData(c echo.Context) error {
	var mdfile model.MdData
	if err := c.Bind(&mdfile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	mdBao := m.dataSource.MdDataDao()
	err := mdBao.UpdateOne(&mdfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "update failed",
		})
	}
	return c.JSON(http.StatusOK, mdfile)
}

func (m *MdDataController) DeleteMdData(c echo.Context) error {
	var mdfile model.MdData
	if err := c.Bind(&mdfile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	mdBao := m.dataSource.MdDataDao()
	err := mdBao.DeleteOne(&mdfile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "delete failed",
		})
	}
	return c.JSON(http.StatusOK, mdfile)
}
