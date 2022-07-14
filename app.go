package app

type directory struct {
	Root   string
	Config string
	Data   string
	Assets string
}

var Dir = &directory{}

func (d *directory) SetRoot(root string) {
	d.Root = root
	d.Config = root + "config/"
	d.Data = root + "data/"
	d.Assets = root + "assets/"
}
