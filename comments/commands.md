# Basics-02

>  wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
> 
>  sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
> 
>  mkdir ~/GolandProjects/rebrain-go
> 
>  echo "export GOPATH=~/GolandProjects/rebrain-go" >> ~/.profile
> 
>  echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
> 
>  echo "export GOPATH=~/GolandProjects/rebrain-go" >> ~/.profile
>
>  mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin
> 
>  echo "export PATH=$PATH:$GOPATH/bin" >> ~/.profile
> 
>  source ~/.profile
> 
>  echo $GOPATH
> 
> go version
> 
> go run $GOPATH/src/GO-01/hello.go