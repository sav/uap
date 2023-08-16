# UAP

A lightweight battery monitor that tracks health and consumption.

## Build

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

