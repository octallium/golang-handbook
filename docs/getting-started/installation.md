# Installing Go
<hr>


## {++ How-To++} Install Go on your local machine 

## Installing on Mac OSx:

???+ question "Installing on Mac OSx"
        
    Installing Go on your mac is pretty straight forward, choose any one method: -

    A) Installing using package installer directly from [golang.org](https://golang.org/dl/)

    OR
    
    B) Using Homebrew

    ???+ note "A) Installing using the package installer"
            
        Package installer method is the simplest way to get you started, just download the .pkg file by clicking on [this link](https://golang.org/dl/) and install it just as you would install any other program on your mac.

        The installer will automatically create everything for you including setting up your environment variables.

        Installer will automatically install it to

            usr/local/go directory

        After installing read "Checking the installation" section below to make sure everything works.

    ??? abstract "B) Using Homebrew"
            
        Before using this method, make sure you have homebrew installed, if it is not installed checkout this [link](https://brew.sh/) and install it.

        Open your terminal and type the following:
            
            brew install go

        brew will automatically fetch the latest binaries and install it for you, make sure they are installed to

            usr/local/go directory

        Its now time to setup your environment variables, check the "Setting up environment variables" to complete the installation and start using Go.

## Installing on FreeBSD & Linux:

??? question "Installing on FreeBSD & Linux"

    Download the archive from [golang.org](https://golang.org/dl/) and extract it to usr/local, for extracting open the terminal in the downloaded directory and type the following:

    If you have downloaded Go v 1.10.3 type the following command, if you have downloaded some other version, change the version in the following command.

    Open terminal in downloaded directory and type:

        tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz

    After extracting set up GOPATH, refer the "Setting up environment variables" section.

## Installing on Windows:

??? question "Installing on Windows"

    Download the msi installer from [golang.org](https://golang.org/dl/) and follow the prompts to intsall it.

    By default it is installed in

        c:\Go

    all the environmental variables are set automatically by the installer.

    If you are using windows, I would highly recommend to install Git & Git Bash.

## Setting up environment variables:
    
??? danger "Setting up environment variables"
            
    If you have installed Go using brew on OSx or on linux you need to set up GOPATH.
    
    A) If you are using bash, open and edit:

        ~/.bash_profile

    B) If you are using zsh, open and edit:

        ~/.zshrc

    and add the following:

        export PATH=$PATH:/usr/local/go/bin
        export GOROOT=/usr/local/go
        export GOPATH=$HOME/go
        export GOBIN=$HOME/go/bin

    save the file and restart the terminal or type

    A) For bash:

        source ~/.bash_profile

    B) For zsh:

        source ~/.zshrc

    ??? question "Editing .bash_profile or .zshrc"

        Follow the instructions to edit .bash_profile or .zshrc, open the terminal and type

            cd
        
        If you are using atom type

            atom .bash_profile

        or

            atom .zshrc
        

## Checking the installation:
    
??? warning "Checking the installation"
            
    Open the terminal and type:

        echo $GOROOT

    it should give the output:

        usr/local/go

    For windows it should give output:

        c:\Go
    
    You can also test using the following commands:

        go env
        go version