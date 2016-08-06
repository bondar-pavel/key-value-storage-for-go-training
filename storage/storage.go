package storage


type Querier interface {
	Query(string, Command) Command
}

type Command struct {
        Cmd string
        Key string
        Value string
        Exists bool
        ReplyChan chan Command
}

type Storage struct {
	Map map[string]string
	Chan chan *Command
}

func (s *Storage) Init() {
	s.Map = make(map[string]string)
	// All SET commands are non-blocking, since does not require any reply,
	// so use buffered channels
        s.Chan = make(chan *Command, 10)
	go s.process()
}

func (s *Storage) set(key, value string) {
	s.Map[key] = value
}

func (s *Storage) get(key string) (string, bool) {
	value, ok := s.Map[key]
	return value, ok
}

func (s *Storage) del(key string) (deleted bool) {
	_, exists := s.Map[key]
	if exists == true {
		delete(s.Map, key)
		deleted = true
	}
	return
}

func (s *Storage) process() {
        for command :=  range s.Chan {
                if command.Cmd == "set" {
                        s.set(command.Key, command.Value)
                } else if command.Cmd == "get" {
                        command.Value, command.Exists = s.get(command.Key)
                        command.ReplyChan <- *command
                } else if command.Cmd == "del" {
                        command.Exists = s.del(command.Key)
                        command.ReplyChan <- *command
                }
        }
}

func (s *Storage) Query(action string, cmd Command) Command {
	cmd.Cmd = action
	if action != "set" {
		cmd.ReplyChan = make(chan Command, 1)
	}
	s.Chan <- &cmd
	if action != "set" {
		cmd = <-cmd.ReplyChan
	}
	return cmd
}

