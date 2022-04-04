package define

type Define struct {
	App *App `mapstructure:"app,omitempty" yaml:"app,omitempty" json:"app,omitempty"`
}

func (app *App) GetTableMap() map[string]*Table {
	res := make(map[string]*Table)
	for _, t := range app.Tables {
		res[t.Name] = t
	}
	return res
}
func (table *Table) GetViewMap() map[string]*View {
	res := make(map[string]*View)
	for _, v := range table.Views {
		res[v.ViewName] = v
	}
	return res
}

type App struct {
	Name   string   `mapstructure:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	Tables []*Table `mapstructure:"tables,omitempty" yaml:"tables,omitempty" json:"tables,omitempty"`
}

type Table struct {
	Revision int      `mapstructure:"revision,omitempty" yaml:"revision,omitempty" json:"revision,omitempty"`
	Name     string   `mapstructure:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	Fields   []*Field `mapstructure:"fields,omitempty" yaml:"fields,omitempty" json:"fields,omitempty"`
	Views    []*View  `mapstructure:"views,omitempty" yaml:"views,omitempty" json:"views,omitempty"`
}
type View struct {
	ViewName string `mapstructure:"view_name,omitempty" yaml:"view_name,omitempty" json:"view_name,omitempty"`
	ViewType string `mapstructure:"view_type,omitempty" yaml:"view_type,omitempty" json:"view_type,omitempty"`
}

type Field struct {
	FieldName string         `mapstructure:"field_name,omitempty" yaml:"field_name,omitempty" json:"field_name,omitempty"`
	Property  *FieldProperty `mapstructure:"property,omitempty" yaml:"property,omitempty" json:"property,omitempty"`
	Type      int            `mapstructure:"type,omitempty"  yaml:"type,omitempty" json:"type,omitempty"`
}

type FieldProperty struct {
	Options    []*FieldPropertyOption `mapstructure:"options,omitempty" yaml:"options,omitempty" json:"options,omitempty"`
	Formatter  string                 `mapstructure:"formatter,omitempty" yaml:"formatter,omitempty" json:"formatter,omitempty"`
	DateFormat string                 `mapstructure:"date_format,omitempty" yaml:"date_format,omitempty" json:"date_format,omitempty"`
	TimeFormat string                 `mapstructure:"time_format,omitempty" yaml:"time_format,omitempty" json:"time_format,omitempty"`
	AutoFill   bool                   `mapstructure:"auto_fill,omitempty" yaml:"auto_fill,omitempty" json:"auto_fill,omitempty"`
	Multiple   bool                   `mapstructure:"multiple,omitempty" yaml:"multiple,omitempty" json:"multiple,omitempty"`
	TableId    string                 `mapstructure:"table_id,omitempty" yaml:"table_id,omitempty" json:"table_id,omitempty"`
	ViewId     string                 `mapstructure:"view_id,omitempty" yaml:"view_id,omitempty"  json:"view_id,omitempty"`
	Fields     []string               `mapstructure:"fields,omitempty" yaml:"fields,omitempty" json:"fields,omitempty"`
}

type FieldPropertyOption struct {
	Name string `mapstructure:"name,omitempty" yaml:"name,omitempty" json:"name,omitempty"`
	Id   string `mapstructure:"id,omitempty" yaml:"id,omitempty" json:"id,omitempty"`
}
