package database

func (t *Users) GetItems() interface{} {
	return &t.Items
}

func (t *Users) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Roles) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Issues) GetItem() interface{} {
	return &t.Items[0]
}

func (t *IssuesLists) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Clients) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Organisations) GetItem() interface{} {
	return &t.Items[0]
}

func (t *WorkGroups) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Developers) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Statuses) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Stages) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Managers) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Projects) GetItem() interface{} {
	return &t.Items[0]
}

func (t *Roles) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Managers) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Projects) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *ProjectsPreview) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Stages) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Statuses) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Tasks) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *WorkGroups) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *WorkGroupLists) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Organisations) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Modules) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Issues) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *IssuesLists) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Users) GetPrimaryKey() string {
	return GetTag(t.Items)
}
func (t *Clients) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
func (t *Developers) GetPrimaryKey() string {
	return GetTag(t.Items[0])
}
