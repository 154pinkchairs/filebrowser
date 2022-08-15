package groups

import (
	os/user
	"github.com/filebrowser/filebrowser/v2/errors"
	"github.com/filebrowser/filebrowser/v2/files"
	"github.com/filebrowser/filebrowser/v2/rules"
)

// ViewMode describes a view mode.
type ViewMode string

const (
	ListViewMode   ViewMode = "list"
	MosaicViewMode ViewMode = "mosaic"
)

// Group describes a group.
type Group struct {
	GID           uint          `storm:"gid,increment" json:"gid"`
	GroupName     string        `storm:"unique" json:"groupname"`
	Scope        string        `json:"scope"`
	MaxFileSize  int		   `json:"maxFileSize"`
	MaxUserCount int		   `json:"maxUserCount"`
	Locale       string        `json:"locale"`
	LockPassword bool          `json:"lockPassword"`
	ViewMode     ViewMode      `json:"viewMode"`
	SingleClick  bool          `json:"singleClick"`
	Perm         Permissions   `json:"perm"`
	Commands     []string      `json:"commands"`
	Sorting      files.Sorting `json:"sorting"`
	Fs           afero.Fs      `json:"-" yaml:"-"`
	Rules        []rules.Rule  `json:"rules"`
	HideDotfiles bool          `json:"hideDotfiles"`
	DateFormat   bool          `json:"dateFormat"`
}

// GetRules implements rules.Provider.
func (g *Group) GetRules() []rules.Rule {
	return g.Rules
}

var checkableFields = []string{
	"GroupName",
	"MaxFileSize",
	"MaxUserCount",
	"Scope",
	"ViewMode",
	"Commands",
	"Sorting",
	"Rules",
}

// Clean cleans up a user and verifies if all its fields
// are alright to be saved.
//nolint:gocyclo
func (g *Group) Clean(baseScope string, fields ...string) error {
	if len(fields) == 0 {
		fields = checkableFields
	}

	for _, field := range fields {
		switch field {
		case "Username":
			if g.GroupName == "" {
				return errors.ErrEmptyGroupName
			}
		case "ViewMode":
			if g.ViewMode == "" {
				g.ViewMode = ListViewMode
			}
		case "Commands":
			if g.Commands == nil {
				g.Commands = []string{}
			}
		case "Sorting":
			if g.Sorting.By == "" {
				g.Sorting.By = "name"
			}
		case "Rules":
			if g.Rules == nil {
				g.Rules = []rules.Rule{}
			}
		}
	}

	if u.Fs == nil {
		scope := u.Scope
		scope = filepath.Join(baseScope, filepath.Join("/", scope)) //nolint:gocritic
		u.Fs = afero.NewBasePathFs(afero.NewOsFs(), scope)
	}

	return nil
}

// CanExecute checks if an user in the group can execute a specific command.
func (g *Group) CanExecute(command string) bool {
	if !g.Perm.Execute {
		return false
	}

	for _, cmd := range u.Commands {
		if regexp.MustCompile(cmd).MatchString(command) {
			return true
		}
	}

	return false
}
