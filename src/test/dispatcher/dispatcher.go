package dispatcher

type dispatcher struct{
	d map[string]func(msg interface{}) (interface{},error)
}

var instance *dispatcher

func Instance() *dispatcher{
	if instance == nil{
		instance = &dispatcher{}
		instance.d = make(map[string]func(msg interface{}) (interface{},error),10)
	}
	return instance
}

func (d *dispatcher)Add(name string,f func(msg interface{}) (interface{},error)){
	d.d[name] = f
}

func (d *dispatcher)Update(name string,f func(msg interface{}) (interface{},error)){
	d.Add(name,f)
}

func (d *dispatcher)Del(name string){
	delete(d.d,name)
}

func (d *dispatcher)Select(name string) func(msg interface{}) (interface{},error){
	f,found := d.d[name]
	if found{
		return f
	}

	return nil
}





