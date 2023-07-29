package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	kind "sigs.k8s.io/kind/pkg/cluster"

	clusterprovider "github.com/kubestellar/kubestellar/pkg/clustermanager/providerclient"
)

var kindProvider *kind.Provider

func main() {
	router := gin.Default()
	kindProvider = kind.NewProvider()
	router.POST("/create/:name", Create)
	router.POST("/delete/:name", Delete)
	router.GET("/get/:name", Get)
	router.GET("/list_clusters", ListClusters)
	router.GET("/list_clusters_names", ListClustersNames)

	router.Run("localhost:8087")
}

func Get(c *gin.Context) {
	lcName := c.Param("name")

	cfg, err := kindProvider.KubeConfig(lcName, false)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Logical cluster not found"})
	}
	lcInfo := clusterprovider.LogicalClusterInfo{
		Name:   lcName,
		Config: cfg,
	}
	c.IndentedJSON(http.StatusOK, lcInfo)
}

func ListClusters(c *gin.Context) {
	lcNames, err := listClustersNames()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error listing logical cluster names"})
	}

	lcInfoList := make([]clusterprovider.LogicalClusterInfo, 0, len(lcNames))
	for _, lcName := range lcNames {
		cfg, err := kindProvider.KubeConfig(lcName, false)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error getting logical cluster config"})
		}

		lcInfoList = append(lcInfoList, clusterprovider.LogicalClusterInfo{
			Name:   lcName,
			Config: cfg,
		})
	}
	c.IndentedJSON(http.StatusOK, lcInfoList)
}

func listClustersNames() ([]string, error) {
	list, err := kindProvider.List()
	if err != nil {
		return nil, err
	}
	logicalNameList := make([]string, 0, len(list))
	logicalNameList = append(logicalNameList, list...)
	return logicalNameList, nil
}

func ListClustersNames(c *gin.Context) {
	logicalNameList, err := listClustersNames()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error listing logical cluster names"})
	}
	c.IndentedJSON(http.StatusOK, logicalNameList)
}

func Create(c *gin.Context) {
	lcName := c.Param("name")
	err := kindProvider.Create(lcName)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to create logical cluster"})
	}

	c.IndentedJSON(http.StatusCreated, lcName)
}

func Delete(c *gin.Context) {
	lcName := c.Param("name")
	err := kindProvider.Delete(lcName, "")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to delete logical cluster"})
	}

	c.IndentedJSON(http.StatusCreated, lcName)
}
