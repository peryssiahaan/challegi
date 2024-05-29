package database

func CreateGroup(group *Group) error {
	result := db.Create(group)
	return result.Error
}

func GetGroups() ([]Group, error) {
	var groups []Group
	result := db.Find(&groups)
	return groups, result.Error
}
