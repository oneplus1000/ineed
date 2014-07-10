package ineed

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"path"
	"errors"
	"reflect"
)

const CONFIG_FILENAME = "ineed.json"

//ineed cmd + ineed config = git cmds
var	ineedCmdTmpls = map[string][]string{
	"status" : []string{ "Alias" },
	"commit" : []string{ "CommitMsg","Alias" },
	"pull" : []string{ "Alias" },
	"push" : []string{ "Alias" },
}

var	gitCmdTmpls = map[string]string{
	"status" : "git -C {{RepoPath}} status",
	"commit" : "git -C {{RepoPath}} commit -a -m \"{{CommitMsg}}\"",
	"pull" : "git -C {{RepoPath}} pull {{Remote}} {{Branch}}",
	"push" : "git -C {{RepoPath}} push {{Remote}} {{Branch}}",
}


type Need struct{
	CurrentPath string
	ConfigInfo Config
}

func (me * Need) Init(currentPath string) error {
	me.CurrentPath = currentPath
	fileconfig := currentPath + "/" + CONFIG_FILENAME
	data,err := ioutil.ReadFile(fileconfig)
	if err != nil {
		return err
	}
	//fmt.Printf("%s\n",string(data))
	err = json.Unmarshal(data,&me.ConfigInfo)
	if err != nil {
		return err
	}
	return nil
}

func (me * Need) ParseIneedCmd(cmdline *CmdLine,cmdtokens []string) error{
	var cmdKey string 
	var cmdTmpl []string
	for key,val := range ineedCmdTmpls {
		if key == cmdtokens[0] {
			cmdKey = key
			cmdTmpl = val
			break
		}
	} 
	
	if cmdKey == "" {
		return errors.New("No ineed command match")
	}
	refl :=	reflect.ValueOf(cmdline).Elem()
	for i,val := range cmdtokens {
		if i == 0 {
			cmdline.Cmd = cmdKey
			continue
		}
		index := i - 1
		param := cmdTmpl[index]
		objparam := refl.FieldByName(param)
		if objparam.IsValid() && objparam.CanSet() {
			objparam.SetString(val)
		}
	}
	return nil
}

func (me * Need) Run(cmdtokens []string) error {
		
	var cmdline CmdLine
	me.ParseIneedCmd(&cmdline,cmdtokens)
	fmt.Printf("%+v",cmdline)
	for _,val := range me.ConfigInfo.Needs	{
		repopath := path.Clean(me.CurrentPath + "/" + val.Path)
		fmt.Printf("\n====================================\n%s\n====================================\n",repopath)
		out, err := exec.Command("git","-C" , repopath , cmdtokens[0]).Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n",out)
	}

	return nil
}


