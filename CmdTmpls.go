package ineed

type CmdTmpls map[string]string

func (me CmdTmpls) GetValByKey(k string) (bool,string){
	for key,val := range me {
		if key == k {
			return true,val
		}
	}
	return false,""
}
