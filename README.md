<h1 align="center">catcat</h1>

<p align="center">
  <a href="https://github.com/crptmem/catcat/stargazers"><img src="https://img.shields.io/github/stars/crptmem/catcat?colorA=151515&colorB=B66467&style=for-the-badge&logo=starship"></a>
  <a href="https://github.com/crptmem/catcat/issues"><img src="https://img.shields.io/github/issues/crptmem/catcat?colorA=151515&colorB=8C977D&style=for-the-badge&logo=bugatti"></a>
  <a href="https://github.com/crptmem/catcat/network/members"><img src="https://img.shields.io/github/forks/crptmem/catcat?colorA=151515&colorB=D9BC8C&style=for-the-badge&logo=github"></a>
</p>

> 🐈 A simple and customizable launcher for running games through Wine.
<p align="center">
  <img align="center" src=https://github.com/crptmem/catcat/assets/88046785/b1f40ce6-fe23-4ead-a466-22bda52e0d43 />
</p>

## 📦 Build and install
You should have GOPATH in your PATH environment variable.
```sh
git clone https://github.com/crptmem/catcat.git && cd catcat
go build .
go install
```

## 🔧 Configuration
Run `catcat` for the first time. <br/>
Open `$HOME/.config/catcat/config.json` in your favourite editor. <br/>
Set `winelocation` to your Wine binary location (like `/etc/eselect/wine/bin/wine` in Gentoo Linux) <br/>
Add games to `gameentries` like this:
```json
{
  "gameentries" : [{
      "name": "Game name",
      "executablepath": "Game executable path",
      "wrappercommand": "Some wrapper command like mangohud",
      "discordpresence": true
  }],
  "winelocation": "your wine location"
}
```
<br/>
Also, if you want to use Discord presence, use <a href="https://github.com/0e4ef622/wine-discord-ipc-bridge">this repository</a>. <br/>
Download or compile `winediscordipcbridge.exe` and put it to `$HOME/.config/catcat/bin/`


