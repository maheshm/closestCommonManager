package memory

type Entry struct {
	Id int
	Data map[string]interface{}
}

var current_id = 1
var memory_store map[int]*Entry

func store_data(data map[string]interface{}) map[string]interface{} {
	var data_to_be_stored map[string]interface{}
	data_to_be_stored = make(map[string]interface{})

	for k,v := range data {
		data_to_be_stored[k] = v
	}

	return data_to_be_stored
}

func Get() {}

func Set(data map[string]interface{}) *Entry {
	entry_data := store_data(data)
	// get lock here
	new_entry := &Entry{Id: current_id, Data: entry_data}
	if memory_store == nil {
		memory_store = make(map[int]*Entry)
	}
	memory_store[current_id] = new_entry
	current_id += 1
	//release lock
	return new_entry
}

func Update(id int, data map[string]interface{}) *Entry {
	return memory_store[0]
}

func Delete () {}
