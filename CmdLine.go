package ineed

//ineed cmd + ineed config = git cmds
var	ineedCmdPatterns = CmdPatterns{
	"status" : []string{ "Alias" },
	"commit" : []string{ "CommitMsg","Alias" },
	"pull" : []string{ "Alias" },
	"push" : []string{ "Alias" },
}

var	gitCmdTmpls = CmdTmpls{
	"status" : "-C {{.RepoPath}} status",
	"commit" : "-C {{.RepoPath}} commit -a -m \"{{.CommitMsg}}\"",
	"pull" : "-C {{.RepoPath}} pull {{.Remote}} {{.Branch}}",
	"push" : "-C {{.RepoPath}} push {{.Remote}} {{.Branch}}",
}

type CmdLine struct{
    Cmd string
    Alias string
    CommitMsg string
    RepoPath string
    Remote string
    Branch string
}




