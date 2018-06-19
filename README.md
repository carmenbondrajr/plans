# Healthcare Plan Information
Provided with a CSV of User information and Plan information, this program  
will return eligible Plans and Residents to a given zip code.  

This program was created using Go 1.10.3, but should be fine for any version >1.8.

## Build
To build, execute the following commands:  
`. ./plans.sh`  
`make build`

## Test
To test, execute the following commands:  
(sourcing `plans.sh` is only necessary once per shell session)  
`. ./plans.sh`  
`make test`

## Run
`./bin/plans -zip=[ZIP_CODE]`  
Where [ZIP_CODE] is a five-digit zip.  
Be sure to build the program first.  

This program will default to the provided `plans.csv` and `users.csv` files.   
However, it is possible to provide your own by utilizing the following flags:  
`./bin/plans -zip=[ZIP_CODE] -plans=[PLANS_FILE.CSV] -users=[USERS_FILE.CSV]`

### Troubleshooting
Utilizing `go env` -   
- Verify that `[path_to_go]/go/bin` is on your `$PATH` variable
- Verify that `$GOROOT` is set to the directory in which you extracted Go. 
- Verify that `$GOPATH` is set to this directory.