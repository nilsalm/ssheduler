# SSHeduler

Async communication and command scheduling for remote devices (IoT).
*Queen* can add commands (basically shell scripts, potentially other stuff) to a queue which the *Pawn* pulls upon next wake and executes.
The idea is to be able to run commands and programs on devices that are not always online or that are behind some form of firewall or 4G-connected, so that you cant remote login to the device itself.

| Short hand | Description                       |
| ---------- | --------------------------------- |
| Queen      | The admin machine or main machine |
| Pawn       | Remote machines, the followers    |

## Code structure

The main model returns either the TUI for the Queen or Pawn, or renders a simple choice between the two which is written to a config file. This logic is coded in `common.go` and the TUI for that is in `commonui.go`.

The Queen UI then lets the user choose between **Scheduling a command** or **Registering a new Pawn**.
The Pawn UI lets the user choose between **Manually executing a command**, **Linking this pawn to a queen** or **Running the pawn in autorun**.

The logic/functionality is coded in `queen.go` and `pawn.go`, respectively.

You can think of it as a tree that represents the three modes of operation:

![docs/architecture.png](docs/architecture.png)







## MVP

Using a [CHARM server.](https://github.com/charmbracelet/charm)

**A** Machine 1 (M1) logs in, uploads a command (echo “Hello World”) and downloads a command to execute. 

**B** M1 uploads a command and R1 downloads it. Requires multiuser setting. -> run “charm link” on M1 and enter the code on R1. 

**C** Differeniate between M and R in the same exe

**D** Wrap the whole thing in a sweetlooking GUI

Let's get started! 🚀

1. [Install GO on raspi](https://www.e-tinkers.com/2019/06/better-way-to-install-golang-go-on-raspberry-pi/) NOTE: go1.18.4.linux-arm64.tar.gz for raspi
2. Installing charm with ```brew tap charmbracelet/tap && brew install charmbracelet/tap/charm```
3. Starting up with ```charm serve``` on M1
4. Linking a new machine with ```charm link``` on master and pasting the code on Rn
5. 




## Compilation

Cross-platform compile for Raspberry Pi on mac
```
env GOOS=linux GOARCH=arm64 go build -o ssheduler-linux-arm64 ssheduler.go
```

Compile for mac on mac

```
go build -o ssheduler-mac ssheduler.go
```

Copy exe to Raspi:
```
scp ssheduler-linux-arm64 herbert@herbert.local:.
```

Shorthand:
```
env GOOS=linux GOARCH=arm64 go build -o ssheduler-linux-arm64 ssheduler.go && scp ssheduler-linux-arm64 herbert@herbert.local:.
env GOOS=linux GOARCH=arm64 go build -o ssheduler-linux-arm64 ssheduler.go && scp ssheduler-linux-arm64 herbert@potatislandet.ddns.net:.
```

