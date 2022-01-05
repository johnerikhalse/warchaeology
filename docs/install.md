# Installation

## Linux
### RPM package
```
curl -LO https://github.com/nlnwa/warchaeology/releases/latest/download/warchaeology_0.1.0-RC.7_x86_64.rpm
sudo rpm -Uvh warchaeology_0.1.0-RC.7_x86_64.rpm
```

### Debian package
```
curl -LO https://github.com/nlnwa/warchaeology/releases/latest/download/warchaeology_0.1.0-RC.7_amd64.deb
sudo dpkg -i warchaeology_0.1.0-RC.7_amd64.deb
```

### Binary download
```
curl -LO https://github.com/nlnwa/warchaeology/releases/latest/download/warc_linux_x86_64
sudo install warc_linux_x86_64 /usr/local/bin/warc
```

#### Command completion
bash
```
sudo sh -c "/usr/local/bin/warc completion bash > /etc/bash_completion.d/warc"
```

zshell
```
source &lt;(warc completion zsh)
```

[{{ site.github.latest_release.name }}]({{ site.github.latest_release }})

[{{ site.github.repository_url }}]({{ site.github.repository_url }})

