package Controllers

type Controller interface {
	get()
	index()
	store()
	update()
	delete()
}
