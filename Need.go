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

	//var cmdKey string
	var cmdPattern []string

	var isContain bool
	isContain,cmdPattern =  ineedCmdPatterns.GetValByKey(cmdtokens[1])
	if !isContain {
		return errors.New("No ineed command match")
	}


	refl :=	reflect.ValueOf(cmdline).Elem()
	for i,val := range cmdtokens {
		index := i
		param := cmdPattern[index]
		objparam := refl.FieldByName(param)
		if objparam.IsValid() && objparam.CanSet() {
			objparam.SetString(val)
		}
	}
	//fmt.Printf(">>>%s\n",cmdline)
	return nil
}

func (me * Need) Run(cmdtokens []string) error {

	//fmt.Printf("=================\n")
	var err error
	for _,val := range me.ConfigInfo.Needs	{

		var cmdline CmdLine
		var gitcmd string

		//fmt.Printf("AAAA>>>>%+v\n",cmdline)
		err = me.ParseIneedCmd(&cmdline,cmdtokens)
		if err != nil {
			return err
		}
		//fmt.Printf("BBBB>>>>%+v\n",cmdline)

		err = me.BindNeedConfigToCmdLine(&cmdline,&val)
		if err != nil {
			return err
		}

		if cmdline.Alias != "-all" {
			if strings.ToLower(cmdline.Alias)!= strings.ToLower(val.Alias) {
				continue
			}
		}

		gitcmd,err = me.CmdLine(&cmdline)
		if err != nil {
			return err
		}

		me.PrintHeader(gitcmd+"     # "+val.Alias+"")
		//fmt.Printf("DDDD>>%s\n\n",gitcmd)
		cmdtokens := strings.Split(gitcmd," ")
		out, _ := exec.Command("git",cmdtokens...).Output()
		me.Print(string(out))

	}

	return nil
}

func (me *Need) PrintHeader(gitcmd string){

	fmt.Printf("\n========================================================================\n")
	fmt.Printf("git %s",gitcmd)
	fmt.Printf("\n========================================================================\n")

}

func (me *Need) Print(outOk string){

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


	keyTmpl = cmdline.Cmd
	tmpl, err = template.New(keyTmpl).Parse(cmdTmpl)
	if err != nil {
		return "",err
	}

	//fmt.Printf(">>>>>>keyTmpl: %s , cmdTmpl: %s\n",keyTmpl,cmdTmpl)

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

