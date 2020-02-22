# cakecloud-cli
How to use cakecloud CLI

1.Install package
```
git clone
cd cakecloud-cli
make
```

2.Log in with your cakecloud username and password
```
cakecloud login
```

2.Start SSH with VM name
```
cakecloud ssh $VM_NAME
```
3.Autocomplete
To enable autocomplete, follow instruction below:
 - Open your ~/.zshrc file and add cakecloud as a plugins
 Example:
```
plugins=(git) ---> plugins=(git cakecloud)
```
 - Run update to fetch list of instance name 
```
cakecloud update
```
