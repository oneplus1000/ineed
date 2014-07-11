package ineed

type CmdPatterns map[string][]string

func (me CmdPatterns) GetValByKey(key string) (bool,[]string) {
	for key,val := range me {
		if key == key {
			return true,val
		}
	}
	return false,nil
}


