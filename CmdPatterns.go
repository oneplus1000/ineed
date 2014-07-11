package ineed

type CmdPatterns map[string][]string

func (me CmdPatterns) GetValByKey(k string) (bool,[]string) {
	for key,val := range me {
		if key == k {
			return true,val
		}
	}
	return false,nil
}


