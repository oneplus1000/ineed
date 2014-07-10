package ineed



type Config struct{
	Needs []ConfigNeed
}

type ConfigNeed struct {
	Alias string
	Path string 
	Remote string
	Branch string
}


//-----------
/*
{
	"needs":[
		{	
			"alias" : "Test01",
			"path" : "",
			"remote" : "origin",
			"branch" : "master"
		},
		{	
			"alias" : "Test02",
			"path" : "../Test02",
			"remote" : "origin",
			"branch" : "master"
		}
	]
}
*/
