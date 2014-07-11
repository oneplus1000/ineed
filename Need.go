package ineed

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"path"
	"errors"
	"reflect"
	"text/template"
	"bytes"
	"strings"
)

const CONFIG_FILENAME = "ineed.json"

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
	err = json.Unmarshal(data,&me.ConfigInfo)
	if err != nil {
		return err
	}
	return nil
}

func (me * Need) ParseIneedCmd(cmdline *CmdLine,cmdtokens []string) error{

	var cmdKey string
	var cmdPattern []string

	var isContain bool
	isContain,cmdPattern =  ineedCmdPatterns.GetValByKey(cmdtokens[0])
	if !isContain {
		return errors.New("No ineed command match")
	}
	cmdKey = cmdtokens[0]

	refl :=	reflect.ValueOf(cmdline).Elem()
	for i,val := range cmdtokens {
		if i == 0 {
			cmdline.Cmd = cmdKey
			continue
		}
		index := i - 1
		param := cmdPattern[index]
		objparam := refl.FieldByName(param)
		if objparam.IsValid() && objparam.CanSet() {
			objparam.SetString(val)
		}
	}
	return nil
}

func (me * Need) Run(cmdtokens []string) error {

	var err error

	for _,val := range me.ConfigInfo.Needs	{

		var cmdline CmdLine
		var gitcmd string
		me.ParseIneedCmd(&cmdline,cmdtokens)

		fmt.Printf(">>>>%+v\n",cmdline)

		err = me.BindNeedConfigToCmdLine(&cmdline,&val)
		if err != nil {
			return err
		}

		gitcmd,err = me.CmdLine(&cmdline)
		if err != nil {
			return err
		}

		cmdtokens := strings.Split(gitcmd," ")
		out, err := exec.Command("git",cmdtokens...).Output()
		if err != nil {
			return err
		}
		me.Print(gitcmd,string(out))
	}

	return nil
}


func (me *Need) Print(gitcmd string,outOk string){
	fmt.Printf("\n===================================================================\n")
	fmt.Printf("git %s",gitcmd)
	fmt.Printf("\n===================================================================\n")
	fmt.Printf("%s\n",outOk)
}

func (me *Need) CmdLine(cmdline *CmdLine) (string,error) {

	var isContain bool
	var cmdTmpl string
	var keyTmpl string
	var err error
	var tmpl *template.Template

	isContain,cmdTmpl =	gitCmdTmpls.GetValByKey(cmdline.Cmd)
	if ! isContain {
		return "",errors.New("No git command match")
	}

	//fmt.Printf(">>>>>>cmdTmpl:%s\n",cmdTmpl)
	keyTmpl = cmdline.Cmd
	tmpl, err = template.New(keyTmpl).Parse(cmdTmpl)
	if err != nil {
		return "",err
	}

	//fmt.Printf("******cmdTmpl:%s\n",cmdTmpl)
	buff := new(bytes.Buffer)
	err = tmpl.Execute(buff,cmdline)
	if err != nil {
		return "",err
	}
	//fmt.Printf(">>>>>>%s %+v\n",buff.String(),cmdline)
	return buff.String(),nil
}

func (me *Need) BindNeedConfigToCmdLine(cmdline *CmdLine,configneed *ConfigNeed) error {

	cmdline.RepoPath = path.Clean(me.CurrentPath + "/" + configneed.Path)
	if cmdline.RepoPath == "" {
		cmdline.RepoPath = "./"
	}

	cmdline.Remote = configneed.Remote
	cmdline.Branch = configneed.Branch
	
	return nil
}

