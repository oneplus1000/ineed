package ineed

//ineed cmd + ineed config = git cmds
var	ineedCmdTmpls = map[string][]string{
	"status" : []string{ "Alias" },
	"commit" : []string{ "CommitMsg","Alias" },
	"pull" : []string{ "Alias" },
	"push" : []string{ "Alias" },
}

var	gitCmdTmpls = CmdTmpls{
	"status" : "git -C {{RepoPath}} status",
	"commit" : "git -C {{RepoPath}} commit -a -m \"{{CommitMsg}}\"",
	"pull" : "git -C {{RepoPath}} pull {{Remote}} {{Branch}}",
	"push" : "git -C {{RepoPath}} push {{Remote}} {{Branch}}",
}

type CmdLine struct{
    Cmd string
    Alias string
    CommitMsg string
    RepoPath string
    Remote string
    Branch string
}


type CmdTmpls map[string]string

func (me * CmdTmpls) GetValByKey(key string) (string,error){
    return "",nil
}