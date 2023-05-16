# bower

```bash
npm i -D bower
npx bower install bootstrap5
```

`bower_components`에 라이브러리가 설치된다.

Github URL을 기반으로 진행.

# bower-installer

 - `bower.json`

```json
{
    "name": "it-is-auto-move-libraries-tools",
    "dependencies": {
        "backbone": "latest",
    },
    "install": {
        "path": "public",
        "sources": {
            "backbone": "bower_components/backbone/backbone.js"
        }
    }
}
```