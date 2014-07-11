package ineed

//ineed cmd + ineed config = git cmds
var	ineedCmdPatterns = CmdPatterns{
	"status" : []string{ "Alias" , "Cmd" },
	"commit" : []string{ "Alias" , "Cmd" ,  "CommitMsg" },
	"pull" : []string{ "Alias" , "Cmd" },
	"push" : []string{ "Alias" , "Cmd"},
	"diff" : []string{ "Alias" , "Cmd"},
	"add" : []string{ "Alias" , "Cmd"},
}

var	gitCmdTmpls = CmdTmpls{
	"status" : "-C {{.RepoPath}} status",
	"commit" : "-C {{.RepoPath}} commit -a -m \"{{.CommitMsg}}\"",
	"pull" : "-C {{.RepoPath}} pull {{.Remote}} {{.Branch}}",
	"push" : "-C {{.RepoPath}} push {{.Remote}} {{.Branch}}",
	"diff" : "-C {{.RepoPath}} diff",
	"add" : "-C {{.RepoPath}} add .",
}

type CmdLine struct{
    Cmd string
    Alias string
    CommitMsg string
    RepoPath string
    Remote string
    Branch string
}




