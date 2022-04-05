package define

import "time"

type Record interface{}

type RecordField interface {
	ToField() map[string]interface{}
}

type Project struct {
	ID          string
	Name        string
	Description string
	WebUrl      string
	TagList     []string
	Topics      []string
}

func (p *Project) ToField() map[string]interface{} {
	return map[string]interface{}{
		"ID":          p.ID,
		"Name":        p.Name,
		"Description": p.Description,
		"WebUrl": map[string]string{
			"link": p.WebUrl,
			"text": p.WebUrl,
		},
	}
}

type Milestone struct {
	ID string `json:"ID"`
}

func (m *Milestone) ToField() map[string]interface{} {
	return map[string]interface{}{
		"ID": m.ID,
	}
}

type Issue struct {
	ID           string
	Title        string
	Description  string
	State        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DueDate      time.Time
	HealthStatus string
	ProjectRef   *ProjectRef
}

type ProjectRef struct {
	Text      string
	Type      string
	TableID   string
	RecordIDs []string
}

func (i *Issue) ToField() map[string]interface{} {
	res := map[string]interface{}{
		"ID":            i.ID,
		"title":         i.Title,
		"description":   i.Description,
		"state":         i.State,
		"created_at":    i.CreatedAt.UnixNano() / 1e6,
		"updated_at":    i.UpdatedAt.UnixNano() / 1e6,
		"due_date":      i.DueDate.UnixNano() / 1e6,
		"health_status": i.HealthStatus,
	}
	if i.ProjectRef != nil {
		// TODO struct oncall
		res["Project"] = []map[string]interface{}{
			{
				"text":       i.ProjectRef.Text,
				"type":       i.ProjectRef.Type,
				"table_id":   i.ProjectRef.TableID,
				"record_ids": i.ProjectRef.RecordIDs,
			},
		}
	}
	return res
}

type Member struct {
	ID string `json:"ID"`
}

func (m *Member) ToField() map[string]interface{} {
	return map[string]interface{}{
		"ID": m.ID,
	}
}
