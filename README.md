<h1 align="center">
  <img alt="Superchan" src="assets/icon.svg" width="500px"/><br/>
  Superchan
</h1>
<p align="center">Superchan is the same chan but generic and with rate calculator, capable of monitoring input and output data transmission rate.</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/superchan?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" />
</p>


# Installation
```shell
go get github.com/mehditeymorian/superchan
```


# Usage
```Go
// Create a new Superchan
c := superchan.New[type](channelSize)

// send data into the channel
c.Send(n)

// receive data from the channel
c.Receive()

// monitor rates
log.Printf("buffer: %d in: %d/s out: %d/s\n", c.BufferedSize(), c.InputRate(), c.OutputRate())
```