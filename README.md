# GolangEnv
Build Environment of Golang labo

## How to Build

- build
```
$ docker build -t CONTAINER_NAME .
```

- run
```
$ docker run -v [LOCAL_DIR_FULLPATH]:/go -it [CONTANER_NAME] /bin/zsh
```

### Detail

- -v option
mount your local directory on contaner
- -it option
inter contaner with using /bin/zsh
