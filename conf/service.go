package conf

type ServiceConf struct {
	Image    string `yaml:"image"`  // 可以直接使用 docker pull 的名称
	Id       string `yaml:"id"`     // 唯一注册 ID
	Expose   []int  `yaml:"expose"` // 暴露端口列表，最少一个
	Icon     string `yaml:"icon"`   // 显示图标(备用，暂不确定如何使用此字段)
	Title    string `yaml:"title"`  // 显示名称
	MapPorts []int  `yaml:"-"`      // 实际映射端口列表
}

func GetServiceConf(path string) *ServiceConf {
	return &ServiceConf{
		Image:    "",
		Id:       "",
		Expose:   make([]int, 0),
		Title:    "",
		MapPorts: make([]int, 0),
	}
}
