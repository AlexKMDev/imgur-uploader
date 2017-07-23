# imgur-uploader

### Installation

1. Create Imgur.com application using [this link](https://api.imgur.com/oauth2/addclient).
2. Paste your Client ID to `~/.config/imgur-uploader/config.yaml`

```yaml
ClientID: 123ff
```

3. Install uploader
```bash
go get github.com/Anakros/imgur-uploader
```

4. Start uploading everything

```bash
./imgur-uploader -i ~/Downloads/12345.jpg
```
5. Or deleting if you uploaded something wrong

```bash
./imgur-uploader -d 345345DeleteHash77
```
