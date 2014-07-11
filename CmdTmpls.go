package ineed

type CmdTmpls map[string]string

func (me CmdTmpls) GetValByKey(key string) (bool,string){
	for key,val := range me {
		if key == key {
			return true,val
		}
	}
	return false,""
}
