<div align="center">
  <h1>
    MyIP CLI Tool
  </h1>
  <div align="center">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/mo7amd/myip">
    <a href="https://github.com/mo7amd/myip/blob/master/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/mo7amd/myip"></a>
    <a href="https://github.com/mo7amd/myip/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/mo7amd/myip"></a>
    <div>
      <a href="https://github.com/mo7amd/myip/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/mo7amd/myip"></a>
      <a href="https://github.com/mo7amd/myip/network"><img alt="GitHub forks" src="https://img.shields.io/github/forks/mo7amd/myip"></a>
    </div>
  </div>
  <br>
  <p>
    I use this tool to get my current Global IP address as I need its value to whitelist Home's IP on my company firewall.I needed this tool since I work from home now and my ISP doesn't provide static IPs.
  </p>
</div>
  <hr>

## Optional Dependency

the tool not only get the IP address it also get geolocation info of your IP address using [ipify](https://geo.ipify.org/). they have a nice free plan of 1000 request per month which is nice for personal use.
so if you want the Geolocation part to work please do the following:
- create an account at ipify
- get the apiKey of your own
- rename example.env file to .env file and past your apiKey there on there


otherwise just ignore that part.

- - -
## Download and Install (Unix systems)
- Download Bin from [here]()
- copy or link it to your bin to use it as a program
copy
```
cp main /usr/local/bin/myip
```
link
```
ln -s $PWD/main /usr/local/bin/myip
```
- - -

## Build the project yourself
- clone the repo
```
git clone git@github.com:mo7amd/myip.git
```
- change dir to it
```
cd myip
```
- build the package
```
go build main.go
```
- copy or link it to your bin to use it as a program
copy
```
cp main /usr/local/bin/myip
```
link
```
ln -s $PWD/main /usr/local/bin/myip
```

I personally prefer the link version as it keep the bin updated when I rebuild the package, you got the options suit yourself.