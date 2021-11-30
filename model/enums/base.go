package enums

type BaseStatus int8

const (
	StatusFalse = iota
	StatusTrue
)

var _BaseStatus = []BaseStatus{StatusFalse, StatusTrue}

func (static BaseStatus) Is() bool {
	for _, v := range _BaseStatus {
		if v == static {
			return true
		}
	}

	return false
}

func (static BaseStatus) Label() string {
	m := map[BaseStatus]string{
		StatusFalse: "禁用",
		StatusTrue:  "启用",
	}

	return m[static]
}

func (static BaseStatus) Value() int8 {
	return int8(static)
}

func BaseStatusMap() []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, v := range _BaseStatus {
		m = append(m, map[string]interface{}{
			"label": v.Label(),
			"value": v.Value(),
		})
	}

	return m
}
