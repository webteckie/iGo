# GO Setup

## Download and Install Compiler
https://golang.org/doc/install?download=go1.13.7.windows-amd64.msi

Ensure the GOPATH is correctly set to the location of where you want your project packages installed (e.g., C:\Users\me\Source\Repos\GitHub\GoPath).

Do not confuse GOPATH with GOROOT!!! GOROOT is the location of the Go compiler binaries (e.g., c:\Go) and thus should be in your system path.

You can see the value of each by running "go env" in a terminal.

## Visual Studio Code Editor
Preferably use Visual Studio Code as the editor of choice for your Go programming. Install the Go extension. Make sure the open instance of Visual Studio Code has picked up the correct GOPATH setting. To verify this you can:

    - open a Terminal window in Visual Studio code
    - Run the "cmd" command
    - Run the "set" command
    - Look for the GOPATH variable in the listing

    If the variable does not show or the value is incorrect ensure the correct settings for the variable in the control panel Environment Variables section and then restart all instances of Visual Studio.

## Git Extensions
Make sure you have installed the Git Extension Pack extension into Visual Studio Code.

## Clone this Repo into your local machine
Using the Visual Studio Code's Git Clone command clone this repo.

## Run The TestConsole
To verify the setup of your Go environment, open a Terminal window in Visual Studio Code and do the following:

    - cd "<cloned-repo-location>\iGO\src
    - run "go run .\TestConsole.g"

    Should see a few prompts with single key answers.
    