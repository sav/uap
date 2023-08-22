# UAP

A lightweight battery monitor that tracks health and consumption.

## 💾 Installation

```
go install github.com/sav/uap@latest
```

## 📝 Configuration

Begin by copying `uaprc.example` to `~/.uaprc` and proceed to make edits as per your requirements.

```
wget https://raw.githubusercontent.com/sav/uap/master/uaprc.example -O ~/.uaprc
```

To receive notifications about the battery's levels and states, it is **essential** to include a field in the `batteries` section containing its corresponding name.

You can usually locate the battery's name within its sysfs directory, often listed in the uevent file. To view all attributes related to power supplies on a Linux system, execute the following command:

```
cat /sys/class/power_supply/*/uevent
```

## 🔨 Build

### Dependencies

#### Linux

[Oto](https://pkg.go.dev/github.com/hajimehoshi/oto) is a package used to play sounds in Go. On Linux, it relies on ALSA to function.

To install ALSA library on Ubuntu or Debian, run this command:

```
apt install libasound2-dev
```

And on RedHat-based Linux distributions, run:

```
dnf install alsa-lib-devel
```
