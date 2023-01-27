package category

type CategoryFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{}
	formatter.ID = category.ID
	formatter.Name = category.Name

	return formatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	formatter := []CategoryFormatter{}

	for _, category := range categories {
		formatter = append(formatter, FormatCategory(category))
	}

	return formatter
}
