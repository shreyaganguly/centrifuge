# Centrifuge

Centrifuge is an application to extract all the issues of a github organization and save it in a csv file.

### Run

```
$ go build
$ ./centrifuge -token <Access Token of the Github Account> -org <Organization's name for which the issues are to be extracted>
```

Filtering based on ```milestone```, ```labels```, ```status``` is also available. Do

```
./centrifuge -h

Usage of ./centrifuge:
  -format string
    	Format to store after extracting issue details(json|csv|html|md) (default "json")
  -labels string
    	Filtering based on the labels marked to the issues(give comma-separated values)
  -milestone string
    	Filtering based on the milestone assigned to the issues
  -name string
    	file to save the extarcted issues, if empty it will print to stdout
  -org string
    	Organization for which issues are to be searched
  -status string
    	Filtering based on the status of the github issue(value can be either open or closed or all)(DEFAULT:all) (default "all")
  -token string
    	Access Token of your github account, if this flag is NOT set it will read from environemnt variable GITHUB_TOKEN
```

#### Start extracting!!! :)
