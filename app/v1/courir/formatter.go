package courir

type CourirFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatCourir(courir Courir) CourirFormatter {
	formatter := CourirFormatter{}
	formatter.ID = courir.ID
	formatter.Name = courir.Name

	return formatter
}

func FormatCourirs(courirs []Courir) []CourirFormatter {
	var formatter []CourirFormatter

	for _, courir := range courirs {
		formatter = append(formatter, FormatCourir(courir))
	}

	return formatter
}
