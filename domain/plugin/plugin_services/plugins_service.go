package plugin_services

import (
	"database/sql"
	"fmt"
	"github.com/gocms-io/gocms/routes"
	"github.com/gocms-io/gocms/domain/plugin/plugin_model"
	"github.com/gocms-io/gocms/init/repository"
)

type IPluginsService interface {
	StartActivePlugins() error
	RegisterActivePluginRoutes(routes *routes.Routes) error
	GetDatabasePlugins() (map[string]*plugin_model.PluginDatabaseRecord, error)
	RefreshInstalledPlugins() error
	GetActivePlugins() map[string]*plugin_model.Plugin
}

type PluginsService struct {
	repositoriesGroup *repository.RepositoriesGroup
	installedPlugins  map[string]*plugin_model.Plugin
	activePlugins     map[string]*plugin_model.Plugin
}

func DefaultPluginsService(rg *repository.RepositoriesGroup) *PluginsService {

	pluginsService := &PluginsService{
		repositoriesGroup: rg,
		installedPlugins:  make(map[string]*plugin_model.Plugin),
		activePlugins:     make(map[string]*plugin_model.Plugin),
	}

	return pluginsService

}

func (ps *PluginsService) GetDatabasePlugins() (map[string]*plugin_model.PluginDatabaseRecord, error) {
	databasePluginRecords, err := ps.repositoriesGroup.PluginRepository.GetDatabasePlugins()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No plugins referenced in database.\n")
			return nil, err
		}
		fmt.Printf("Error getting database plugins: %v\n", err.Error())
		return nil, err
	}

	databasePluginsMap := make(map[string]*plugin_model.PluginDatabaseRecord)
	for _, databasePlugin := range databasePluginRecords {
		databasePluginsMap[databasePlugin.PluginId] = databasePlugin
	}

	return databasePluginsMap, nil
}

func (ps *PluginsService) GetActivePlugins() map[string]*plugin_model.Plugin {
	return ps.activePlugins
}
